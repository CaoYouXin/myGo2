package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// PKG xxx
type PKG struct {
	Deps []string `json:"Deps"`
}

func main() {
	cmd := exec.Command("go", "list", "-json", os.Args[1])

	if outs, err := cmd.Output(); err != nil {
		fmt.Fprintf(os.Stderr, "pkg: %v\n", err)
		os.Exit(1)
	} else {
		var pkg PKG
		err = json.Unmarshal(outs, &pkg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pkg: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(pkg.Deps)
	}
}
