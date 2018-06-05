import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CreateClusterDTO } from './dto.model';
import * as obj from '../shared/utils/object.utils'
import { Select, Store } from '@ngxs/store';
import { HomeState, HomeStateModel } from './home.state';
import { Observable } from 'rxjs';
import { CreateCluster, RefreshCluster, DeleteCluster } from './home.events';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html'
})
export class HomeComponent implements OnInit {

  
  @Select(HomeState) homeState$: Observable<HomeStateModel>;
  model = new CreateClusterDTO();
  deploymentType = "docker";

  constructor(private store: Store, private modalService: NgbModal) { }
  
  ngOnInit(): void {
    this.model.dbType = "mariadb"
  }

  open(content): void {
    this.modalService.open(content).result.then((result) => {
    }, (reason) => {
    });
  }

  submitCreateForm(): void {
    console.log("submit form")
    this.store.dispatch([
      new CreateCluster({deployerType: this.deploymentType, dto: this.model}),
    ]);
  }

  refreshCluster(deployerType: string, clusterName: string) {
    console.log(`refresh ${deployerType} [${clusterName}]`)
    this.store.dispatch([
      new RefreshCluster({deployerType: deployerType, clusterName: clusterName}),
    ]);
  }


  deleteCluster(deployerType: string, clusterName: string) {
    console.log(`delete ${deployerType} [${clusterName}]`);
    this.store.dispatch([
      new DeleteCluster({deployerType: deployerType, clusterName: clusterName}),
    ]);
  }

  portsText(): string {
    let text = "Ports: ";
    if (obj.exists(this.model.clusterSize) && obj.exists(this.model.firstHostPort)) {
      for (let i = 0; i < this.model.clusterSize; i++) {
        text = text + "'Node-" + i + "':" + (this.model.firstHostPort + i) + " - "
      }
    }
    return text;
  }

}
