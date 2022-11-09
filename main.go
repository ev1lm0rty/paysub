package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	sub_file := "/tmp/p-subs.txt"

	// Parsing Command Line Args
	sub := flag.String("d", "sub", "subdomain")
	output := flag.String("o", "output", "output")
	flag.Parse()

	app := "grep"

	arg0 := "-E"
	arg1 := ".*" + "\\." + string(*sub) + "$"
	arg2 := sub_file

	cmd := exec.Command(app, arg0, arg1, arg2)

	stdout, err := cmd.Output()

	if err != nil {
		return
	}

	fmt.Print(string(stdout))
	f, _ := os.Create(*output)
	_, err2 := f.WriteString(string(stdout))
	_ = err2
}
