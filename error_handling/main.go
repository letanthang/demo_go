package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	filename := "hello/test.txt"
	for try := 0; try < 2; try++ {
		_, err := os.Create(filename)
		if err == nil {
			fmt.Println("Success! file created", err)
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			// deleteTempFiles() // Recover some space.

			fmt.Println("Cannot create the file: diskspace, err", e)
			continue
		}

		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
			// deleteTempFiles() // Recover some space.

			fmt.Println("Cannot create the file: no such path/directory, err", e)
			continue
		}

		return
	}

}
