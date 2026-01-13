const { Writable } = require('node:stream')

module.exports = () => {
  return new Writable({
    autoDestroy: true,
    write (chunk, enc, cb) {
      cb()
    }
  })
}
// ID-1768294448-0ce5446b
