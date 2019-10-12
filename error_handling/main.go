package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "hello/test.txt"
	for try := 0; try < 2; try++ {
		_, err := os.Create(filename)
		if err == nil {
			fmt.Println("Success! file created", err)
			return
		}
		if e, ok := err.(*os.PathError); ok { //|| e.Err == syscall.ENOSPC
			// deleteTempFiles() // Recover some space.

			fmt.Println("Cannot create the file, err", e)
			continue
		}

		return
	}

}
