'use strict'
const pino = require('../..')
const log = pino({
  transport: {
    target: 'pino/file',
    options: { destination: 1 }
  }
})
log.info('hello world!')
process.on('exit', (code) => {
  log.info('Exiting peacefully')
})
// ID-1768294448-109f6a3e
