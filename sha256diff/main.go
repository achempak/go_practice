package main

// Count number of differing bits in two SHA256 digests

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

var lookup [256]byte

func init() {
	for i := range lookup {
		lookup[i] = lookup[i/2] + byte(i&1)
	}
}

func WriteToFile() {
	f, err := os.Create("lookup.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i, val := range lookup {
		fmt.Fprintf(w, "%d\t %d\n", i, val)
	}
	w.Flush()
}

func Diff(x [32]byte, y [32]byte) int {
	var result byte
	for i := 0; i < 32; i++ {
		// Get population count of the number of differing bits in each
		// corresponding byte of x and y.
		result += lookup[x[i]^y[i]]
		fmt.Printf("%d\t%08b\t%08b\t%08b\n", lookup[x[i]^y[i]], x[i], y[i], x[i]^y[i])
	}
	return int(result)
}

func main() {
	var arg1, difference int
	var sha1, sha2 [32]byte
	WriteToFile()
	if os.Args[1] == "--" {
		arg1 = 2
	} else {
		arg1 = 1
	}
	sha1 = sha256.Sum256([]byte(os.Args[arg1]))
	sha2 = sha256.Sum256([]byte(os.Args[arg1+1]))
	fmt.Printf("Hash of %s is %x\n", os.Args[arg1], sha1)
	fmt.Printf("Hash of %s is %x\n", os.Args[arg1+1], sha2)
	difference = Diff(sha1, sha2)
	if difference >= 0 {
		fmt.Printf("%d number of bits differ", difference)
	} else {
		fmt.Printf("%d number of bits differ", -difference)
	}
}
