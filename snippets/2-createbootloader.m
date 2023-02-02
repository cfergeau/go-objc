func createBootLoader(kernelURL: URL, initialRamdiskURL: URL) -> VZBootLoader {
    let bootLoader = VZLinuxBootLoader(kernelURL: kernelURL)
    bootLoader.initialRamdiskURL = initialRamdiskURL


    let kernelCommandLineArguments = [
        // Use the first virtio console device as system console.
        "console=hvc0",
        // Stop in the initial ramdisk before attempting to transition to
        // the root file system.
        "rd.break=initqueue"
    ]


    bootLoader.commandLine = kernelCommandLineArguments.joined(separator: " ")


    return bootLoader
}
