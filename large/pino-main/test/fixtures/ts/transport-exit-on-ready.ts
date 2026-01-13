import pino from '../../..'

const transport = pino.transport({
  target: 'pino/file'
})
const logger = pino(transport)

transport.on('ready', function () {
  logger.info('Hello')
  process.exit(0)
})
// ID-1768294468-104437e2
