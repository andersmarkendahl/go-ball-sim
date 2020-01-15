# ball-sim-go

A learning exercise for a golang project: Ball simulator.

I have tried to follow the recommendations:
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Project Layout](https://github.com/golang-standards/project-layout)

Visual representation is made using Hajime Hoshis [Ebiten Game Library](https://ebiten.org/)

Lastly, big thanks to my daughter who has made the images used for the graphics :).

<img src="assets/screenshot.png" alt="Ball Simulator Screenshot" width="80%" />

## Build

```
go build -o gravity cmd/gravity/main.go
```

## Test

```
go test ./pkg/...
```
**Note** Only external packages are unit tested.

## Run

```
./gravity
```
