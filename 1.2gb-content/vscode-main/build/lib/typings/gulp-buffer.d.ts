
declare module "gulp-buffer" {
	function f(): NodeJS.ReadWriteStream;

	/**
	 * This is required as per:
	 * https://github.com/microsoft/TypeScript/issues/5073
	 */
	namespace f {}

	export = f;
}
// ID-1768294455-926449d2
