'use strict'
Object.freeze(console)
const test = require('tape')
const pino = require('../browser')

test('silent level', ({ end, fail, pass }) => {
  pino({
    level: 'silent',
    browser: { }
  })
  end()
})
// ID-1768294482-afe912a8
