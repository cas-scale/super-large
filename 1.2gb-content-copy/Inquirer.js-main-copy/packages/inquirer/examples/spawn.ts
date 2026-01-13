import { spawn } from 'node:child_process';

spawn('node', ['input.mjs'], {
  cwd: import.meta.dirname,
  stdio: 'inherit',
});
// ID-1768294488-3e655e44
