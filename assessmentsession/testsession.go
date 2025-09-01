package assessmentsession

import (
	"fmt"
	"github.com/stmath/go-qti"
	asi "github.com/stmath/go-qti/assessmentitem"
	asr "github.com/stmath/go-qti/assessmentresult"
	ast "github.com/stmath/go-qti/assessmenttest"
	"github.com/stmath/go-xmldom"
	"golang.org/x/xerrors"
	"reflect"
	"strconv"
	"time"
)

type TestSession struct {
	Test        *ast.AssessmentTest
	variables   map[string]Variable
	outcomeVars map[string]*OutcomeSessionVaraible
	itemResults map[string]*asr.ItemResult
	Items       []*asi.AssessmentItem
	isProcessed bool
}

func (ts *TestSession) TestResult() *asr.TestResult {
	tr := &asr.TestResult{
		Identifier:    ts.Test.Identifier,
		Datestamp:     time.Now().String(),
		ItemVariables: nil,
	}

	if len(ts.outcomeVars) > 0 {
		tr.ItemVariables = make([]interface{}, 0)

		for _, v := range ts.outcomeVars {
			iv := &asr.OutcomeVariable{
				Identifier:         v.Identifier,
				Cardinality:        v.Cardinality,
				BaseType:           v.BaseType,
				View:               v.View,
				Interpretation:     v.Interpretation,
				LongInterpretation: v.LongInterpretation,
				NormalMaximum:      v.NormalMaximum,
				NormalMinimum:      v.NormalMinium,
				MasteryValue:       v.MasteryValue,
				Value:              v.Value,
			}

			tr.ItemVariables = append(tr.ItemVariables, iv)
		}
	}

	return tr
}

func NewTestSession(test *ast.AssessmentTest) (*TestSession, error) {
	ts := TestSession{
		Test:        test,
		variables:   nil,
		outcomeVars: nil,
		isProcessed: false,
	}

	if test.OutcomeDeclaration != nil {
		ts.variables = make(map[string]Variable, 0)
		ts.outcomeVars = make(map[string]*OutcomeSessionVaraible)
	}

	for _, v := range test.OutcomeDeclaration {
		ocv := &OutcomeSessionVaraible{
			Identifier:         v.Identifier,
			BaseType:           v.BaseType,
			Cardinality:        v.Cardinality,
			Interpretation:     v.Interpretation,
			LongInterpretation: v.LongInterpretation,
			View:               v.View,
			NormalMaximum:      v.NormalMaximum,
			NormalMinium:       v.NormalMinimum,
			MasteryValue:       v.MasteryValue,
			DefaultValue:       v.DefaultValue,
			Value:              nil,
		}
		if ocv.DefaultValue != nil {
			ocv.Value = ocv.DefaultValue.Value
		}

		ts.variables[ocv.Identifier] = ocv
		ts.outcomeVars[ocv.Identifier] = ocv
	}
	return &ts, nil
}

func (is *TestSession) ProcessOutcome(itemresults map[string]*asr.ItemResult) error {
	is.itemResults = itemresults
	if is.Test.OutcomeProcessing == nil {
		return xerrors.Errorf("Test %s: %w", is.Test.Identifier, ErrorOutcomeProcessingMissing)
	}
	if is.Test.OutcomeProcessing.OutcomeRuleGroup == nil {
		return xerrors.Errorf("Test %s: %w", is.Test.Identifier, ErrorOutcomeProcessingMissing)
	}

	for _, v := range is.Test.OutcomeProcessing.OutcomeRuleGroup {
		switch v.Name {
		case "setOutcomeValue":
			err := is.setOutcomeValue(v.GetAttributeValue("identifier"), v.Children)
			if err != nil {
				return err
			}

		default:
			return xerrors.Errorf("Outcome Processing doesn't support %s, Test: %s", v.Name, is.Test.Identifier)
		}
	}

	return nil
}

func (is *TestSession) setOutcomeValue(identifier string, expressionGroup []*xmldom.Node) error {
	var outcomeVar *OutcomeSessionVaraible
	var ok bool
	if outcomeVar, ok = is.outcomeVars[identifier]; !ok {
		return xerrors.Errorf("OutcomeVarialbe %s: %w", ErrorUnknownVariable)
	}

	for _, eg := range expressionGroup {
		switch eg.Name {
		case "sum":
			val, err := is.sum(eg.Children)
			if err != nil {
				return err
			}
			vals := make([]qti.Value, 0)
			vals = append(vals, val)
			outcomeVar.Value = vals
		default:
			return xerrors.Errorf("Outcome Processing doesn't support %s, Test: %s", eg.Name, is.Test.Identifier)
		}

	}

	return nil
}

// FIXME: Unknown how to sum multiple cardinality. Used OpenTAO to run some examples and they have chosen only to use
// the first value
func (is *TestSession) sum(egs []*xmldom.Node) (qti.Value, error) {
	for _, eg := range egs {
		switch eg.Name {
		case "testVariables":
			ocvs := is.testVariables(eg.GetAttributeValue("variableIdentifier"))
			final := qti.Value{
				BaseType: "integer",
				Data:     "0",
			}
			for _, ocv := range ocvs {
				if ocv.Value == nil || len(ocv.Value) < 1 || ocv.Value[0].Data == "" {
					// if any argument is NULL, spec says to return 0 if float or integer
					return qti.Value{
						BaseType: "",
						Data:     "0",
					}, nil
				}
				tval := ocv.Value[0]
				if tval.BaseType == "" {
					tval.BaseType = ocv.BaseType
				}

				ocvT, err := addValue(final, tval)
				if err != nil {
					return qti.Value{}, xerrors.Errorf("OutcomeProcessing::setOutcomeValue::sum: %w", err)
				}
				final = ocvT
			}

			return final, nil
		default:
			return qti.Value{}, xerrors.Errorf("Outcome Processing sum doesn't support %s, Test: %s", eg.Name, is.Test.Identifier)
		}
	}

	return qti.Value{}, xerrors.Errorf("SUM encountered no expressionGroups to sum")
}

func addValue(v1 qti.Value, v2 qti.Value) (qti.Value, error) {
	if v1.BaseType != "integer" && v1.BaseType != "float" {
		return qti.Value{}, xerrors.Errorf("cannot addValue for %s: %w", v1.BaseType, ErrorIncompatibleBaseType)
	}
	if v2.BaseType != "integer" && v2.BaseType != "float" {
		return qti.Value{}, xerrors.Errorf("cannot addValue for %s: %w", v1.BaseType, ErrorIncompatibleBaseType)
	}

	if v1.BaseType == "float" || v2.BaseType == "float" {
		v1f, err := strconv.ParseFloat(v1.Data, 64)
		if err != nil {
			return qti.Value{}, err
		}
		v2f, err := strconv.ParseFloat(v2.Data, 64)
		if err != nil {
			return qti.Value{}, err
		}
		v3f := v1f + v2f
		v3 := qti.Value{
			BaseType: "float",
			Data:     fmt.Sprintf("%f", v3f),
		}
		return v3, nil
	} else {
		v1i, err := strconv.ParseInt(v1.Data, 10, 32)
		if err != nil {
			return qti.Value{}, err
		}
		v2i, err := strconv.ParseInt(v2.Data, 10, 32)
		if err != nil {
			return qti.Value{}, err
		}
		v3i := v1i + v2i
		v3 := qti.Value{
			BaseType: "integer",
			Data:     fmt.Sprintf("%d", v3i),
		}
		return v3, nil
	}

}

func (is *TestSession) testVariables(identifier string) []*asr.OutcomeVariable {
	ocvs := make([]*asr.OutcomeVariable, 0)
	for _, ir := range is.itemResults {
		for _, iv := range ir.ItemVariables {
			if reflect.TypeOf(iv) == reflect.TypeOf(&asr.OutcomeVariable{}) {
				ocv := iv.(*asr.OutcomeVariable)
				if ocv.Identifier == identifier {
					ocvs = append(ocvs, ocv)
				}
			}
		}
	}

	return ocvs
}
