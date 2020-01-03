package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type params struct {
	*flag.FlagSet `json:"-"`

	dir  string
	port string

	help bool
}

func main() {

	f := &params{}
	f.FlagSet = flag.NewFlagSet("static server Global Params", flag.ContinueOnError)

	f.StringVar(&f.dir, "dir", "", "dir")
	f.StringVar(&f.port, "port", "8080", "port")

	f.BoolVar(&f.help, "h", false, "help")

	err := f.Parse(os.Args[1:])
	if err != nil {
		f.Usage()
		os.Exit(0)
	}

	if f.dir == "" {
		f.Usage()
		return
	}

	http.Handle("/", http.FileServer(http.Dir(f.dir)))
	fmt.Println("static server run ", f.port, "  dir:", f.dir)
	http.ListenAndServe(":"+f.port, nil)

}
