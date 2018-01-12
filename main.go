//https://gist.github.com/icholy/5314423

package main

import "os"
import "fmt"
import (
	"crypto/sha256"

	"github.com/fatih/color"
)

func main() {
	color.Blue("hiihi")
	color.Green("day la green nay")
	color.Black("day la black ne")

	return
	var args = os.Args
	if len(args) == 1 {
		printHelp()
		return
	}

	if len(args) >= 2 {
		test1()
	}
}

func help() string {
	return "Help"
}

func printHelp() {
	fmt.Println(help())
}

func test1() {
	var enc = sha256.New()
	var bytes = []byte("a")
	var bytes1 = append(bytes, bytes...)

	fmt.Printf("%x\n", enc.Sum(bytes))
	fmt.Println("size: ", enc.Size())

	fmt.Printf("%x\n", sha256.Sum256(bytes))

	fmt.Printf("%x\n", enc.Sum(bytes))
	fmt.Println("size: ", enc.Size())

	fmt.Printf("%x\n", sha256.Sum256(bytes1))

	enc.Write(bytes)
	enc.Sum(nil)
	enc.Write(bytes)
	enc.Reset()

	enc.Write(bytes)
	enc.Sum(nil)
	enc.Write(bytes)
	fmt.Printf("%x\n", enc.Sum(nil))
}
