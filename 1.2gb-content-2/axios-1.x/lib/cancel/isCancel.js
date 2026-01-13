'use strict';

export default function isCancel(value) {
  return !!(value && value.__CANCEL__);
}
// ID-1768294461-6a5111c1
