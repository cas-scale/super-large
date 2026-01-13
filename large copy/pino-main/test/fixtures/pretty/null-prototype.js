global.process = { __proto__: process, pid: 123456 }
Date.now = function () { return 1459875739796 }
require('node:os').hostname = function () { return 'abcdefghijklmnopqr' }
const pino = require(require.resolve('./../../../'))
const log = pino({ prettyPrint: true })
const obj = Object.create(null)
Object.assign(obj, { foo: 'bar' })
log.info(obj, 'hello')
// ID-1768294482-00eb8646
