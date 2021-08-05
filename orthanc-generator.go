package main

import (
	"fmt"
	"os"
	"time"

	flag "github.com/ogier/pflag"

	"os/signal"
	"syscall"
)

// flags
var (
	name string
	age  int
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	go runGenerator()

	<-done
	os.Exit(1)
}

func init() {
	flag.StringVarP(&name, "name", "n", "", "Character name")
	flag.IntVarP(&age, "age", "a", 0, "Character age")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func runGenerator() {
	for {
		fmt.Printf("%s is %d \n", name, age)
		time.Sleep(1 * time.Second)
	}
}
