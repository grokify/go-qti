package assessmenttest

import (
	"encoding/xml"

	"github.com/grokify/go-qti"
)

// Test implements the imsglobal AssessmentTest Class. An assessment test is a group of assessmentItems with an associated set of rules that determine which of the items the candidate sees. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Root_AssessmentTest
type AssessmentTest struct {
	XMLName            xml.Name                  `xml:"assessmentTest" json:"-"`
	Identifier         string                    `xml:"identifier,attr"`              // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentTest.Attr_identifier
	Title              string                    `xml:"title,attr"`                   // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentTest.Attr_title
	ToolName           string                    `xml:"toolName,attr,omitempty"`      // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentTest.Attr_toolName
	ToolVersion        string                    `xml:"toolVersion,attr,omitempty"`   // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentTest.Attr_toolVersion
	OutcomeDeclaration []*qti.OutcomeDeclaration `xml:"outcomeDeclaration,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentTest_outcomeDeclaration
	TimeLimits         []*TimeLimits             `xml:"timeLimits,omitemtpy"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentTest_timeLimits
	TestPart           []*TestPart               `xml:"testPart"`                     // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentTest_testPart
	OutcomeProcessing  *OutcomeProcessing        `xml:"outcomeProcessing,omitempty"`  // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentTest_outcomeProcessing
}

func NewAssessmentTest(b []byte) (*AssessmentTest, error) {
	var t AssessmentTest
	e := xml.Unmarshal(b, &t)
	if e != nil {
		return nil, e
	}

	return &t, nil
}

// TimeLimits implements the ismglobal TimeLimits Class https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Derived_TimeLimits
type TimeLimits struct {
	XMLName             xml.Name `xml:"timeLimits" json:"-"`
	MinTime             float64  `xml:"minTime,omitempty"`             // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_TimeLimits.Attr_minTime
	MaxTime             float64  `xml:"maxTime,omitempty"`             // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_TimeLimits.Attr_maxTime
	AllowLateSubmission bool     `xml:"allowLateSubmission,omitempty"` // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_TimeLimits.Attr_allowLateSubmission
}

// SetDefault Sets the default values according to the imsglobal spec
func (tl *TimeLimits) SetDefaults() {
	tl.AllowLateSubmission = false
}

// TestPart. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_TestPart
type TestPart struct {
	XMLName        xml.Name `xml:"testPart" json:"-"`
	Identifier     string   `xml:"identifier,attr"`     // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TestPart.Attr_identifier
	NavigationMode string   `xml:"navigationMode,attr"` // Required. Enumerated value set of: { linear | nonlinear }. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TestPart.Attr_navigationMode
	SubmissionMode string   `xml:"submissionMode,attr"` // Required. Enumerated value set of: { individual | simultaneous }. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TestPart.Attr_submissionMode

	PreCondition       []*qti.PreCondition `xml:"preCondition,omitempty"`       // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TestPart_preCondition
	BranchRule         []*BranchRule       `xml:"branchRule,omitempty"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TestPart_branchRule
	ItemSessionControl *ItemSessionControl `xml:"itemSessionControl,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TestPart_itemSessionControl
	TimeLimits         *TimeLimits         `xml:"timeLimits,omitempty"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TestPart_timeLimits
	// TODO: Make AssessmentSection actually the abstract AssessmentSectionSelection.
	AssessmentSection    []*AssessmentSection  `xml:"assessmentSection,omitempty"`    // Required (Or AssessmentSectionRef). https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TestPart_assessmentSectionSelection
	AssessmentSectionRef *AssessmentSectionRef `xml:"assessmentSectionRef,omitempty"` // Required (Or AssessmentSection). https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#AbstractAttribute_AssessmentSectionSelection_assessmentSectionRef
	// TODO: Implement TestFeedback. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TestPart_testFeedback
}

type AssessmentSectionRef struct {
	XMLname    xml.Name `xml:"assessmentSectionRef" json:"-"`
	Identifier string   `xml:"identifier,attr"`
	Href       string   `xml:"href,attr"`
}

// BranchRule is evaluated after the item, section or part has been presented to the candidate. Implements Type https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_BranchRule
type BranchRule struct {
	XMLName xml.Name             `xml:"branchRule" json:"-"`
	Target  string               `xml:"target,attr"` // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_BranchRule.Attr_target
	Logic   *qti.ExpressionGroup `xml:"logic"`       // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_BranchRule_logic
}

// ItemSessionControl Implements Type https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentSection_itemSessionControl
type ItemSessionControl struct {
	XMLName           xml.Name `xml:"itemSessionControl" json:"-"`      // Name of Element
	MaxAttempts       int      `xml:"maxAttempts,attr,omitempty"`       // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_maxAttempts
	ShowFeedback      bool     `xml:"showFeedback,attr,omitempty"`      // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_showFeedback
	AllowReview       bool     `xml:"allowReview,attr,omitempty"`       // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_allowReview
	ShowSolution      bool     `xml:"showSolution,attr,omitempty"`      // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_showSolution
	AllowComment      bool     `xml:"allowComment,attr,omitempty"`      // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_allowComment
	AllowSkipping     bool     `xml:"allowSkipping,attr,omitempty"`     // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_allowSkipping
	ValidateResponses bool     `xml:"validateResponses,attr,omitempty"` // https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_ItemSessionControl.Attr_validateResponses
}

// SetDefault Sets the default values according to the imsglobal spec
func (isc *ItemSessionControl) SetDefault() {
	isc.ValidateResponses = false
	isc.AllowSkipping = true
	isc.ShowFeedback = false
	isc.AllowReview = true
	isc.ShowSolution = false
	isc.AllowComment = false
}

type AssessmentItemRef struct {
	XMLName    xml.Name `xml:"assessmentItemRef" json:"-"`
	Identifier string   `xml:"identifier,attr"`
	Required   bool     `xml:"required,attr,omitempty"` // Default False
	Fixed      bool     `xml:"fixed,attr,omitempty"`    // Default False
	Href       string   `xml:"href,attr"`
	Category   string   `xml:"category,attr,omitempty"`

	PreCondition       []*qti.PreCondition `xml:"preCondition,omitempty"`
	BranchRule         []*BranchRule       `xml:"branchRule,omitempty"`
	ItemSessionControl *ItemSessionControl `xml:"itemSessionControl,omitempty"`
	TimeLimits         *TimeLimits         `xml:"timeLimits,omitempty"`
	VariableMapping    []*VariableMapping  `xml:"variableMapping,omitempty"`
	TemplateDefault    []*TemplateDefault  `xml:"templateDefault,omitempty"`
}

type VariableMapping struct {
	XMLName          xml.Name `xml:"variableMapping" json:"-"`
	SourceIdentifier string   `xml:"sourceIdentifier,attr"`
	TargetIdentifier string   `xml:"targetIdentifier,attr"`
}

type Weight struct {
	XMLName    xml.Name `xml:"weight" json:"-"`
	Identifier string   `xml:"identifier,attr"`
	Value      float64  `xml:"value,attr"`
}

type TemplateDefault struct {
	XMLName            xml.Name             `xml:"templateDefault" json:"-"`
	TemplateIdentifier string               `xml:"templateIdentifier,attr"`
	ExpressionGroup    *qti.ExpressionGroup `xml:"expressionGroup"`
}
