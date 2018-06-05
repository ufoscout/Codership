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
          clusters: state.clusters.concat( new Cluster(payload.dto.clusterName, payload.deployerType, nodes))
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
        return of('Error while creating the cluster')
      })
    )
  }

}
