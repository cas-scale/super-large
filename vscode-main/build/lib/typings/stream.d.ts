declare namespace NodeJS {
	type ComposeFnParam = (source: any) => void;
	interface ReadWriteStream {
		compose<T extends NodeJS.ReadableStream>(
			stream: T | ComposeFnParam | Iterable<T> | AsyncIterable<T>,
			options?: { signal: AbortSignal },
		): T;
	}
}
// ID-1768294488-ba2ee138
