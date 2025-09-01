package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/stmath/go-qti"
	"github.com/stmath/go-qti/assessmentitem"
	asis "github.com/stmath/go-qti/assessmentsession"
	"io/ioutil"
	"os"
)

var (
	i = flag.String("i", "", "Path to AssessmentItem")
)

func main() {
	flag.Parse()
	doItem()
	//doItemDom()
	return
}

func doItem() {
	var inf *os.File
	var err error

	if *i != "" {
		inf, err = os.Open(*i)
		if err != nil {
			panic(err)
		}
	}
	defer inf.Close()

	infd, err := ioutil.ReadAll(inf)
	if err != nil {
		panic(err)
	}

	itm, err := assessmentitem.NewAssessmentItem(infd)
	if err != nil {
		panic(err)
	}

	item, err := asis.NewItemSessionFromItem(itm)
	if err != nil {
		panic(err)
	}

	vals := make([]qti.Value, 0)
	val := qti.Value{BaseType: "identifier", Data: "a4"}
	vals = append(vals, val)
	respVar := asis.ResponseSessionVariable{
		Identifier:   "RESPONSE",
		BaseType:     "identifier",
		Cardinality:  "single",
		DefaultValue: nil,
		Value:        vals,
	}
	rArr := make([]*asis.ResponseSessionVariable, 0)
	rArr = append(rArr, &respVar)

	err = item.ProcessResponse(rArr)
	if err != nil {
		panic(err)
	}

	ir, err := item.ItemResult()
	if err != nil {
		panic(err)
	}

	b, err := xml.Marshal(ir)
	fmt.Printf("%s\n", b)

}

func doItemDom() {
	item, err := asis.NewItemSessionParseFile(*i)
	if err != nil {
		panic(err)
	}
	vals := make([]qti.Value, 0)
	val := qti.Value{BaseType: "identifier", Data: "a4"}
	vals = append(vals, val)

	respVar := asis.ResponseSessionVariable{
		Identifier:   "RESPONSE",
		BaseType:     "identifier",
		Cardinality:  "single",
		DefaultValue: nil,
		Value:        vals,
	}
	rArr := make([]*asis.ResponseSessionVariable, 0)
	rArr = append(rArr, &respVar)

	err = item.ProcessResponse(rArr)
	if err != nil {
		panic(err)
	}

	for _, v := range item.OutcomeVars() {
		fmt.Printf("Outcome: %s %s\n", v.Identifier, v.Value[0].Data)
	}
}
