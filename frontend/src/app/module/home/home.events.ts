import { CreateClusterDTO } from './dto.model';

export class CreateCluster {
    static type = '[Home] Create Cluster';
    constructor(public readonly payload: {deployerType: string, dto: CreateClusterDTO}) { }
}

export class RefreshCluster {
    static type = '[Home] Refresh Cluster';
    constructor(public readonly payload: {deployerType: string, clusterName: string}) { }
}

export class DeleteCluster {
    static type = '[Home] Delete Cluster';
    constructor(public readonly payload: {deployerType: string, clusterName: string}) { }
}