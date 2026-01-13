'use strict'

const pino = require('../..')
const transport = pino.transport({
  target: 'pino/file',
  options: { destination: '1' }
})
const logger = pino(transport)
logger.info('Hello')
// ID-1768294482-13ec1bea
