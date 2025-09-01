package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/grokify/go-qti/assessmentresult"
)

var (
	i = flag.String("r", "", "XML File for QTI Item")
)

func main() {
	flag.Parse()
	var inf *os.File
	var err error

	if *i != "" {
		inf, err = os.Open(*i)
		if err != nil {
			panic(err)
		}
	}

	infd, err := io.ReadAll(inf)
	if err != nil {
		panic(err)
	}

	var aResult assessmentresult.AssessmentResult
	err = xml.Unmarshal(infd, &aResult)
	if err != nil {
		panic(err)
	}

	rs, err := xml.MarshalIndent(aResult, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%s\n\n", rs)

}
