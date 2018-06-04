import { NgModule } from '@angular/core';
import { CommonModule as NgCommonModule } from '@angular/common';
import { CommonModule } from '../common';
import { HomeComponent } from './home.component';
import { HomeRoutingModule } from './home.routing';
import { NgbModalModule } from '@ng-bootstrap/ng-bootstrap/modal/modal.module';
import { FormsModule } from '@angular/forms';
import { NgxsModule } from '@ngxs/store';
import { HomeState } from './home.state';
import { ClusterService } from './home.service';

@NgModule({
  declarations: [
    HomeComponent,
  ],
  imports: [
    FormsModule,
    NgCommonModule,
    CommonModule,
    HomeRoutingModule,
    NgbModalModule,
    NgxsModule.forFeature([HomeState]),
  ],
  providers: [
    ClusterService,
  ]
})
export class HomeModule { }
