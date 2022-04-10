package main

/* "go mod init square" makes a module square and /square points to the directory square where code is */
import (
	"square/square"
)

func main() {
	//square.SquareA(10, 10)
	square.SquareB(10, 10)
}
