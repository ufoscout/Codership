export class CreateClusterDTO {
    public clusterName: string;
	public dbType: string;
	public clusterSize: number;
	public firstHostPort: number;
}

export class Cluster {
	constructor(
		public name: string,
		public deployerType: string,
		public nodes: Array<Node>
	) {}
}

export class Node {
	public Id: string;
	public Status: string;
	public Port: number;
}

export class Error {
	public error: string;
}