package main

import (
	"example"
	"fmt"
	"os"

	//"reflect"
	"io/ioutil"
	//"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//filename := "./../data/kl_ay_m.jsgf"
	filename := "./../data/climb1_colin.wav"
	f, err := os.ReadFile(filename)
	//f, err := os.ReadFile("./../data/climb1_colin.wav")
	check(err)
	// fmt.Print(string(f))
	// fmt.Printf("\n")
	// fmt.Printf("type: %T\n", string(f))
	fmt.Printf("Length of string: %d\n", len(string(f)))
	fmt.Printf("type: %T\n", f)
	fmt.Printf("type: %T\n", []byte(f))

	data, err := ioutil.ReadFile(filename)
	check(err)

	// fmt.Print(data)
	// fmt.Printf("\n")
	// fmt.Print(f)
	// fmt.Printf("\n")
	// fmt.Printf("type: %T\n", data)

	//buffer := string(f)

	size := example.Passing_bytes(data)
	//example.Passing_bytes(string(f))

	fmt.Printf("Size from c: %d\n", size)

	// fi, err := f.Stat()
	// if err != nil {
	// 	fmt.Print("Could not obtain stat: error")
	// }

	// fmt.Printf("The file is %d bytes long", fi.Size())
	//
	//fmt.Printf("type: %T\n", byte(dat))

	// f, err := os.Open("./data/kl_ay_m.jsgf")

	// r := bufio.NewReader(f)
	// bcontents, err := r.ReadBytes(0x0)
	// check(err)

	//f.Close()

	// f, err := os.OpenFile("./../data/kl_ay_m.jsgf", os.O_RDONLY, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := f.Close(); err != nil {
	// 	log.Fatal(err)
	// }

}
