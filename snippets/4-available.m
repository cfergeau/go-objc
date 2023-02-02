void *newVZSingleDirectoryShare(void *sharedDirectory)
{
    if (@available(macOS 12, *)) {
        return [[VZSingleDirectoryShare alloc] initWithDirectory:(VZSharedDirectory *)sharedDirectory];
    }

    RAISE_UNSUPPORTED_MACOS_EXCEPTION();
}
