package common

type Deployer interface {
	DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) ([]string, error)
	RemoveCluster(clusterName string) (bool, error)
	ClusterStatus(clusterName string) (map[string]string, error)
}
