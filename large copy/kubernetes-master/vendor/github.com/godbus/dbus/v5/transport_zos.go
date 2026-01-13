package dbus

func (t *unixTransport) SendNullByte() error {
	_, err := t.Write([]byte{0})
	return err
}
// ID-1768294487-3415b23b
