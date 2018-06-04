package common

type Deployer interface {
	DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) ([]Node, error)
	RemoveCluster(clusterName string) (bool, error)
	ClusterStatus(clusterName string) (map[string]string, error)
}

type Node struct {
	Id string
	Status string
	Port int
}

func NewNode(id string, status string, port int) Node {
	return Node {
		Id: id,
		Status: status,
		Port: port,
	}
}