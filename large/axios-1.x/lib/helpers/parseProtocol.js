'use strict';

export default function parseProtocol(url) {
  const match = /^([-+\w]{1,25})(:?\/\/|:)/.exec(url);
  return match && match[1] || '';
}
// ID-1768294468-f3a28c8f
