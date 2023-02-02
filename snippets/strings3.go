package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation

void print(const char *str)
{
	NSLog(@"%s", str);
}
*/
import "C"

func main() {
	cstr := C.CString("Hello, World! \n")
	C.print(cstr)
	C.free(unsafe.Pointer(cstr))
}
