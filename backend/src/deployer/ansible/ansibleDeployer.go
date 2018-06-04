package ansible

import (
	"github.com/ufoscout/Codership/backend/src/deployer/common"
	"errors"
)

type ansibleDeployer struct {
	name string
}

func NewAnsibleDeployer() common.Deployer {
	return &ansibleDeployer{}
}

func (k *ansibleDeployer) DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) ([]common.Node, error) {
	return nil, errors.New("Ansible deployment not yet implemented")
}

func (k *ansibleDeployer) RemoveCluster(clusterName string) (bool, error) {
	return false, errors.New("Ansible deployment not yet implemented")
}

func (k *ansibleDeployer) ClusterStatus(clusterName string) (map[string]string, error) {
	return nil, errors.New("Ansible deployment not yet implemented")
}