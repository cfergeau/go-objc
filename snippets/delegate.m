@implementation VZVirtioSocketListenerDelegateImpl {
    void *_cgoHandler;
}

- (instancetype)initWithHandler:(void *)cgoHandler
{
    self = [super init];
    _cgoHandler = cgoHandler;
    return self;
}

- (BOOL)listener:(VZVirtioSocketListener *)listener shouldAcceptNewConnection:(VZVirtioSocketConnection *)connection fromSocketDevice:(VZVirtioSocketDevice *)socketDevice;
{
    return (BOOL)shouldAcceptNewConnectionHandler(_cgoHandler, connection, socketDevice);
}
@end

void *newVZVirtioSocketListener(void *cgoHandlerPtr)
{
    VZVirtioSocketListener *ret = [[VZVirtioSocketListener alloc] init];
    [ret setDelegate:[[VZVirtioSocketListenerDelegateImpl alloc] initWithHandler:cgoHandlerPtr]];
    return ret;
}
