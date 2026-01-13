'use strict';

export default function isCancel(value) {
  return !!(value && value.__CANCEL__);
}
// ID-1768294488-0c5c2fe2
