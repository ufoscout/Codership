import { Injectable } from '@angular/core';
import * as obj from '../utils/object.utils';

@Injectable()
export class StorageService {

    fromOrDefault<T>(key: string, defaultValue: T): T {
        const value = localStorage.getItem(key);
        if (!obj.exists(value)) {
            return defaultValue;
        }
        return JSON.parse(value);
    }

    store<T>(key: string, value: T): void {
        localStorage.setItem(key, JSON.stringify(value));
    }

    remove<T>(key: string): void {
        localStorage.removeItem(key);
    }

    clear(): void {
        localStorage.clear();
    }
}
