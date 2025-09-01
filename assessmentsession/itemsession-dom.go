package assessmentsession

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/stmath/go-xmldom"
	"golang.org/x/xerrors"

	"github.com/grokify/go-qti"
	"github.com/grokify/go-qti/assessmentitem"
	asr "github.com/grokify/go-qti/assessmentresult"
)

var (
	ErrorResponseIdentNotFound       = errors.New("ResponseDeclaration Identifier Not Found")
	ErrorResponseIfNotFound          = errors.New("Required responseIf not found")
	ErrorUnexpectedNode              = errors.New("Unexpected Node")
	ErrorUnknownResponseVariable     = errors.New("Unknown Response Variable")
	ErrorIncompatibleBaseType        = errors.New("Incompatible BaseType")
	ErrorBaseTypeMisMatch            = errors.New("BaseType Mismatch")
	ErrorUnsupportedFeature          = errors.New("Unsupported Feature")
	ErrorUnknownVariable             = errors.New("Not Found")
	ErrorMissingValue                = errors.New("Missing Value")
	ErrorDuplicateVariableIdentifier = errors.New("Duplicate variable identifiers")
	ErrorMissingTemplate             = errors.New("QTI Template Not Defined")
	ErrorResponseProcessingMissing   = errors.New("Missing ResponseProcessing")
	ErrorOutcomeProcessingMissing    = errors.New("Missing OutcomeProcessing")
)

type ItemSession struct {
	dom         *xmldom.Document
	itm         *assessmentitem.AssessmentItem
	identifier  string
	variables   map[string]Variable
	responseVar map[string]*ResponseSessionVariable
	outcomeVar  map[string]*OutcomeSessionVaraible
	processed   bool
}

func NewItemSessionParseFile(filename string) (*ItemSession, error) {
	dom, err := xmldom.ParseFile(filename)
	if err != nil {
		return nil, err
	}

	is := &ItemSession{dom: dom}
	err = is.initResponseVars()
	if err != nil {
		return nil, err
	}

	err = is.initOutcomeVars()
	if err != nil {
		return nil, err
	}
	is.processed = false

	is.identifier = is.dom.Root.GetAttributeValue("identifier")

	return is, nil
}

func NewItemSession(itemxml string) (*ItemSession, error) {
	dom, err := xmldom.ParseXML(itemxml)
	if err != nil {
		return nil, err
	}

	is := &ItemSession{dom: dom, variables: make(map[string]Variable), responseVar: make(map[string]*ResponseSessionVariable), outcomeVar: make(map[string]*OutcomeSessionVaraible)}
	err = is.initResponseVars()
	if err != nil {
		return nil, err
	}

	err = is.initOutcomeVars()
	if err != nil {
		return nil, err
	}
	is.processed = false

	is.identifier = is.identifier

	return is, nil
}

func (is *ItemSession) initResponseVars() error {
	respNodes := is.dom.Root.GetChildren("responseDeclaration")
	if is.responseVar == nil {
		is.responseVar = make(map[string]*ResponseSessionVariable, len(respNodes))
	}
	if is.variables == nil {
		is.variables = make(map[string]Variable, len(respNodes))
	}

	for _, decnode := range respNodes {
		var respDec assessmentitem.ResponseDeclaration
		err := xml.Unmarshal([]byte(decnode.XML()), &respDec)
		if err != nil {
			panic(err)
		}

		v := &ResponseSessionVariable{
			Identifier:   respDec.Identifier,
			BaseType:     respDec.BaseType,
			Cardinality:  respDec.Cardinality,
			DefaultValue: respDec.DefaultValue,
			Value:        nil,
		}

		is.responseVar[respDec.Identifier] = v
		is.variables[respDec.Identifier] = v
	}

	return nil
}

func (is *ItemSession) initOutcomeVars() error {
	outcomeNodes := is.dom.Root.GetChildren("outcomeDeclaration")
	if is.variables == nil {
		is.variables = make(map[string]Variable, len(outcomeNodes))
	}
	if is.outcomeVar == nil {
		is.outcomeVar = make(map[string]*OutcomeSessionVaraible, len(outcomeNodes))
	}

	for _, decnode := range outcomeNodes {
		var outDec qti.OutcomeDeclaration
		err := xml.Unmarshal([]byte(decnode.XML()), &outDec)
		if err != nil {
			panic(err)
		}

		v := &OutcomeSessionVaraible{
			Identifier:    outDec.Identifier,
			BaseType:      outDec.BaseType,
			Cardinality:   outDec.Cardinality,
			NormalMaximum: outDec.NormalMaximum,
			NormalMinium:  outDec.NormalMinimum,
			DefaultValue:  outDec.DefaultValue,
			View:          outDec.View,
			Value:         nil,
		}

		if outDec.DefaultValue != nil {
			v.Value = outDec.DefaultValue.Value
		}

		is.outcomeVar[outDec.Identifier] = v
		is.variables[outDec.Identifier] = v
	}

	return nil
}

func (is *ItemSession) ProcessResponse(vars []*ResponseSessionVariable) error {
	// Populate declared response variables with passed in values
	for _, v := range vars {
		passed_ident := v.Identifier
		var rvar Variable
		var ok bool
		if rvar, ok = is.variables[passed_ident]; !ok {
			return xerrors.Errorf("Item: %s ResponseDeclaration Identifier %s: %w", is.identifier, passed_ident, ErrorUnknownResponseVariable)
		}
		// Check that types match
		if rvar.GetBaseType() != v.BaseType {
			return xerrors.Errorf("Item %s, Response Identifier %s BaseType: %s, ResponseSessonVariable %s BaseType: %s: %w", is.identifier, v.Identifier, v.BaseType, rvar.GetIdentifier(), rvar.GetBaseType(), ErrorBaseTypeMisMatch)
		}
		rvar.SetValue(v.Value)
	}

	if is.itm != nil {
		return is.processResponseITEM(vars)

	} else if is.dom != nil {
		return is.processResponseDOM(vars)
	} else {
		return errors.New("No Processing Defined")
	}

	is.processed = true
	return nil
}

func (is *ItemSession) processResponseITEM(vars []*ResponseSessionVariable) error {
	var procNodes []*xmldom.Node
	if procNodes == nil {
		procNodes = make([]*xmldom.Node, 0)
	}

	if is.itm.ResponseProcessing.Template == "" {
		procNodes = append(procNodes, is.itm.ResponseProcessing.ResponseRuleGroup...)
	} else {
		doc, err := loadRPTemplate(is.itm.ResponseProcessing.Template)
		if err != nil {
			return err
		}

		procNodes = append(procNodes, doc.Root.Children...)
	}

	for _, pn := range procNodes {
		err := responseCondition(is, pn)
		if err != nil {
			return err
		}
	}

	is.processed = true
	return nil
}

func (is *ItemSession) processResponseDOM(vars []*ResponseSessionVariable) error {
	var procNodes []*xmldom.Node
	if procNodes == nil {
		procNodes = make([]*xmldom.Node, 0)
	}

	rpNode := is.dom.Root.GetChild("responseProcessing")
	if rpNode == nil {
		return xerrors.Errorf("No responseProcessing for %s: %w", is.identifier, ErrorResponseProcessingMissing)
	}

	if rpNode.GetAttributeValue("template") == "" {
		rpNodes := is.dom.Root.GetChildren("responseProcessing")
		for _, v := range rpNodes {
			procNodes = append(procNodes, v.GetChild("responseCondition"))
		}
	} else {
		doc, err := loadRPTemplate(rpNode.GetAttributeValue("template"))
		if err != nil {
			return err
		}

		procNodes = append(procNodes, doc.Root.Children...)
	}

	for _, pn := range procNodes {
		err := responseCondition(is, pn)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadRPTemplate(id string) (*xmldom.Document, error) {
	if xml, ok := qti.ResponseTemplates[id]; !ok {
		return nil, xerrors.Errorf("ResponseProcessing Template %s not defined: %w", id, ErrorMissingTemplate)
	} else {
		doc, err := xmldom.ParseXML(xml)
		if err != nil {
			return nil, xerrors.Errorf("ResponseProcessing Template %s parse error: %w", id, err)
		}
		return doc, nil
	}
}

func responseCondition(is *ItemSession, condNode *xmldom.Node) error {
	fmt.Println("Running ResponseCondition")
	ifNode := condNode.GetChild("responseIf")
	if ifNode == nil {
		return xerrors.Errorf("XML: %s: %w", condNode.XML(), ErrorResponseIfNotFound)
	}

	ifResult, err := responseIf(is, ifNode)
	if err != nil {
		return xerrors.Errorf("XML %s: %w", condNode.XML(), err)
	}
	if ifResult {
		return nil
	}

	elseIfNodes := condNode.GetChildren("responseElseIf")
	for _, elseIfNode := range elseIfNodes {
		ifResult, err = responseIf(is, elseIfNode)
		if err != nil {
			return xerrors.Errorf("XML %s: %w", condNode.XML(), err)
		}
		if ifResult {
			return nil
		}
	}

	elseNode := condNode.GetChild("responseElse")
	if elseNode != nil {
		fmt.Println("Running responseElse")
		err := responseRuleGroup(is, elseNode)
		if err != nil {
			return xerrors.Errorf("XML %s: %w", condNode.XML(), err)
		}
	}

	return nil
}

func responseRuleGroup(is *ItemSession, node *xmldom.Node) error {

	return nil
}

func responseIf(is *ItemSession, node *xmldom.Node) (bool, error) {
	fmt.Printf("Running %s\n", node.Name)
	if node.Name != "responseIf" && node.Name != "responseElseIf" {
		return false, xerrors.Errorf("ExpectedNode: responseIf, ActualNode: %s: %w", node.Name, ErrorUnexpectedNode)
	}
	children := node.Children
	for i, child := range children {
		if i == 0 {
			result, err := expressionGroup(is, child)
			if err != nil {
				return result, xerrors.Errorf("XML %s: %w", child.XML(), err)
			}

			if !result {
				return false, nil
			}
		} else {
			switch child.Name {
			case "responseCondition":
				responseCondition(is, child)
			case "setOutcomeValue":
				err := setOutcomeValue(is, child)
				if err != nil {
					return false, err
				}
			case "exitResponse":
				return true, nil
			default:
				// TODO: Report correct error
				return true, xerrors.Errorf("XML %s: %w", node.XML(), errors.New("shouldn't have reached default switch"))
			}
		}

	}

	return true, nil
}

func expressionGroup(is *ItemSession, node *xmldom.Node) (bool, error) {
	fmt.Println("Running expressionGroup")

	switch node.Name {
	case "match":
		matchChildren := node.Children
		if len(matchChildren) == 2 {
			var args [2]*xmldom.Node
			args[0] = matchChildren[0]
			args[1] = matchChildren[1]
			return match(is, args)
		}

	}
	return false, nil
}

func match(is *ItemSession, nodes [2]*xmldom.Node) (bool, error) {
	var a Variable
	var b Variable
	var err error
	var ok bool
	if nodes[0].Name != "variable" && nodes[0].Name != "correct" {
		return false, xerrors.Errorf("%s in expressionGroup.match: %w", ErrorUnsupportedFeature)
	}
	if nodes[1].Name != "variable" && nodes[1].Name != "correct" {
		return false, xerrors.Errorf("%s in expressionGroup.match: %w", ErrorUnsupportedFeature)
	}

	idA := nodes[0].GetAttribute("identifier").Value
	idB := nodes[0].GetAttribute("identifier").Value

	if nodes[0].Name == "variable" {
		if a, ok = is.variables[idA].(Variable); !ok {
			return false, xerrors.Errorf("%s: %w", idA, ErrorUnknownVariable)
		}
	} else if nodes[0].Name == "correct" {
		a, err = is.correct(idA)
		if err != nil {
			return false, err
		}
	}

	if nodes[1].Name == "variable" {
		b = is.variables[idB].(Variable)
	} else if nodes[1].Name == "correct" {
		b, err = is.correct(idB)
		if err != nil {
			return false, err
		}
	}

	return a.Equal(b), nil
}

func (is *ItemSession) correct(identifier string) (Variable, error) {
	fmt.Println("Running Get Correct")
	if is.dom != nil {

		respNodes := is.dom.Root.QueryOne(fmt.Sprintf(`//responseDeclaration[@identifier='%s']`, identifier))
		if respNodes == nil {
			return nil, xerrors.Errorf("GetCorrect %s: %w", identifier, ErrorUnknownResponseVariable)
		}

		node := respNodes
		r := &ResponseSessionVariable{
			Identifier:  node.GetAttributeValue("identifier"),
			BaseType:    node.GetAttributeValue("baseType"),
			Cardinality: node.GetAttributeValue("cardinality"),
		}

		if node.GetChild("correctResponse") != nil {
			vs := node.GetChild("correctResponse").GetChildren("value")
			if vs != nil {
				vsar := make([]qti.Value, 0)
				for _, v := range vs {
					val := qti.Value{
						FieldIdentifier: v.GetAttributeValue("fieldIdentifier"),
						BaseType:        v.GetAttributeValue("baseType"),
						Data:            v.Text,
					}
					vsar = append(vsar, val)
				}
				if len(vsar) < 1 {
					return nil, xerrors.Errorf("%s: %w", node.GetChild("correctResponse").XML(), ErrorMissingValue)
				}

				r.Value = vsar
				return r, nil
			}
		} else {
			return nil, xerrors.Errorf("GetCorrect Identifier %s, %s: %w", identifier, node.XML(), ErrorMissingValue)
		}
	} else if is.itm != nil {
		var rspS *assessmentitem.ResponseDeclaration
		var cr *qti.CorrectResponse

		for _, rs := range is.itm.ResponseDeclaration {
			if rs.Identifier == identifier {
				rspS = rs
				if rspS.CorrectResponse != nil {
					cr = rspS.CorrectResponse
				}
			}
		}

		if cr == nil {
			return nil, xerrors.Errorf("GetCorrect Identifier %s: %w", identifier, ErrorMissingValue)
		}

		r := &ResponseSessionVariable{
			Identifier:  identifier,
			BaseType:    rspS.BaseType,
			Cardinality: rspS.Cardinality,
			Value:       cr.Value,
		}
		return r, nil
	}

	return nil, xerrors.Errorf("GetCorrect Identifier %s: %w", identifier, ErrorMissingValue)
}

func setOutcomeValue(is *ItemSession, node *xmldom.Node) error {
	// TODO: Only supports Single Cardinality
	var ok bool
	var outComeVar *OutcomeSessionVaraible
	if outComeVar, ok = is.outcomeVar[node.GetAttribute("identifier").Value]; !ok {
		return xerrors.Errorf("setOutcome Identifier [\"%s\"]: %w", node.GetAttribute("identifier").Value, ErrorUnknownVariable)
	}

	switch node.FirstChild().Name {
	case "sum":
		val, baseType, err := sum(is, node.FirstChild().Children)
		if err != nil {
			return err
		}
		if baseType != outComeVar.BaseType {
			return xerrors.Errorf("OutcomeSessionVaraible baseType %s != sum() baseType. %s, %s", outComeVar.BaseType, baseType, node.XML())
		}
		arrVal := make([]qti.Value, 0)
		arrVal = append(arrVal, qti.Value{
			BaseType: baseType,
			Data:     val.(string),
		})
		outComeVar.Value = arrVal
	case "baseValue":
		val := baseValue(is, node.FirstChild())
		if outComeVar.BaseType != val.BaseType {
			return xerrors.Errorf("Item %s, OutcomeVar Identifier %s BaseType: %s, baseValue BaseType: %s: %w", is.identifier, outComeVar.Identifier, outComeVar.BaseType, val.BaseType, ErrorBaseTypeMisMatch)
		}
		valAr := make([]qti.Value, 0)
		valAr = append(valAr, *val)
		outComeVar.Value = valAr

	default:
		return xerrors.Errorf("Unknown setOutcomeValue operation %s: %w", node.FirstChild().Name, ErrorUnexpectedNode)
	}

	return nil
}

func baseValue(is *ItemSession, node *xmldom.Node) *qti.Value {
	return &qti.Value{BaseType: node.GetAttributeValue("baseType"), Data: node.Text}
}

func sum(is *ItemSession, nodes []*xmldom.Node) (interface{}, string, error) {
	vars := make([]Variable, 0)
	for _, v := range nodes {
		var ok bool
		ident := v.GetAttribute("identifier").Value
		switch v.Name {
		case "variable":
			var vvar Variable
			if vvar, ok = is.variables[ident]; ok {
				vars = append(vars, vvar)
			}
		case "baseValue":
			strVal := v.Text
			baseType := v.GetAttribute("baseType").Value
			arVar := make([]qti.Value, 0)
			arVar = append(arVar, qti.Value{Data: strVal, BaseType: baseType})
			rsvp := &ResponseSessionVariable{
				Identifier:  "-",
				BaseType:    baseType,
				Cardinality: "single",
				Value:       arVar,
			}
			vars = append(vars, rsvp)
		case "default":
			return "", "", xerrors.Errorf("sum type not supported [%s]: %w", v.Name, ErrorUnexpectedNode)
		}
	}

	intMode := true
	for _, v := range vars {
		if v.GetBaseType() != "integer" {
			intMode = false
			break
		}
	}
	if intMode {
		for _, v := range vars {
			if v.GetBaseType() != "integer" {
				intMode = false
				break
			}
		}
	}

	if intMode {
		var sum int
		for _, v := range vars {
			i, err := strconv.ParseInt(v.GetValue()[0].Data, 10, 32)
			if err != nil {
				return "", "", xerrors.Errorf("Invalid Integer String %s: %w", v.GetValue()[0].Data, errors.New("Conversion Error"))
			}
			sum += int(i)
		}
		return fmt.Sprintf("%d", sum), "integer", nil

	} else {
		var sum float64
		for _, v := range vars {
			i, err := strconv.ParseFloat(v.GetValue()[0].Data, 64)
			if err != nil {
				return "", "", xerrors.Errorf("Invalid Integer String %s: %w", v.GetValue()[0].Data, errors.New("Conversion Error"))
			}
			sum += i
		}
		return fmt.Sprintf("%f", sum), "float", nil
	}

	return "", "", errors.New("Sum failed to return correctly")

}

func (is *ItemSession) ItemResult() (*asr.ItemResult, error) {
	if !is.processed {
		return nil, xerrors.Errorf("ItemSession has not run processing")
	}

	ir := &asr.ItemResult{
		XMLName: xml.Name{
			Space: "",
			Local: "itemResult",
		},
		Identifier:       is.itm.Identifier,
		Datestamp:        time.Now().String(),
		SessionStatus:    "final",
		CandidateComment: "",
	}

	for _, oc := range is.outcomeVar {
		if ir.ItemVariables == nil {
			ir.ItemVariables = make([]interface{}, 0)
		}

		roc := &asr.OutcomeVariable{
			XMLName:            xml.Name{Local: "outcomeVariable"},
			Identifier:         oc.Identifier,
			Cardinality:        oc.Cardinality,
			BaseType:           oc.BaseType,
			View:               oc.View,
			Interpretation:     oc.Interpretation,
			LongInterpretation: oc.LongInterpretation,
			NormalMaximum:      oc.NormalMaximum,
			NormalMinimum:      oc.NormalMinium,
			MasteryValue:       oc.MasteryValue,
			Value:              oc.Value,
		}

		ir.ItemVariables = append(ir.ItemVariables, roc)
	}

	for _, rs := range is.responseVar {
		if ir.ItemVariables == nil {
			ir.ItemVariables = make([]interface{}, 0)
		}

		rspv := &asr.ResponseVariable{
			XMLName:     xml.Name{Local: "responseVariable"},
			Identifier:  rs.Identifier,
			Cardinality: rs.Cardinality,
			BaseType:    rs.BaseType,
		}
		cr, err := is.correct(rs.Identifier)
		if err == nil {
			rspv.CorrectResponse = &qti.CorrectResponse{
				Value: cr.GetValue(),
			}
		}
		rspv.CandidateResponse = &asr.CandidateResponse{
			Value: rs.Value,
		}

		ir.ItemVariables = append(ir.ItemVariables, rspv)
	}

	return ir, nil
}

func (is *ItemSession) OutcomeVars() []*OutcomeSessionVaraible {
	if len(is.outcomeVar) < 1 {
		return nil
	}

	ovs := make([]*OutcomeSessionVaraible, 0)
	for _, v := range is.outcomeVar {
		ovs = append(ovs, v)
	}

	return ovs
}
