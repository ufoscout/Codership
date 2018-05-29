import { Store, NgxsModule } from '@ngxs/store';
import { async, TestBed } from '@angular/core/testing';
import { CommonState, CommonStateModel, defaultLanguage } from './common.state';
import { SetLanguage } from './common.events';
import { TranslateService } from '@ngx-translate/core';

class MockTranslateService {

    defaultLang: String;
    lang: String;

    setDefaultLang(language: String) {
        this.defaultLang = language;
    }
    use(language: String): any {
        this.lang = language;
    }
  }

describe('[Common] Common State', () => {
    let store: Store;
    let translateService: MockTranslateService;

    beforeEach(async(() => {
        TestBed.configureTestingModule({
            imports: [NgxsModule.forRoot([CommonState])],
            providers: [
                {provide: TranslateService, useClass: MockTranslateService}
              ]
        }).compileComponents();
        store = TestBed.get(Store);
        translateService = TestBed.get(TranslateService);
    }));

    it('it should initialize the default languages', () => {
        expect(translateService.defaultLang).toBe(defaultLanguage);
        expect(translateService.lang).toBe(defaultLanguage);
    });

    it('it should set the current language', () => {
        const newLanguage = 'fr';
        store.dispatch(new SetLanguage({ language: newLanguage }));
        store.selectOnce(CommonState).subscribe((state: CommonStateModel) => {
            expect(state.language).toBe(newLanguage);
            expect(state.defaultLanguage).toBe(defaultLanguage);
        });
        expect(translateService.defaultLang).toBe(defaultLanguage);
        expect(translateService.lang).toBe(newLanguage);
    });

});
