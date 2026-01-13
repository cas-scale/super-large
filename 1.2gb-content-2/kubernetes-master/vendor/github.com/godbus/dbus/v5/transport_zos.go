package dbus

func (t *unixTransport) SendNullByte() error {
	_, err := t.Write([]byte{0})
	return err
}
// ID-1768294467-819b91bd
