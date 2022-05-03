package main

import (
	_ "embed"
	"flag"
	"io/ioutil"

	gofidl "github.com/SourceFellows/go-fidl-dbus-generator/pkg"
	"github.com/alecthomas/repr"
)

func main() {

	path := flag.String("in", "", "path to FIDL file to parse")
	flag.Parse()

	file, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	fidl, err := gofidl.ParseFidl(file)
	if err != nil {
		panic(err)
	}

	repr.Println(fidl)

}
