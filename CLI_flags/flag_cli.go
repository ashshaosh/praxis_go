package main

import (
	//"flag"
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"Name to greeting"`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish lang"`
}

//define vars to hold flags
// var name = flag.String("name", "World", "A name to say hello to")
// var spanish bool
// var help bool

// func init() {
// 	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
// 	flag.BoolVar(&spanish, "s", false, "Use Spanish lang")
// 	flag.BoolVar(&help, "help", false, "Show help")
// }

func main() {
	flags.Parse(&opts)
	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}

	// if help == true {
	// 	// Set non standart help for flag
	// 	flag.VisitAll(func(flag *flag.Flag) {
	// 		format := "\t-%s: %s (DEFAULT: '%s')\n"
	// 		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	// 	})
	// }

}
