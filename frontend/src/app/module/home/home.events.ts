import { CreateClusterDTO } from './dto.model';

export class CreateCluster {
    static type = '[Home] Create Cluster';
    constructor(public readonly payload: {deployerType: string, dto: CreateClusterDTO}) { }
}
