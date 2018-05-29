import { NgModule } from '@angular/core';
import { NgxsModule } from '@ngxs/store';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { HttpClientModule } from '@angular/common/http';
import { AuthState } from './auth.state';
import { AuthInterceptor } from './auth.interceptor';

@NgModule({
  declarations: [
  ],
  imports: [
    HttpClientModule,
    NgxsModule.forFeature([AuthState])
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ]
})
export class AuthModule { }
