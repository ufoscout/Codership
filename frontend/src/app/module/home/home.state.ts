import { State, Action, StateContext } from '@ngxs/store';
import { TranslateService } from '@ngx-translate/core';
import * as events from './home.events';
import { Cluster } from './dto.model';
import { ClusterService } from './home.service';
import { tap, catchError } from 'rxjs/operators';
import { of } from 'rxjs';
import { NgxSpinnerService } from 'ngx-spinner';
import { ApplicationRef } from '@angular/core';

export const defaultLanguage = 'en';

export class HomeStateModel {
  clusters: Array<Cluster> = []
}

@State<HomeStateModel>({
  name: 'home',
  defaults: new HomeStateModel(),
})
export class HomeState {

  constructor(private service: ClusterService,
    private spinner: NgxSpinnerService,
    private ref: ApplicationRef) {
  }

  @Action(events.CreateCluster)
  createCluster({ getState, setState }: StateContext<HomeStateModel>, { payload }: events.CreateCluster) {
    this.spinner.show();
    return this.service.createCluster(payload.deployerType, payload.dto).pipe(
      tap(nodes => {
        const state = getState();
        setState({
          ...state,
          clusters: state.clusters.slice().concat( new Cluster(payload.dto.clusterName, payload.deployerType, nodes))
        });
        this.spinner.hide();
        this.ref.tick();
      }
      ),
      catchError(err => {
        this.spinner.hide();
        this.ref.tick();
        console.info(err);
        alert(err.error.error)
        return of(err.error.error)
      })
    )
  }

  @Action(events.RefreshCluster)
  refreshCluster({ getState, setState }: StateContext<HomeStateModel>, { payload }: events.RefreshCluster) {
    this.spinner.show();
    return this.service.getClusterStatus(payload.deployerType, payload.clusterName).pipe(
      tap(nodes => {
        const state = getState();
        let clusters = state.clusters.slice();
        clusters.forEach(cluster => {
          if (cluster.name === payload.clusterName) {
            cluster.nodes.forEach(node => {
              node.status = nodes[node.id];
            });
          }
        });
        setState({
          ...state,
          clusters: clusters
        });
        this.spinner.hide();
        this.ref.tick();
      }
      ),
      catchError(err => {
        this.spinner.hide();
        this.ref.tick();
        console.info(err);
        alert(err.error.error)
        return of(err.error.error)
      })
    )
  }

  @Action(events.DeleteCluster)
  deleteCluster({ getState, setState }: StateContext<HomeStateModel>, { payload }: events.DeleteCluster) {
    this.spinner.show();
    return this.service.deleteCluster(payload.deployerType, payload.clusterName).pipe(
      tap(nodes => {
        const state = getState();
        let clusters = state.clusters.slice();
        for (var i = clusters.length - 1; i >= 0; i--) {
          if (clusters[i].name === payload.clusterName) { 
            clusters.splice(i, 1);
          }
      }
        setState({
          ...state,
          clusters: clusters
        });
        this.spinner.hide();
        this.ref.tick();
      }
      ),
      catchError(err => {
        this.spinner.hide();
        this.ref.tick();
        console.info(err);
        alert(err.error.error)
        return of(err.error.error)
      })
    )
  }

}
