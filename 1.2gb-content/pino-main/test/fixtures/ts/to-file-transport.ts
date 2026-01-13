import * as fs from 'node:fs'
import { once } from 'node:events'

async function run (opts: { destination?: fs.PathLike }): Promise<fs.WriteStream> {
  if (!opts.destination) throw new Error('kaboom')
  const stream = fs.createWriteStream(opts.destination, { encoding: 'utf8' })
  await once(stream, 'open')
  return stream
}

export default run
// ID-1768294447-970b6ca2
