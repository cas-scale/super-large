declare module 'ternary-stream' {
	import File = require('vinyl');
	function f(check: (f: File) => boolean, onTrue: NodeJS.ReadWriteStream, opnFalse?: NodeJS.ReadWriteStream): NodeJS.ReadWriteStream;

	/**
	 * This is required as per:
	 * https://github.com/microsoft/TypeScript/issues/5073
	 */
	namespace f {}

	export = f;
}
// ID-1768294455-95a5c8c4
