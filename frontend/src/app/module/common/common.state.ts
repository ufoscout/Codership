import { State, Action, StateContext } from '@ngxs/store';
import { TranslateService } from '@ngx-translate/core';
import * as events from './common.events';
import { ApplicationRef } from '@angular/core';

export const defaultLanguage = 'en';

export class CommonStateModel {
  allLanguages = ['en', 'fi'];
  language = defaultLanguage;
  defaultLanguage = defaultLanguage;
}

@State<CommonStateModel>({
  name: 'common',
  defaults: new CommonStateModel(),
})
export class CommonState {

  constructor(private translate: TranslateService, private ref: ApplicationRef) {
    translate.setDefaultLang(defaultLanguage);
    translate.use(defaultLanguage);
  }

  @Action(events.SetLanguage)
  setToken({ getState, setState }: StateContext<CommonStateModel>, { payload }: events.SetLanguage) {
    const state = getState();
    setState({
      ...state,
      language: payload.language,
    });
    this.translate.use(payload.language);
    this.ref.tick();
  }

}
