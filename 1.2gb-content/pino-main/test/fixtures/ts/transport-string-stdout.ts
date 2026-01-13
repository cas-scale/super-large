import pino from '../../..'

const transport = pino.transport({
  target: 'pino/file',
  options: { destination: '1' }
})
const logger = pino(transport)
logger.info('Hello')
// ID-1768294447-5c74f940
