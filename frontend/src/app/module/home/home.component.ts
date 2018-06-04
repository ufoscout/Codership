import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CreateClusterDTO } from './dto.model';
import * as obj from '../shared/utils/object.utils'

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html'
})
export class HomeComponent implements OnInit {

  constructor(private modalService: NgbModal) { }

  model = new CreateClusterDTO();
  deploymentType = "docker";

  ngOnInit() {
    this.model.dbType = "mariadb"
  }

  open(content) {
    this.modalService.open(content).result.then((result) => {
    }, (reason) => {
    });
  }

  portsText(): string {
    let text = "Ports: ";
    if (obj.exists(this.model.clusterSize) && obj.exists(this.model.firstHostPort)) {
      for (let i = 0; i < this.model.clusterSize; i++) {
        text = text + "'Node-" + i + "':" + (this.model.firstHostPort + i) + " "
      }
    }
    console.log(this.deploymentType)
    console.log(JSON.stringify(this.model))
    return text;
  }

}
