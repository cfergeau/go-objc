package main

/*
#cgo darwin CFLAGS: -mmacosx-version-min=11 -x objective-c -fno-objc-arc
#cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Virtualization

#import "disk.h"

void *newVZVirtioBlockDeviceConfiguration(void *attachment)
{
    return [[VZVirtioBlockDeviceConfiguration alloc] initWithAttachment:(VZStorageDeviceAttachment *)attachment];
}

void releaseNSObject(void* o)
{
    [(NSObject*)o release];
}
*/
import "C"
import (
	"unsafe"
)

type VirtioBlockDeviceConfiguration struct {
	pointer unsafe.Pointer
}

func NewVirtioBlockDeviceConfiguration(attachment *DiskImageStorageDeviceAttachment) (*VirtioBlockDeviceConfiguration, error) {
	objcConfig := C.newVZVirtioBlockDeviceConfiguration(attachment.pointer)

	config := &VirtioBlockDeviceConfiguration{
		pointer: objcConfig,
	}
	return config, nil
}
