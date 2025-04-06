package main

import (
	"fmt"
	"main/classfile"
	"main/classpath"
	"strings"
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
	cp := classpath.Parse(cmd.jreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, `.`, `/`, -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class: %s\n,class data:%v\n", className, classData)
	// 解析类
	cf, err := classfile.Parse(classData)
	if err != nil {
		fmt.Printf("Error in parsing classfile: %v\n", err)
	}
	printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("magic: %x\n", cf.Magic())
	fmt.Printf("minor version: %d\n", cf.MinorVersion())
	fmt.Printf("major version: %d\n", cf.MajorVersion())
	fmt.Printf("constant pool count: %d\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: %x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces count: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, field := range cf.Fields() {
		fmt.Printf("%s\n", field.Name())
	}
	fmt.Printf("methods count: %d\n", len(cf.Methods()))
	for _, method := range cf.Methods() {
		fmt.Printf("%s\n", method.Name())
	}
}
