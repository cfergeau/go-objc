package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation

#import "objc.h"
#import "stdlib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Call objc function
	C.helloWorld()

	// Call objc function returning a string, and print it from go
	helloc := C.getHelloWorld()
	hello := C.GoString(helloc)
	fmt.Printf(hello)
	C.free(unsafe.Pointer(helloc))

	// Call objc function with a string argument
	cstr := C.CString("Hello, World! \n")
	C.print(cstr)
	C.free(unsafe.Pointer(cstr))
}
