package main

import (
	"encoding/hex"
	"flag"
	"github.com/jclab-joseph/go-winexetag"
	"log"
	"os"
)

func main() {
	var input string
	var output string
	var newtag string
	flag.StringVar(&input, "i", "", "")
	flag.StringVar(&output, "o", "", "")
	flag.StringVar(&newtag, "tag", "", "")
	flag.Parse()

	fin, err := os.Open(input)
	defer fin.Close()
	if err != nil {
		log.Fatalln(err)
	}
	bin, err := winexetag.NewPE32Binary(fin)
	if err != nil {
		log.Fatalln(err)
	}
	tag, err := bin.GetTag()
	if err != nil {
		log.Println("gettag: ", err)
	} else {
		log.Println("TAG: " + hex.EncodeToString(tag))
	}

	fout, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer fout.Close()
	err = bin.SetTag(fout, []byte(newtag))
	if err != nil {
		log.Fatalln(err)
	}
}
