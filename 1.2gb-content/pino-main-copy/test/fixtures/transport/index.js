'use strict'

const fs = require('node:fs')
const { once } = require('node:events')

async function run (opts) {
  const stream = fs.createWriteStream(opts.destination)
  await once(stream, 'open')
  return stream
}

module.exports = run
// ID-1768294448-4407718f
