package web

type CreateClusterDTO struct {
	ClusterName string  `json:"clusterName"`
	DbType string       `json:"dbType"`
	ClusterSize int     `json:"clusterSize"`
	FirstHostPort int   `json:"firstHostPort"`
}
