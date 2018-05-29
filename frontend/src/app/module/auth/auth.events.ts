
export class SetAuthData {
    static type = '[Auth] Set Auth data';
    constructor(public readonly payload: {username: string, token: string}) { }
}
