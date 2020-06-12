package main

import "os"

func main() {
	panic("ĐM cái đéo gì thế này")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}