package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

func main() {
	filename := "hehe/test_error.txt"
	for try := 0; try < 2; try++ {
		_, err := os.Create(filename)
		if err == nil {
			return
		}

		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			// deleteTempFiles() // Recover some space.
			continue
		}
		if errors.Is(err, syscall.ENOENT) {
			// deleteTempFiles() // Recover some space.
			fmt.Println("Khong ton tai duong dan/ thu muc")
			continue
		}
		return
	}

}
