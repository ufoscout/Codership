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
	public id: string;
	public status: string;
	public port: number;
}

export class Error {
	public error: string;
}