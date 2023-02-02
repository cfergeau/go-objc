let virtualMachine = VZVirtualMachine(configuration: configuration)


let delegate = Delegate()
virtualMachine.delegate = delegate


virtualMachine.start { (result) in
    if case let .failure(error) = result {
        print("Failed to start the virtual machine. \(error)")
        exit(EXIT_FAILURE)
    }
}


RunLoop.main.run(until: Date.distantFuture)
