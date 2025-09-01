package assessmentsession

import "github.com/stmath/go-qti"

type Variable interface {
	GetIdentifier() string
	SetIdentifier(string)
	GetBaseType() string
	GetCardinality() string
	GetValue() []qti.Value
	SetValue([]qti.Value)
	AddValue(value qti.Value)
	Equal(v Variable) bool
}

type ResponseSessionVariable struct {
	Identifier   string            `xml:"identifier,attr" json:"identifier"`
	BaseType     string            `xml:"baseType,attr" json:"baseType"`
	Cardinality  string            `xml:"cardinality,attr" json:"cardinality"`
	DefaultValue *qti.DefaultValue `xml:"defaultValue,attr" json:"defaultValue"`
	Value        []qti.Value       `xml:"value" json:"values"`
}

func (rv *ResponseSessionVariable) GetIdentifier() string {
	return rv.Identifier
}
func (rv *ResponseSessionVariable) SetIdentifier(id string) {
	rv.Identifier = id
}
func (rv *ResponseSessionVariable) GetBaseType() string {
	return rv.BaseType
}
func (rv *ResponseSessionVariable) GetCardinality() string {
	return rv.Cardinality
}
func (rv *ResponseSessionVariable) GetValue() []qti.Value {
	return rv.Value
}
func (rv *ResponseSessionVariable) SetValue(vals []qti.Value) {
	rv.Value = vals
}
func (rv *ResponseSessionVariable) AddValue(val qti.Value) {
	if rv.Value != nil {
		rv.Value = make([]qti.Value, 0)
	}
	rv.Value = append(rv.Value, val)
}
func (rv *ResponseSessionVariable) Equal(v Variable) bool {
	if rv.Cardinality != v.GetCardinality() {
		return false
	}
	if rv.BaseType != rv.BaseType {
		return false
	}

	found := false
	for _, ov := range rv.Value {
		found = false
		for _, j := range v.GetValue() {
			if j.Data == ov.Data {
				found = true
				break
			}
		}
		if found != true {
			return false
		}
	}

	return found
}

type OutcomeSessionVaraible struct {
	Identifier         string
	BaseType           string
	Cardinality        string
	Interpretation     string
	LongInterpretation string
	View               string
	NormalMaximum      float64
	NormalMinium       float64
	MasteryValue       float64
	DefaultValue       *qti.DefaultValue
	Value              []qti.Value
}

func (ov *OutcomeSessionVaraible) GetIdentifier() string {
	return ov.Identifier
}
func (ov *OutcomeSessionVaraible) SetIdentifier(id string) {
	ov.Identifier = id
}
func (ov *OutcomeSessionVaraible) GetBaseType() string {
	return ov.BaseType
}
func (ov *OutcomeSessionVaraible) GetCardinality() string {
	return ov.Cardinality
}
func (ov *OutcomeSessionVaraible) GetValue() []qti.Value {
	return ov.Value
}
func (ov *OutcomeSessionVaraible) SetValue(vals []qti.Value) {
	ov.Value = vals
}
func (ov *OutcomeSessionVaraible) AddValue(val qti.Value) {
	if ov.Value != nil {
		ov.Value = make([]qti.Value, 0)
	}
	ov.Value = append(ov.Value, val)
}
func (ov *OutcomeSessionVaraible) Equal(v Variable) bool {
	if ov.Cardinality != v.GetCardinality() {
		return false
	}
	if ov.BaseType != ov.BaseType {
		return false
	}

	found := false
	for _, ov := range ov.Value {
		found = false
		for _, j := range v.GetValue() {
			if j.Data == ov.Data {
				found = true
				break
			}
		}
		if found != true {
			return false
		}
	}

	return found
}
