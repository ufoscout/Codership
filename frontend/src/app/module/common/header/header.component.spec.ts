import { TestBed, async, inject, ComponentFixture, fakeAsync, tick } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { CommonModule } from '../';
import { HeaderComponent } from './header.component';
import { NgxsModule, Store } from '@ngxs/store';
import { CommonState, defaultLanguage, CommonStateModel } from '../common.state';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

describe('HeaderComponent', () => {
    let fixture: ComponentFixture<HeaderComponent>;
    let component: HeaderComponent;
    let store: Store;

    beforeEach(async(() => {
        TestBed.configureTestingModule({
            imports: [
                HttpClientTestingModule,
                RouterTestingModule,
                NgbModule.forRoot(),
                NgxsModule.forRoot([]),
                CommonModule,
            ],
            providers: []
        }).compileComponents().then(() => {
            fixture = TestBed.createComponent(HeaderComponent);
            component = fixture.componentInstance;
            store = TestBed.get(Store);
        });
    }));

    it('should be instanciated', async(() => {
        expect(component).not.toBeUndefined();
    }));

    it('should show the current language', fakeAsync(() => {
        fixture.detectChanges();
        const state = store.selectSnapshot<CommonStateModel>(s => s.common);
        const selectedLanguage = fixture.debugElement.nativeElement.querySelector('#dropdownLanguage');
        expect(selectedLanguage.textContent.trim()).toBe(state.language);
    }));

    it('should list all languages', fakeAsync(() => {
        fixture.detectChanges();
        const state = store.selectSnapshot<CommonStateModel>(s => s.common);
        const languages: HTMLCollection = fixture.debugElement.nativeElement.querySelector('#dropdownLanguages').children;

        expect(languages.length).toBe(state.allLanguages.length);
        for (let i = 0; i < state.allLanguages.length; i++) {
            // console.log(languages.item(i));
            expect(languages.item(i).textContent.trim()).toBe(state.allLanguages[i]);
        }

    }));

    it('should change the selected language', fakeAsync(() => {
        fixture.detectChanges();
        spyOn(component, 'setLanguage').and.callThrough();

        let state = store.selectSnapshot<CommonStateModel>(s => s.common);

        // Select a language different than the current one
        const previouslySelectedLanguage = state.language;
        let notDefaultLanguageIndex = 0;
        for (let i = 0; i < state.allLanguages.length; i++) {
            if (previouslySelectedLanguage !== state.allLanguages[i]) {
                notDefaultLanguageIndex = i;
            }
        }

        // click on the dropdown entry for the new language
        const languages: HTMLCollection = fixture.debugElement.nativeElement.querySelector('#dropdownLanguages').children;
        (<any>languages.item(notDefaultLanguageIndex)).click();

        tick();
        fixture.detectChanges();

        expect(component.setLanguage).toHaveBeenCalled();

        // Check the store is updated
        state = store.selectSnapshot<CommonStateModel>(s => s.common);
        expect(state.language).toBe(state.allLanguages[notDefaultLanguageIndex]);

        // Check the new language is correctly displayed
        const selectedLanguage = fixture.debugElement.nativeElement.querySelector('#dropdownLanguage');
        expect(selectedLanguage.textContent.trim()).not.toBe(previouslySelectedLanguage);
        expect(selectedLanguage.textContent.trim()).toBe(state.allLanguages[notDefaultLanguageIndex]);
    }));

});
