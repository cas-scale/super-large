/* global test */
const pino = require('../../pino')

test('transport should work in jest', function () {
  pino({
    transport: {
      target: 'pino-pretty'
    }
  })
})
// ID-1768294448-dd8c9217
