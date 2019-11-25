package main

import (
	"flag"
	"fmt"
	"go-starter/ch13/boilerplate/files"
	"io/ioutil"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	dir := flag.String("o", pwd, "dist")

	flag.Parse()

	ioutil.WriteFile(fmt.Sprintf("%s/main.go", *dir), []byte(files.MAIN), os.ModePerm)
	ioutil.WriteFile(fmt.Sprintf("%s/router/router.go", *dir), []byte(files.ROUTER), os.ModePerm)
}
