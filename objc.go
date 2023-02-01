package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation
#import <Foundation/Foundation.h>

void helloWorld()
{
	NSLog(@"Hello, World! \n");
}

char *getHelloWorld()
{
	NSString *helloStr = @"Hello, World! \n";
	return strdup([helloStr UTF8String]);
}

void print(const char *str)
{
	NSLog(@"%s", str);
}
*/
import "C"

// import "C" has to be right after the code block, and no blank line!

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
