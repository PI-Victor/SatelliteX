#!/usr/bin/env bash

# source of inspiration was Docker's make file and install scripts
# https://github.com/docker/docker/blob/master/project/make.sh
# only tested on Centos 7 so far

ARG=$1
GOBIN="/usr/bin/go"
# all main entry points that need to be build will be 
# nested in the cmd folder
# CMDDIR="./cmd" 
# application workspace
APPWS="application/replaceme"
GOBUILD=$GOPATH/build
# what do we want to execut when there's no arg passed?
# in the current context it's bloated 
# keep for now remember to simplify later
DEFAULT_BUNDLES=( 
    install
    clean
    build
    gotest 
)
# gobuild for SU is
GOBUILDSU="/home/vectra/projects/go_projects/build"

# if it's just a build, add the build dir to GOPATH
# remember NOT to add $GOPATH/bin to $GOPATH, but rather
# add $GOPATH/build to $PATH

cleanall() {
    # clean all files from the build dir
    # clean all files from the pkg source
    # we're gonna recompile everything
    if [ -d $GOBUILD ]; then
	rm -f $GOBUILD/*
	echo "Removed directory content of $GOBUILD"
    fi
    # assume we're running on 64bits
    # remove workspace objects before recompiling them
    rm -rf $GOPATH/pkg/linux_amd64/$APPWS
}

compile() {
    # compile all main binaries to build dir
    for CMDDIR in "$(ls ./cmd)"; do
	echo "Building... "
	go build -o $GOBUILD/$CMDDIR ./cmd/$CMDDIR
	echo "Compiled $CMDDIR to $GOBUILD/$CMDDIR"
    done
}

install() {
    # this also installs the binary as a soft link to
    # /usr/bin
    for APPBIN in "$(ls $GOBUILDSU)"; do
	echo "Trying to create symlink for $GOBUILD/$APPBIN in /usr/bin"
	ln -s $GOBUILDSU/$APPBIN /usr/bin/$APPBIN
	if [ "$?" -eq 0 ];then
	    echo "Created symlink"
	fi
    done
}

makedir() {
    # check if main build dir is in the go path
    if [ -z "$GOBUILD" ]; then
	echo "Creating directory $GOBUILD"
	mkdir -p $GOBUILD
	# but add everything to $PATH so that it desrupts the
	# current GO env of the host as less as possilbe
	# and make compiled bin available
	export PATH=$PATH:$GOBUILD # THIS IS NOT A GOOD IDEA, CHANGE THIS WHENEVER
    fi
}

main() {
    # exit if we don't find the go bin
    if [ ! -f "$GOBIN" ] ; then
	echo "Go binary not found in" $GOBIN
	exit
    fi
    
    # must refactor logic
    case "$1" in
	clean)
            cleanall
	    ;;
	build)
	    cleanall
	    makedir
	    compile
	    ;;
	install)
	    install
	    ;;
	gotest)
	    echo "Testing not implemented yet"
	    ;;
    esac
}

main "$@"
