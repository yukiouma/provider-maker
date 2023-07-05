package main

import (
	"flag"
	"log"

	"github.com/yukiouma/provider-maker/internal/providermaker"
)

func main() {
	in := flag.String("in", ".", "path of project")
	out := flag.String("out", ".", "path of generated provider")
	flag.Parse()
	iter, err := providermaker.NewDirIterator(*in, providermaker.DefaultFilters...)
	if err != nil {
		log.Fatal(err)
	}
	p := providermaker.NewProvider()
	for f, ok := iter.Next(); ok; {
		if err := p.Analyse(f); err != nil {
			log.Fatalf("error: failed to analyse file [%s], because: %v", f, err)
		}
		f, ok = iter.Next()
	}
	if err := p.Generate(*out); err != nil {
		log.Fatalf("error: failed to generate provider because: %v", err)
	}
}
