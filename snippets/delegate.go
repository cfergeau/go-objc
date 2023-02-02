
//export shouldAcceptNewConnectionHandler
func shouldAcceptNewConnectionHandler(cgoHandlerPtr, connPtr, devicePtr unsafe.Pointer) C.bool {
	cgoHandler := *(*cgo.Handle)(cgoHandlerPtr)
	handler := cgoHandler.Value().(func(*VirtioSocketConnection, error))

	// see: startHandler
	conn, err := newVirtioSocketConnection(connPtr)
	go handler(conn, err)
	return (C.bool)(true)
}

func (v *VirtioSocketDevice) Listen(port uint32) (*VirtioSocketListener, error) {
	ch := make(chan connResults, 1) // should I increase more caps?

	handler := cgo.NewHandle(func(conn *VirtioSocketConnection, err error) {
		ch <- connResults{conn, err}
	})
	ptr := C.newVZVirtioSocketListener(
		unsafe.Pointer(&handler),
	)
