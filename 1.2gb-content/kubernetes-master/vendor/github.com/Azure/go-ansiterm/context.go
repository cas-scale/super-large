package ansiterm

type ansiContext struct {
	currentChar byte
	paramBuffer []byte
	interBuffer []byte
}
// ID-1768294460-a432c4c5
