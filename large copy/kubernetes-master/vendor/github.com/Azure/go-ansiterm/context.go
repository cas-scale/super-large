package ansiterm

type ansiContext struct {
	currentChar byte
	paramBuffer []byte
	interBuffer []byte
}
// ID-1768294487-c81ed65b
