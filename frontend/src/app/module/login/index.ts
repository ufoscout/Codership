import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login.component';
import { LoginRoutingModule } from './login.routing';
import { AuthModule } from '../auth';
import { SharedModule } from '../shared';

@NgModule({
  declarations: [
    LoginComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    AuthModule,
    LoginRoutingModule,
    SharedModule,
  ],
  providers: []
})
export class LoginModule { }
