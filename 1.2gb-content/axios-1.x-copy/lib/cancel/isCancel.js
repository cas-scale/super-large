'use strict';

export default function isCancel(value) {
  return !!(value && value.__CANCEL__);
}
// ID-1768294448-d5251039
