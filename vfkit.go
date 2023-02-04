package main

import (
    "strings"
    "github.com/crc-org/vfkit/pkg/config"
)

func vfkitCmdline() string {
    bootloader := config.NewEFIBootloader("/Users/teuf/efi-variable-store", true)
    vmConfig:= config.NewVirtualMachine(2, 4*1024*1024*1024, bootloader)
    
    disk, _ := config.VirtioBlkNew("/Users/teuf/vz-test.img")
    vmConfig.AddDevice(disk)
    
    serial, _ := config.VirtioSerialNew("/Users/teuf/console/log")
    vmConfig.AddDevice(serial)
    
    cmdline, _ :=  vmConfig.ToCmdLine()
    return strings.Join(cmdline, " ")
}
