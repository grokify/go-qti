package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/grokify/go-qti/assessmentitem"
	"github.com/grokify/go-qti/assessmenttest"
)

var (
	i = flag.String("i", "", "XML File for QTI Item")
	t = flag.String("t", "", "XML File for QTI Item")
)

func main() {
	fmt.Println("Running main tester")
	flag.Parse()
	var inf *os.File
	var err error

	if *i != "" {
		inf, err = os.Open(*i)
		if err != nil {
			panic(err)
		}
	} else if *t != "" {
		inf, err = os.Open(*t)
		if err != nil {
			panic(err)
		}
	}

	infd, err := io.ReadAll(inf)
	if err != nil {
		panic(err)
	}

	if *i != "" {
		it, err := assessmentitem.NewAssessmentItem(infd)
		if err != nil {
			panic(err)
		}
		nb, err := it.ItemBody.UpdateImagePath("MYITEM")
		it.ItemBody.Xml = string(nb)

		newB, err := it.ItemBody.ToWebComponent()

		wb := assessmentitem.ItemWebComponent{Xml: string(newB), Identifier: it.Identifier, TimeDependent: it.TimeDependent, Label: it.Label, Adaptive: it.Adaptive, Title: it.Title}

		x, e := xml.Marshal(wb)
		if e != nil {
			panic(e)
		}

		fmt.Printf("%s", x)
	} else if *t != "" {
		ia, err := assessmenttest.NewAssessmentTest(infd)
		if err != nil {
			panic(err)
		}

		xmlOut, err := xml.Marshal(ia)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", xmlOut)

	} else {
		panic("No File")
	}
}
