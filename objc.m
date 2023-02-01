#import <Foundation/Foundation.h>
#import "objc.h"

void helloWorld()
{
	NSLog(@"Hello, World! \n");
}

char *getHelloWorld()
{
	NSString *helloStr = @"Hello, World! \n";
	return strdup([helloStr UTF8String]);
}

void print(const char *str)
{
	NSLog(@"%s", str);
}
