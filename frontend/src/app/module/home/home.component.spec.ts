import { TestBed, async, inject } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HomeModule } from './';
import { HomeComponent } from './home.component';
import { NgxsModule } from '@ngxs/store';

describe(`HomeComponent`, () => {

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [
                HttpClientTestingModule,
                RouterTestingModule,
                NgxsModule.forRoot([]),
                HomeModule,
            ],
            providers: [
                HomeComponent,
            ],
        });
    });

    it(`should instantiate the component`, async(
        inject([HomeComponent],
            (component: HomeComponent) => {
                expect(component).not.toBeUndefined();
            })));

});
