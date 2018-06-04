package common

type Deployer interface {
	DeployCluster(clusterName string, dbType string, clusterSize int, firstHostPort int) (Nodes, error)
	RemoveCluster(clusterName string) (bool, error)
	ClusterStatus(clusterName string) (map[string]string, error)
}

type Nodes []Node

type Node struct {
	Id string      `json: id`
	Status string  `json: status`
	Port int       `json: port`
}

func NewNode(id string, status string, port int) Node {
	return Node {
		Id: id,
		Status: status,
		Port: port,
	}
}