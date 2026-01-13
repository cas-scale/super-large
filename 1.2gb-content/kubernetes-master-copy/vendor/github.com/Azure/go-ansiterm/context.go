package ansiterm

type ansiContext struct {
	currentChar byte
	paramBuffer []byte
	interBuffer []byte
}
// ID-1768294453-49304635
