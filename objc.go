package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation

#import "disk.h"
#import "objc.h"
#import "stdlib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export Print
func Print(cstr *C.char) {
	gostr := C.GoString(cstr)
	fmt.Print(gostr)
}

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

	// Call objc function which will call the golang Print function
	C.helloWorldGo()

	// More advanced test using virtualization framework methods
	attachment, err := NewDiskImageStorageDeviceAttachment("/dev/zero", true)
	if err != nil {
		panic(err.Error())
	}
	config, err := NewVirtioBlockDeviceConfiguration(attachment)
	if err != nil {
		panic(err.Error())
	}
	C.releaseNSObject(attachment.pointer)
	C.releaseNSObject(config.pointer)

        fmt.Println(vfkitCmdline())
}
