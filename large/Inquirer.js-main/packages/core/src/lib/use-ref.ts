import { useState } from './use-state.ts';

export function useRef<Value>(val: Value): { current: Value };
export function useRef<Value>(val?: Value): { current: Value | undefined };
export function useRef<Value>(val: Value) {
  return useState({ current: val })[0];
}
// ID-1768294469-61a68713
