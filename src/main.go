package main

import (
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1\n")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("starting JVM...\n")
	fmt.Printf("Classpath: %s, class: %s, args: %v\n", cmd.cpOption, cmd.class, cmd.args)
}
