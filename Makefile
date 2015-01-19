PHONY: all clean test build install
#use this in conjuction with cleaning binaries from the build dir
#and other stuff
all:
	install/install.sh
clean:
	install/install.sh clean
build:
	install/install.sh build
#use this just for a normal install
install:
	install/install.sh install
test:
	install/install.sh gotest
