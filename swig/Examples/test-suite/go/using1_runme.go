package main

import "swigtests/using1"

func main() {
	if using1.Spam(37) != 37 {
		panic(0)
	}
}
