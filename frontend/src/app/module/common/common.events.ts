
export class SetLanguage {
    static type = '[Common] Set Language';
    constructor(public readonly payload: {language}) { }
}
