import { NgModule } from '@angular/core';
import { CommonModule } from '../common';
import { HomeComponent } from './home.component';
import { HomeRoutingModule } from './home.routing';
import { NgbModalModule } from '@ng-bootstrap/ng-bootstrap/modal/modal.module';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    HomeComponent,
  ],
  imports: [
    FormsModule,
    CommonModule,
    HomeRoutingModule,
    NgbModalModule,
  ],
  providers: []
})
export class HomeModule { }
