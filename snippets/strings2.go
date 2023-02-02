package main

import (
	"fmt"
)

/*
#cgo ...

char *getHelloWorld()
{
	NSString *helloStr = @"Hello, World! \n";
	return strdup([helloStr UTF8String]);
}
*/
import "C"

func main() {
	helloc := C.getHelloWorld()
	hello := C.GoString(helloc)
	fmt.Printf(hello)
	C.free(unsafe.Pointer(helloc))
}
