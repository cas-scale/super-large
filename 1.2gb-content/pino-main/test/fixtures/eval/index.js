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
// ID-1768294447-e12cf75b
