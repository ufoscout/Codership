import { Component } from '@angular/core';
import { Store, Select } from '@ngxs/store';
import { SetLanguage } from '../common.events';
import { CommonState, CommonStateModel } from '../common.state';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-common-header',
  templateUrl: './header.component.html'
})
export class HeaderComponent {

  @Select(CommonState) commonState$: Observable<CommonStateModel>;
  navbarCollapsed = true;

  constructor(private store: Store) {
  }

  setLanguage(language: String): void {
    this.store.dispatch([
      new SetLanguage({language: language}),
    ]);
  }

}
