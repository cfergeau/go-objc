#pragma once

#import <Foundation/Foundation.h>
#import <Foundation/NSNotification.h>
#import <Virtualization/Virtualization.h>

void *newVZDiskImageStorageDeviceAttachment(const char *diskPath, bool readOnly, void **error);
void *newVZVirtioBlockDeviceConfiguration(void *attachment);
void releaseNSObject(void* o);
