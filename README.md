# ball-sim-go

A learning exercise for a golang project: Ball simulator.

I have tried to follow the recommendations:
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Project Layout](https://github.com/golang-standards/project-layout)

Visual representation is made using Hajime Hoshis [Ebiten Game Library](https://ebiten.org/)

Lastly, big thanks to my daughter who has made the images used for the graphics :).

<img src="assets/screenshot.png" alt="Ball Simulator Screenshot" width="80%" />

## Prerequisites

Ubuntu:
```
sudo apt install xorg-dev libgl1-mesa-dev
```

## Build

```
go build -o bounce cmd/bounce/main.go
go build -o collision cmd/collision/main.go
```

## Test

```
go test ./pkg/...
```
**Note** Only external packages are currently unit tested.

## Run

```
./bounce -nballs=10
```
or...
```
./collision -nballs=10
```
