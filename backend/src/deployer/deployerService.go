package deployer

type DeployerService interface {
	Something()
}

type deployerService struct {

}

func NewDeployerService() DeployerService {
	return &deployerService {

	}
}

func (s *deployerService) Something() {

}