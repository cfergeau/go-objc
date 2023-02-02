package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation

#import "objc.h"

void helloWorld()
{
	NSLog(@"Hello, World! \n");
}
*/
import "C"

func main() {
	// Call objc function
	C.helloWorld()
}
