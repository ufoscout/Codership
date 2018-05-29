import * as _ from 'lodash';

/**
 * Returns true if the object is not undefined and not null.
 */
export function exists(obj: any): boolean {
  return (!_.isUndefined(obj) && !_.isNull(obj));
}

/**
 * Returns value if exists, otherwise defaultValue.
 */
export function getOrDefault<T>(value: T, defaultValue: T): T {
  return exists(value) ? value : defaultValue;
}
