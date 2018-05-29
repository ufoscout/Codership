import * as obj from './object.utils';

describe('[Shared] Object utils', () => {

    it('should not exist if undefined', () => {
        expect(
            obj.exists(undefined)
        )
            .toBeFalsy();
    });

    it('should not exist if null', () => {
        expect(
            obj.exists(null)
        )
            .toBeFalsy();
    });

    it('should exist if not undefined and not null', () => {
        expect(
            obj.exists('undefined')
        )
            .toBeTruthy();
    });

    it('should return defaul if value is undefined', () => {
        expect(
            obj.getOrDefault(undefined, 'default')
        )
            .toBe('default');
    });

    it('should return defaul if value is null', () => {
        expect(
            obj.getOrDefault(null, 'default')
        )
            .toBe('default');
    });

    it('should return value if is not undefined and not null', () => {
        expect(
            obj.getOrDefault('value', 'default')
        )
            .toBe('value');
    });
});
