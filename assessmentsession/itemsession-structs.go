package assessmentsession

import "github.com/stmath/go-qti/assessmentitem"

func NewItemSessionFromItem(itm *assessmentitem.AssessmentItem) (*ItemSession, error) {
	is := &ItemSession{
		dom:       nil,
		itm:       itm,
		processed: false,
	}

	err := is.initResponseVarsFromItem()
	if err != nil {
		return nil, err
	}

	err = is.initOutcomeVarsFromItem()
	if err != nil {
		return nil, err
	}

	return is, nil
}

func (is *ItemSession) initResponseVarsFromItem() error {
	if is.variables == nil {
		is.variables = make(map[string]Variable)
	}
	if is.responseVar == nil {
		is.responseVar = make(map[string]*ResponseSessionVariable, 0)
	}

	for _, declar := range is.itm.ResponseDeclaration {
		r := &ResponseSessionVariable{
			Identifier:   declar.Identifier,
			BaseType:     declar.BaseType,
			Cardinality:  declar.Cardinality,
			DefaultValue: declar.DefaultValue,
			Value:        nil,
		}
		is.responseVar[r.Identifier] = r
		is.variables[r.Identifier] = r
	}

	return nil
}

func (is *ItemSession) initOutcomeVarsFromItem() error {
	if is.variables == nil {
		is.variables = make(map[string]Variable)
	}
	if is.outcomeVar == nil {
		is.outcomeVar = make(map[string]*OutcomeSessionVaraible, 0)
	}

	for _, declar := range is.itm.OutcomeDeclaration {
		r := &OutcomeSessionVaraible{
			Identifier:         declar.Identifier,
			BaseType:           declar.BaseType,
			Cardinality:        declar.Cardinality,
			Interpretation:     declar.Interpretation,
			LongInterpretation: declar.LongInterpretation,
			NormalMaximum:      declar.NormalMaximum,
			NormalMinium:       declar.NormalMinimum,
			DefaultValue:       declar.DefaultValue,
			MasteryValue:       declar.MasteryValue,
			View:               declar.View,
			Value:              nil,
		}

		if r.DefaultValue != nil {
			r.Value = r.DefaultValue.Value
		}

		is.outcomeVar[r.Identifier] = r
		is.variables[r.Identifier] = r
	}

	return nil
}
