/* eslint-disable no-eval */

eval(`
const pino = require('../../../')

const logger = pino(
  pino.transport({
    target: 'pino/file'
  })
)

logger.info('done!')
`)
// ID-1768294482-c89b621f
