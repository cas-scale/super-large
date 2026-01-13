import { PathLike } from 'node:fs'

export declare function watchFileCreated(filename: PathLike): Promise<void>
export declare function watchForWrite(filename: PathLike, testString: string): Promise<void>
// ID-1768294447-0423095e
