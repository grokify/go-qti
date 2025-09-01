package assessmentresult

import (
	"encoding/xml"
	"github.com/stmath/go-qti"
)

// ItemVariable. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#Abstract_ItemVariable
type ItemVariable interface{}

type AssessmentResult struct {
	XMLName    xml.Name      `xml:"assessmentResult"`
	Context    Context       `xml:"context"`
	TestResult *TestResult   `xml:"testResult,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#RootAttribute_AssessmentResult_testResult
	ItemResult []*ItemResult `xml:"itemResult,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#RootAttribute_AssessmentResult_itemResult
}

func (ar *AssessmentResult) AddItemResult(ir *ItemResult) error {
	if ar.ItemResult == nil {
		ar.ItemResult = make([]*ItemResult, 0)
	}

	ar.ItemResult = append(ar.ItemResult, ir)
	return nil
}

type Context struct {
	XMLName           xml.Name            `xml:"context"`
	SourcedId         string              `xml:"sourcedId,attr,omitempty"`
	SessionIdentifier []SessionIdentifier `xml:"sessionIdentifier,omitempty"`
}

type SessionIdentifier struct {
	XMLName    xml.Name `xml:"sessionIdentifier"`
	SourceID   string   `xml:"sourceID,attr"`
	Identifier string   `xml:"sessionIdentifier,attr"`
}

// OutcomeVariable. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#AbstractAttribute_ItemVariable_outcomeVariable
type OutcomeVariable struct {
	XMLName            xml.Name    `xml:"outcomeVariable"`
	Identifier         string      `xml:"identifier,attr"`                   // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_identifier
	Cardinality        string      `xml:"cardinality,attr"`                  // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_cardinality
	BaseType           string      `xml:"baseType,attr,omitempty"`           // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_baseType
	View               string      `xml:"view,attr,omitempty"`               // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_view
	Interpretation     string      `xml:"interpretation,attr,omitempty"`     // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_interpretation
	LongInterpretation string      `xml:"longInterpretation,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_longInterpretation
	NormalMaximum      float64     `xml:"normalMaximum,attr,omitempty"`      // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_normalMaximum
	NormalMinimum      float64     `xml:"normalMinimum,attr,omitempty"`      // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_normalMinimum
	MasteryValue       float64     `xml:"masteryValue,attr,omitempty"`       // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_OutcomeVariable.Attr_masteryValue
	Value              []qti.Value `xml:"value,omitempty""`                  // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataAttribute_OutcomeVariable_value
}

// ResponseVariable. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#Data_ResponseVariable
type ResponseVariable struct {
	XMLName        xml.Name `xml:"responseVariable"`
	Identifier     string   `xml:"identifier,attr"`               // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ResponseVariable.Attr_identifier
	Cardinality    string   `xml:"cardinality,attr"`              // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ResponseVariable.Attr_cardinality
	BaseType       string   `xml:"baseType,attr,omitempty"`       // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ResponseVariable.Attr_baseType
	ChoiceSequence string   `xml:"choiceSequence,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#List_IdentifierList

	CorrectResponse   *qti.CorrectResponse `xml:"correctResponse"`   // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataAttribute_ResponseVariable_correctResponse
	CandidateResponse *CandidateResponse   `xml:"candidateResponse"` // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataAttribute_ResponseVariable_candidateResponse
}

// CandidateResponse. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#Data_CandidateResponse
type CandidateResponse struct {
	XMLName xml.Name    `xml:"candidateResponse"`
	Value   []qti.Value `xml:"value,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataAttribute_CandidateResponse_value
}
