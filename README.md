# ball-sim-go

A learning exercise for a golang project: Ball simulator.

I have tried to follow the recommendations:
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Project Layout](https://github.com/golang-standards/project-layout)

Visual representation is made using Hajime Hoshis [Ebiten Game Library](https://ebiten.org/)

Lastly, big thanks to my daughter who has made the images used for the graphics :).

<img src="assets/images/screenshot.png" alt="Ball Simulator Screenshot" width="80%" />

## Prerequisites

Ubuntu:
```
sudo apt install xorg-dev libgl1-mesa-dev
```

## Quick Start

The following assumes installed go binaries are placed in $PATH

To install programs and run each program:
```
./script/bsg.sh install
collision -nballs=50
bounce -nballs=200
```

## Developer Instruction

The script bsg.sh supports developers

```
./scripts/bsg.sh help

 bsg.sh --

	Script for the ball-sim-go program collection.

	Preparation;

	Install the following packages:
	Ubuntu
		sudo apt-get install xorg-dev libgl1-mesa-dev

 Commands;

	build [--clean]
		Compiles ball-sim-go programs

	install
		Installs ball-sim-go programs

	test
		Unit test the ball-sim-go programs

	format
		Lint and format check

	smoketest
		Execute build, test and format
```
