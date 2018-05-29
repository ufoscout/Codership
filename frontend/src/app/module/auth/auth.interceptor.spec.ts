import { TestBed } from '@angular/core/testing';
import {
    HttpClientTestingModule,
    HttpTestingController,
} from '@angular/common/http/testing';
import { AuthInterceptor } from './auth.interceptor';
import { HTTP_INTERCEPTORS, HttpClient } from '@angular/common/http';
import { Store, NgxsModule } from '@ngxs/store';
import { AuthState } from './auth.state';
import { SetAuthData } from './auth.events';

describe(`[Auth] AuthInterceptor`, () => {

    let store: Store;
    let httpMock: HttpTestingController;
    let httpClient: HttpClient;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [
                HttpClientTestingModule,
                NgxsModule.forRoot([AuthState]),
            ],
            providers: [
                {
                    provide: HTTP_INTERCEPTORS,
                    useClass: AuthInterceptor,
                    multi: true,
                },
            ],
        });
        httpMock = TestBed.get(HttpTestingController);
        httpClient = TestBed.get(HttpClient);
        store = TestBed.get(Store);
    });

    it('should add an Authorization header', () => {
        const token = 'token-' + new Date().getTime();
        store.dispatch(new SetAuthData({ username: '', token: token }));

        httpClient.get('/url').subscribe();
        const httpRequest = httpMock.expectOne(`/url`);

        expect(httpRequest.request.headers.get('Authorization')).toBe(`Bearer ${token}`);
    });

    it('should not add an Authorization header if the token is not present', () => {
        httpClient.get('/url').subscribe();
        const httpRequest = httpMock.expectOne(`/url`);
        expect(httpRequest.request.headers.get('Authorization')).toBeNull();
    });

});
