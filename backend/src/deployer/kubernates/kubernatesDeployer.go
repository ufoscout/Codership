package kubernates

import (
	"github.com/ufoscout/Codership/backend/src/deployer/common"
	"errors"
)

type kubernatesDeployer struct {
	name string
}

func NewKubernatesDeployer() common.Deployer {
	return &kubernatesDeployer{}
}

func (k *kubernatesDeployer) DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) ([]common.Node, error) {
	return nil, errors.New("Kubernates deployment not yet implemented")
}

func (k *kubernatesDeployer) RemoveCluster(clusterName string) (bool, error) {
	return false, errors.New("Kubernates deployment not yet implemented")
}

func (k *kubernatesDeployer) ClusterStatus(clusterName string) (map[string]string, error) {
	return nil, errors.New("Kubernates deployment not yet implemented")
}