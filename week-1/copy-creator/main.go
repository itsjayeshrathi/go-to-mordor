package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(fileDestination, srcDestination string) (int64, error) {
	src, err := os.Open(fileDestination)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	dst, err := os.Create(srcDestination)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	fmt.Println("©️©️©️ WELCOME TO COPY CREATOR ©️©️©️")

	var fileDestination, srcDestination string

	for {
		fmt.Println("Enter the file destination that you would like to copy.")
		fmt.Scanln(&fileDestination)
		fmt.Println("Enter the source destination where you would like to make copy.")
		fmt.Scanln(&srcDestination)

		fileDestination = os.ExpandEnv(fileDestination)
		srcDestination = os.ExpandEnv(srcDestination)

		_, err := copyFile(fileDestination, srcDestination)

		if err == nil {
			fmt.Println("file copied successfully")
		} else {
			fmt.Println("there was some error", err)
		}
	}

}
