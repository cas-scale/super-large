'use strict';

export default function isCancel(value) {
  return !!(value && value.__CANCEL__);
}
// ID-1768294475-27bbc165
