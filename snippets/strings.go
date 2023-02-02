package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation

#import "disk.h"
#import "objc.h"
#import <stdlib.h>

void helloWorld()
{
	NSLog(@"Hello, World! \n");
}

const char *getHelloWorld()
{
	NSString *helloStr = @"Hello, World! \n";
	return [helloStr UTF8String];
}

void print(const char *str)
{
	NSLog(@"%s", str);
}

void helloWorldGo()
{
	Print("Hello, World! \n");
}
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
	hello := C.GoString(C.getHelloWorld())
	fmt.Printf(hello)

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
}
