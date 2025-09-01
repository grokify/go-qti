package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/grokify/go-qti/assessmentitem"
)

var (
	in  = flag.String("i", "./", "Input Directory of XML Files")
	out = flag.String("o", "./out", "Output Directory of MathML Files")
	h   = flag.Bool("h", false, "Print Help")
)

func main() {
	flag.Parse()

	info, err := os.Stat(*in)
	if err != nil {
		printHelp(err.Error())
	}
	if !info.IsDir() {
		printHelp("Input -i Not a Directory")
		return
	}

	outfo, err := os.Stat(*out)
	if err != nil {
		printHelp(err.Error())
	}
	if !outfo.IsDir() {
		printHelp("Output -o Not a Directory")
		return
	}

	fl, err := os.ReadDir(*in)
	if err != nil {
		panic(err)
	}

	for _, qtif := range fl {
		inf, err := os.Open(path.Join(*in, qtif.Name()))
		if err != nil {
			panic(err)
		}

		infd, err := io.ReadAll(inf)
		if err != nil {
			panic(err)
		}

		it, err := assessmentitem.NewAssessmentItem(infd)
		if err != nil {
			panic(err)
		}

		b, err := it.ItemBody.MathML()
		if err != nil {
			panic(err)
		}

		for _, v := range b {
			writeFile(v)
		}
		inf.Close()
	}
}

func writeFile(ml string) {
	cr := sha1.Sum([]byte(ml))

	fname := fmt.Sprintf("%x-ml.xml", cr)
	_, err := os.Stat(path.Join(*out, fname))
	if err != nil {
		if os.IsNotExist(err) {
			outF, err := os.OpenFile(path.Join(*out, fname), os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(outF, strings.NewReader(ml))
			if err != nil {
				panic(err)
			}
			outF.Close()
		}
	} else {
		fmt.Printf("Skipping: %s\n", fname)
	}
}

func printHelp(msg string) {
	if msg != "" {
		fmt.Printf("Error: %s\n", msg)
	}

	fmt.Println("-i\t\tInput Directory for QTI Item XML Files")
	fmt.Println("-o\t\tOutput Directory for MathML XML Files")
	fmt.Println("-h\t\tPrint Help")
}
