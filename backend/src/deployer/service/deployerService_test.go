package service

import (
	"testing"
	"github.com/ufoscout/Codership/backend/src/deployer/common"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnErrorIfNotPresent(t *testing.T) {

	deployers := map[string]common.Deployer {
		"one": FakeDeployer{"one"},
		"two": FakeDeployer{"two"},
	}

	service := NewDeployerService(deployers)

	_, err := service.GetDeployer("three")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Deployer of type [three] does not exist")

}

func Test_ShouldReturnTheExpectedDeployer(t *testing.T) {

	deployers := map[string]common.Deployer {
		"one": FakeDeployer{"one"},
		"two": FakeDeployer{"two"},
	}

	service := NewDeployerService(deployers)

	deployer, err := service.GetDeployer("one")
	assert.Nil(t, err)
	assert.Equal(t, "one", deployer.(FakeDeployer).name)

}

type FakeDeployer struct {
	name string
}

func (FakeDeployer) DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) ([]common.Node, error) {
	return nil, nil
}

func (FakeDeployer) RemoveCluster(clusterName string) (bool, error) {
	return false, nil
}

func (FakeDeployer) ClusterStatus(clusterName string) (map[string]string, error) {
	return nil, nil
}