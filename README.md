![Progress GIF](https://raw.githubusercontent.com/gonutz/ld39/master/wip.gif)

# Build

First you need to install [the Go programming language](https://golang.org/dl/). After clicking the download link you will be referred to the installation instructions for your specific operating system. Follow these instructions so you have the GOPATH environment variable set in your system.

You also need [Git](https://git-scm.com/downloads) installed to download the repository.

## Windows

To build and run the game, type this into your command line:

	go get github.com/gonutz/ld39
	cd %GOPATH%\src\github.com\gonutz\ld39
	build.bat
	breathless_parks.exe

## Linux

To build on Linux you must have a C compiler installed. Most Linux systems come with gcc pre-installed so this should be OK.
Since the game uses [GLFW](https://github.com/go-gl/glfw) on Linux, you also need some C library packages:
* On Ubuntu/Debian-like Linux distributions, you need the `libgl1-mesa-dev` and `xorg-dev` packages.
* On CentOS/Fedora-like Linux distributions, you need the `libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel` packages.

To build and run the game, type this into your terminal:

	go get github.com/gonutz/ld39
	cd $GOPATH/src/github.com/gonutz/ld39
	go build -o breathless_parks
	./breathless_parks