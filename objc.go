package main

import "C"

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c
#cgo darwin LDFLAGS: -lobjc -framework Foundation
#import <Foundation/Foundation.h>

void helloWorld()
{
	NSLog(@"Hello, World! \n");
}
*/

func main() {
	// Call objc function
	C.helloWorld()
}
