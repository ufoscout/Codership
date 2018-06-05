import { NgModule } from '@angular/core';
import { CommonModule as NgCommonModule } from '@angular/common';
import { HttpClientModule, HttpClient } from '@angular/common/http';
import { TranslateModule, TranslateLoader } from '@ngx-translate/core';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';
import { NgbCollapseModule } from '@ng-bootstrap/ng-bootstrap/collapse/collapse.module';
import { NgbDropdownModule } from '@ng-bootstrap/ng-bootstrap/dropdown/dropdown.module';
import { RouterModule } from '@angular/router';
import { HeaderComponent } from './header/header.component';
import { NgxsModule } from '@ngxs/store';
import { CommonState } from './common.state';
import { NgxSpinnerModule } from 'ngx-spinner';

export function HttpLoaderFactory(http: HttpClient): TranslateHttpLoader {
  return new TranslateHttpLoader(http);
}

@NgModule({
  declarations: [
    HeaderComponent,
  ],
  imports: [
    HttpClientModule,
    NgCommonModule,
    NgxsModule.forFeature([CommonState]),
    NgbCollapseModule,
    NgbDropdownModule,
    NgxSpinnerModule,
    RouterModule,
    TranslateModule.forRoot({
      loader: {
        provide: TranslateLoader,
        useFactory: HttpLoaderFactory,
        deps: [HttpClient]
      }
    })
  ],
  providers: [],
  exports: [
    TranslateModule,
    HeaderComponent,
  ]
})
export class CommonModule { }
