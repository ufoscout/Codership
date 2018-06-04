import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { CreateClusterDTO, Node } from './dto.model';

@Injectable()
export class ClusterService {

  private baseUrl = '/api/v1/cluster';

  constructor(private http: HttpClient) {
  }

  createCluster(deployerType: string, dto: CreateClusterDTO): Observable<Array<Node>> {
    return this.http.post<Array<Node>>(`${this.baseUrl}/${deployerType}/`, dto)
  }

  getClusterStatus(deployerType: string, clusterName: string): Observable<Map<string,string>> {
    return this.http.get<Map<string,string>>(`${this.baseUrl}/${deployerType}/${clusterName}`);
  }

  deleteCluster(deployerType: string, clusterName: string): Observable<string> {
    return this.http.delete<string>(`${this.baseUrl}/${deployerType}/${clusterName}`);
  }

}
