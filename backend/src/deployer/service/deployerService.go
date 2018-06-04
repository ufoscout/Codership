package service

import (
	"github.com/ufoscout/Codership/backend/src/deployer/common"
	"errors"
)

type DeployerService interface {
	GetDeployer(deployerType string) (common.Deployer, error)
}

type deployerService struct {
	deployers map[string]common.Deployer
}

func NewDeployerService(deployers map[string]common.Deployer) DeployerService {
	return &deployerService {
		deployers: deployers,
	}
}

func (s *deployerService) GetDeployer(deployerType string) (common.Deployer, error) {
	deployer, ok := s.deployers[deployerType]
	if ok {
		return deployer, nil
	}
	return nil, errors.New("Deployer of type [" + deployerType + "] does not exist")
}