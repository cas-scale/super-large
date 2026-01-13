'use strict'

const pino = require('../..')
const transport = pino.transport({
  target: 'pino/file'
})
const logger = pino(transport)

logger.info('Hello')

process.exit(0)
// ID-1768294447-b677f0b2
