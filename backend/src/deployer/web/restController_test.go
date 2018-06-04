package web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"github.com/ufoscout/Codership/backend/src/deployer/common"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
	"encoding/json"
	"strconv"
	"io"
	"bytes"
	"fmt"
)

func Test_Get_ShouldUseReturnClusterStatusIfExists(t *testing.T) {
	router := gin.Default()
	deployer := &MockDeployer{
		result: map[string]string{
			"one": "statusOne",
			"two": "statusTwo",
		},
	}
	service := &MockDeployerService{
		deployer: deployer,
	}
	web := NewRestController(router, service)
	web.Start()

	response := performRequest(router, "GET", "/api/v1/cluster/dock/12345", nil)
	assert.Equal(t, "dock", service.calledDeployerType)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, "status", deployer.methodCalled)
	assert.Equal(t, "12345", deployer.params[0])

	var body map[string]string
	json.Unmarshal([]byte(response.Body.String()), &body)
	assert.Equal(t, 2, len(body))
	assert.Equal(t, "statusOne", body["one"])
	assert.Equal(t, "statusTwo", body["two"])
}

func Test_Get_ShouldFailIfUnknownDeployer(t *testing.T) {
	router := gin.Default()
	service := &MockDeployerService{
		err: errors.New("CustomError"),
	}
	web := NewRestController(router, service)
	web.Start()

	response := performRequest(router, "GET", "/api/v1/cluster/dock/12345", nil)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	var body ErrorResponse
	json.Unmarshal([]byte(response.Body.String()), &body)
	assert.Equal(t, "CustomError", body.Error)
}

func Test_Delete_ShouldRemoveTheCluster(t *testing.T) {
	router := gin.Default()
	deployer := &MockDeployer{
		result: true,
	}
	service := &MockDeployerService{
		deployer: deployer,
	}
	web := NewRestController(router, service)
	web.Start()

	response := performRequest(router, "DELETE", "/api/v1/cluster/docker/12345", nil)
	assert.Equal(t, "docker", service.calledDeployerType)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, "remove", deployer.methodCalled)
	assert.Equal(t, "12345", deployer.params[0])

	var body map[string]bool
	json.Unmarshal([]byte(response.Body.String()), &body)
	assert.Equal(t, 1, len(body))
	assert.Equal(t, true, body["deleted"])
}

func Test_Delete_ShouldReturnErrorFromDeployer(t *testing.T) {
	router := gin.Default()
	deployer := &MockDeployer{
		err: errors.New("CustomRemoveError"),
	}
	service := &MockDeployerService{
		deployer: deployer,
	}
	web := NewRestController(router, service)
	web.Start()

	response := performRequest(router, "DELETE", "/api/v1/cluster/dock/12345", nil)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.Equal(t, "remove", deployer.methodCalled)
	assert.Equal(t, "12345", deployer.params[0])

	var body ErrorResponse
	json.Unmarshal([]byte(response.Body.String()), &body)
	assert.Equal(t, "CustomRemoveError", body.Error)
}

func Test_Post_ShouldCreateNewCluster(t *testing.T) {
	router := gin.Default()
	deployer := &MockDeployer{
		result: common.Nodes{
			common.NewNode("one", "oneStatus", 321),
			common.NewNode("two", "twoStatus", 123),
		},
	}
	service := &MockDeployerService{
		deployer: deployer,
	}
	web := NewRestController(router, service)
	web.Start()

	dto := CreateClusterDTO{
		ClusterName: "newClusterName",
		DbType: "mariadb",
		ClusterSize: 3,
		FirstHostPort: 12345,
	}

	response := performRequest(router, "POST", "/api/v1/cluster/docker/", dto)
	assert.Equal(t, "docker", service.calledDeployerType)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.Equal(t, "deploy", deployer.methodCalled)
	assert.Equal(t, "newClusterName", deployer.params[0])
	assert.Equal(t, "mariadb", deployer.params[1])
	assert.Equal(t, "3", deployer.params[2])
	assert.Equal(t, "12345", deployer.params[3])

	fmt.Println("Body is: " + response.Body.String())

	var body []common.Node
	json.Unmarshal([]byte(response.Body.String()), &body)
	assert.Equal(t, 2, len(body))

	assert.Equal(t, "one", body[0].Id)
	assert.Equal(t, "oneStatus", body[0].Status)
	assert.Equal(t, 321, body[0].Port)

	assert.Equal(t, "two", body[1].Id)
	assert.Equal(t, "twoStatus", body[1].Status)
	assert.Equal(t, 123, body[1].Port)
}



func performRequest(r http.Handler, method string, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody io.Reader
	if (body!=nil) {
		jsonReq, err := json.Marshal(body)
		if err!=nil {
			panic(err)
		}
		reqBody = bytes.NewReader(jsonReq)
	}
	req, _ := http.NewRequest(method, path, reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// Mocks

type MockDeployerService struct {
	err error
	deployer *MockDeployer
	calledDeployerType string
}

func (s *MockDeployerService) GetDeployer(deployerType string) (common.Deployer, error) {
	s.calledDeployerType = deployerType
	if s.err != nil {
		return nil, s.err
	}
	return s.deployer, nil
}

type MockDeployer struct {
	err error
	result interface{}
	methodCalled string
	params []string
}

func (k *MockDeployer) DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) (common.Nodes, error) {
	k.methodCalled = "deploy"
	k.params = []string{
		clusterName,
		dbType,
		strconv.Itoa(clusterSize),
		strconv.Itoa(firstHostPort),
	}
	if k.err != nil {
		return nil, k.err
	}
	return k.result.(common.Nodes), nil
}

func (k *MockDeployer) RemoveCluster(clusterName string) (bool, error) {
	k.methodCalled = "remove"
	k.params = []string{
		clusterName,
	}
	if k.err != nil {
		return false, k.err
	}
	return k.result.(bool), nil
}

func (k *MockDeployer) ClusterStatus(clusterName string) (map[string]string, error) {
	k.methodCalled = "status"
	k.params = []string{
		clusterName,
	}
	if k.err != nil {
		return nil, k.err
	}
	return k.result.(map[string]string), nil
}