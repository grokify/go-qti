package assessmentitem

import (
	"encoding/xml"

	"github.com/grokify/go-qti"
)

type AssessmentItem struct {
	XMLName               xml.Name                 `xml:"assessmentItem" json:"-"`
	Identifier            string                   `xml:"identifier,attr"`
	Title                 string                   `xml:"title,attr"`
	Label                 string                   `xml:"label,attr"`
	Language              string                   `xml:"language,attr"`
	ToolName              string                   `xml:"toolName,attr,omitempty"`
	ToolVersion           string                   `xml:"toolVersion,attr,omitempty"`
	Adaptive              bool                     `xml:"adaptive,attr"`
	TimeDependent         bool                     `xml:"timeDependent,attr"`
	ResponseDeclaration   []*ResponseDeclaration   `xml:"responseDeclaration,omitempty"`   // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_responseDeclaration
	OutcomeDeclaration    []qti.OutcomeDeclaration `xml:"outcomeDeclaration,omitempty"`    // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_outcomeDeclaration
	TemplateDeclaration   []*TemplateDeclaration   `xml:"templateDeclaration,omitempty"`   // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Root_AssessmentItem
	TemplateProcessing    *TemplateProcessing      `xml:"templateProcessing,omitempty"`    // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_templateProcessing
	AssessmentStimulusRef []*AssessmentStimulusRef `xml:"assessmentStimulusRef,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_assessmentStimulusRef
	Stylesheet            []*qti.StyleSheet        `xml:"stylesheet,omitempty"`            // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_stylesheet
	ItemBody              *ItemBody                `xml:"itemBody,omitempty"`              // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_itemBody
	ResponseProcessing    *ResponseProcessing      `xml:"responseProcessing,omitempty"`    // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_responseProcessing
	ModalFeedback         []ModalFeedback          `xml:"modalFeedback,omitempty"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_modalFeedback
	ApipAccessibility     *qti.ApipAccessibility   `xml:"apipAccessibility,omitempty"`     // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentItem_apipAccessibility
}

func NewAssessmentItem(b []byte) (*AssessmentItem, error) {
	var i AssessmentItem
	e := xml.Unmarshal(b, &i)
	if e != nil {
		return nil, e
	}

	return &i, nil
}

type ItemBody struct {
	XMLName xml.Name `xml:"itemBody" json:"-"`
	Xml     string   `xml:",innerxml"`
}

// ResponseDeclaration implements imsglobal ResponseDeclaration Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_ResponseDeclaration
type ResponseDeclaration struct {
	XMLName     xml.Name `xml:"responseDeclaration" json:"-"`
	Identifier  string   `xml:"identifier,attr"`         // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ResponseDeclaration.Attr_identifier
	Cardinality string   `xml:"cardinality,attr"`        // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ResponseDeclaration.Attr_cardinality
	BaseType    string   `xml:"baseType,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ResponseDeclaration.Attr_baseType

	DefaultValue    *qti.DefaultValue    `xml:"defaultValue,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_ResponseDeclaration_defaultValue
	CorrectResponse *qti.CorrectResponse `xml:"correctResponse"`        // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_ResponseDeclaration_correctResponse
	Mapping         *Mapping             `xml:"mapping,omitempty"`      // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_ResponseDeclaration_mapping
	AreaMapping     *AreaMapping         `xml:"areaMapping,omitempty"`  // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_ResponseDeclaration_areaMapping
}

// Mapping implements the imsglobal Mapping Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_Mapping
type Mapping struct {
	XMLName      xml.Name          `xml:"mapping" json:"-"`
	LowerBound   *float64          `xml:"lowerBound,attr,omitempty"`   // Optional, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Mapping.Attr_lowerBound
	UpperBound   *float64          `xml:"upperBound,attr,omitempty"`   // Optional, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Mapping.Attr_upperBound
	DefaultValue *qti.DefaultValue `xml:"defaultValue,attr,omitempty"` // Optional, Default 0. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Mapping.Attr_defaultValue
	MapEntry     []MapEntry        `xml:"mapEntry"`                    // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_Mapping_mapEntry
}

// MapEntry implements the imsglobal MapEntry Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Derived_MapEntry
type MapEntry struct {
	XMLName       xml.Name `xml:"mapEntry" json:"-"`
	MapKey        string   `xml:"mapKey,attr"`                  // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_MapEntry.Attr_mapKey
	MappedValue   float64  `xml:"mappedValue,attr"`             // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_MapEntry.Attr_mappedValue
	CaseSensitive bool     `xml:"caseSensitive,attr,omitempty"` // Optional. Default false. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_MapEntry.Attr_caseSensitive
}

// AreaMapping implements the imsglobal AreaMapping Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_AreaMapping
type AreaMapping struct {
	XMLName      xml.Name          `xml:"areaMapping" json:"-"`
	LowerBound   *float64          `xml:"lowerBound,attr,omitempty"`   // Optional, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_AreaMapping.Attr_lowerBound
	UpperBound   *float64          `xml:"upperBound,attr,omitempty"`   // Optional, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_AreaMapping.Attr_upperBound
	DefaultValue *qti.DefaultValue `xml:"defaultValue,attr,omitempty"` // Optional, Default 0. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_AreaMapping.Attr_defaultValue
	AreaMapEntry []AreaMapEntry    `xml:"mapEntry"`                    // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_AreaMapping_areaMapEntry
}

// AreaMapEntry implements the imsglobal AreaMapEntry Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_AreaMapping_areaMapEntry
type AreaMapEntry struct {
	XMLName     xml.Name   `xml:"areaMapEntry" json:"-"`
	Shape       qti.Shape  `xml:"shape,attr"`       // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_AreaMapEntry.Attr_shape
	Coords      qti.Coords `xml:"coords,attr"`      // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_AreaMapEntry.Attr_coords
	MappedValue float64    `xml:"mappedValue,attr"` // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_AreaMapEntry.Attr_mappedValue
}

// TemplateDeclaration implements imsglobal TemplateDeclaration Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_TemplateDeclaration
type TemplateDeclaration struct {
	XMLName       xml.Name          `xml:"templateDeclaration" json:"-"`
	Identifier    string            `xml:"identifier,attr"`              // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TemplateDeclaration.Attr_identifier
	Cardinality   string            `xml:"cardinality,attr"`             // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TemplateDeclaration.Attr_cardinality
	BaseType      string            `xml:"baseType,attr,omitempty"`      // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TemplateDeclaration.Attr_baseType
	ParamVariable bool              `xml:"paramVariable,attr,omitempty"` // Optional. Default false. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TemplateDeclaration.Attr_paramVariable
	MathVariable  bool              `xml:"mathVariable,attr,omitempty"`  // Optional. Default false. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_TemplateDeclaration.Attr_mathVariable
	DefaultValue  *qti.DefaultValue `xml:"defaultValue"`                 // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TemplateDeclaration_defaultValue
}

// TemplateProcessing implements imsglobal TemplateProcessing Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_TemplateProcessing
type TemplateProcessing struct {
	XMLName           xml.Name            `xml:"templateProcessing" json:"-"`
	TemplateRuleGroup []TemplateRuleGroup `xml:"templateRuleGroup"` // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_TemplateProcessing_templateRuleGroup
}

// TemplateRuleGroup implements imsglobal TemplateRuleGroup Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Abstract_TemplateRuleGroup
type TemplateRuleGroup struct {
	XMLName xml.Name `xml:"templateRuleGroup" json:"-"`
	// TODO: Build out the rest of TemplateRuleGroup
}

// AssessmentStimulusRef implements imsglobal AssessmentStimulusRef Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Derived_AssessmentStimulusRef
type AssessmentStimulusRef struct {
	XMLName    xml.Name `xml:"assessmentStimulusRef" json:"-"`
	Identifier string   `xml:"identifier,attr"` // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_AssessmentStimulusRef.Attr_identifier
	Href       string   `xml:"href,attr"`       // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_AssessmentStimulusRef.Attr_href
}

// ModalFeedback implements imsglobal ModalFeedback Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_ModalFeedback
type ModalFeedback struct {
	XMLName           xml.Name               `xml:"modalFeedback" json:"-"`
	OutcomeIdentifier string                 `xml:"outcomeIdentifier,attr"`      // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ModalFeedback.Attr_outcomeIdentifier
	ShowHide          string                 `xml:"showHide,attr"`               // Required. Enumerated value set of: { show | hide } https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ModalFeedback.Attr_showHide
	Identifier        string                 `xml:"identifier,attr"`             // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ModalFeedback.Attr_identifier
	Title             string                 `xml:"title,attr,omitempty"`        // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_ModalFeedback.Attr_title
	StyleSheet        []*qti.StyleSheet      `xml:"styleSheet,omitempty"`        // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_ModalFeedback_stylesheet
	ApipAccessibility *qti.ApipAccessibility `xml:"apipAccessibility,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_ModalFeedback_apipAccessibility

	// TODO: Finish ModalFeedback::feedbackFlowStaticGroup
}
