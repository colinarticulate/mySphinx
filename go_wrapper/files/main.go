package main

import (
	"fmt"
	"os"
	//"reflect"
	//"io"
	//"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./../data/kl_ay_m.jsgf")
	check(err)
	fmt.Print(string(dat))
	fmt.Printf("type: %T\n", string(dat))
	fmt.Printf("type: %T\n", dat)
	//
	//fmt.Printf("type: %T\n", byte(dat))

	// f, err := os.Open("./data/kl_ay_m.jsgf")

	// r := bufio.NewReader(f)
	// bcontents, err := r.ReadBytes(0x0)
	// check(err)

	// f.Close()
}
