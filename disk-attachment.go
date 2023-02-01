package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization

#import "disk.h"

void *newVZDiskImageStorageDeviceAttachment(const char *diskPath, bool readOnly, void **error)
{
    NSString *diskPathNSString = [NSString stringWithUTF8String:diskPath];
    NSURL *diskURL = [NSURL fileURLWithPath:diskPathNSString];
    return [[VZDiskImageStorageDeviceAttachment alloc]
        initWithURL:diskURL
           readOnly:(BOOL)readOnly
              error:(NSError *_Nullable *_Nullable)error];
}
*/
import "C"
import (
	"unsafe"
)

type DiskImageStorageDeviceAttachment struct {
	pointer unsafe.Pointer
}

func NewDiskImageStorageDeviceAttachment(diskPath string, readOnly bool) (*DiskImageStorageDeviceAttachment, error) {
	diskPathChar := C.CString(diskPath)
	defer C.free(unsafe.Pointer(diskPathChar))
	objcAttachment := C.newVZDiskImageStorageDeviceAttachment(diskPathChar, C.bool(readOnly), nil)

	return &DiskImageStorageDeviceAttachment{
		pointer: objcAttachment,
	}, nil
}
