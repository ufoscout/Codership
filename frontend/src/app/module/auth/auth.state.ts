import { State, Action, StateContext } from '@ngxs/store';
import * as events from './auth.events';

export class AuthStateModel {
    username = '';
    token = '';
}

@State<AuthStateModel>({
  name: 'auth',
  defaults: new AuthStateModel(),
})
export class AuthState {

    @Action(events.SetAuthData)
    setToken({ getState, setState }: StateContext<AuthStateModel>, {payload}: events.SetAuthData) {
      const state = getState();
      setState({
        ...state,
        username: payload.username,
        token: payload.token
      });
    }

}
