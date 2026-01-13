package ansiterm

type ansiContext struct {
	currentChar byte
	paramBuffer []byte
	interBuffer []byte
}
// ID-1768294494-f1b4263a
