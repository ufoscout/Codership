import { Store, NgxsModule } from '@ngxs/store';
import { async, TestBed } from '@angular/core/testing';
import { AuthState } from './auth.state';
import { SetAuthData } from './auth.events';

describe('[Auth] Auth State', () => {
    let store: Store;

    beforeEach(async(() => {
        TestBed.configureTestingModule({
            imports: [NgxsModule.forRoot([AuthState])],
        }).compileComponents();
        store = TestBed.get(Store);
    }));

    it('it should set username and password', () => {
        const username = 'username-' + new Date().getTime();
        const token = 'token-' + new Date().getTime();
        store.dispatch(new SetAuthData({ username: username, token: token }));
        store.selectOnce(AuthState).subscribe(state => {
            expect(state.username).toBe(username);
            expect(state.token).toBe(token);
        });
    });

});
