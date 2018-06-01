package common

type Deployer interface {
	DeployCluster(clusterName string, dbType string, instances int) (bool, error)
	RemoveCluster(clusterName string) (bool, error)
}
