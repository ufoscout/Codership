import { TestBed } from '@angular/core/testing';
import { StorageService } from './storage.service';

describe(`[Shared] StorageService`, () => {

    const KEY_1 = 'key1-' + new Date().getTime();
    const KEY_2 = 'key2-' + new Date().getTime();
    let storage: StorageService;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [],
            providers: [StorageService],
        });
        storage = TestBed.get(StorageService);
        storage.clear();
    });

    it('should persit a string', () => {

        const source = 'a string value';

        storage.store(KEY_1, source);

        expect(
            storage.fromOrDefault(KEY_1, '')
        )
            .toEqual(source);
    });

    it('should persit based on the key', () => {

        const source1 = 'a string value 1';
        const source2 = 'a string value 2';

        storage.store(KEY_1, source1);
        storage.store(KEY_2, source2);

        expect(
            storage.fromOrDefault(KEY_1, '')
        )
            .toEqual(source1);

        expect(
            storage.fromOrDefault(KEY_2, '')
        )
            .toEqual(source2);
    });

    it('should persit a number', () => {

        const source = 123456;

        storage.store(KEY_1, source);

        expect(
            storage.fromOrDefault(KEY_1, 0)
        )
            .toEqual(source);
    });

    it('should return the default value', () => {
        const defaultValue = 2134545;
        expect(
            storage.fromOrDefault(KEY_1, defaultValue)
        )
            .toEqual(defaultValue);
    });

    it('should persit an object', () => {

        const source = {
            age: 123456,
            name: 'Arthur Dent',
        };

        storage.store(KEY_1, source);

        expect(
            storage.fromOrDefault(KEY_1, {})
        )
            .toEqual(source);
    });

    it('should remove based on key', () => {

        const source = 'a string value';
        storage.store(KEY_1, source);
        storage.remove(KEY_1);

        expect(
            storage.fromOrDefault(KEY_1, '')
        )
            .toEqual('');
    });

    it('should clear the stored values when asked for', () => {
        storage.store(KEY_1, '');
        storage.store(KEY_2, '');

        storage.clear();

        expect(
            storage.fromOrDefault(KEY_1, '')
        )
            .toEqual('');

        expect(
            storage.fromOrDefault(KEY_2, '')
        )
            .toEqual('');
    });

});
