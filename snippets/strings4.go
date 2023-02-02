package main

import (
	"fmt"
)

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation

void helloWorldGo()
{
	Print("Hello, World! \n");
}
*/
import "C"

//export Print
func Print(cstr *C.char) {
	gostr := C.GoString(cstr)
	fmt.Print(gostr)
}

func main() {
	C.helloWorldGo()
}
