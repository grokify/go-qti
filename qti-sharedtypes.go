package qti

import (
	"encoding/xml"
)

// DefaultValue implements the imsglobal DefaultValue Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_DefaultValue
type DefaultValue struct {
	XMLName        xml.Name `xml:"defaultValue" json:"-"`
	Interpretation string   `xml:"interpretation,attr,omitempty"` // Optional. Human readable interpretation of the default value. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_DefaultValue.Attr_interpretation
	Value          []Value  `xml:"value"`                         // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_DefaultValue_value
}

// Value implements the imsglobal Value Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Derived_Value
type Value struct {
	XMLName         xml.Name `xml:"value" json:"-"`
	FieldIdentifier string   `xml:"fieldIdentifier,attr,omitempty" json:"fieldIdentifier,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_Value.Attr_fieldIdentifier
	BaseType        string   `xml:"baseType,attr,omitempty" json:"baseType,omitempty"`               // Optional. Enumerated value set of: { boolean | directedPair | duration | file | float | identifier | integer | pair | point | string | uri } https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_Value.Attr_baseType
	Data            string   `xml:",chardata" json:"data,omitempty"`
}

// CorrectResponse implements imsglobal CorrectResponse Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_CorrectResponse
type CorrectResponse struct {
	XMLName        xml.Name `xml:"correctResponse" json:"-"`
	Interpretation string   `xml:"interpretation,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_CorrectResponse.Attr_interpretation
	Value          []Value  `xml:"value"`                    // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataAttribute_CorrectResponse_value
}

// OutcomeDeclaration implments the imsglobal 4.5 OutcomeDeclaration Root Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Root_OutcomeDeclaration
type OutcomeDeclaration struct {
	XMLName               xml.Name      `xml:"outcomeDeclaration" json:"-"`
	Identifier            string        `xml:"identifier,attr"`                 // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_identifier
	Cardinality           string        `xml:"cardinality,attr"`                // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_cardinality
	BaseType              string        `xml:"baseType,attr,omitempty"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_baseType
	View                  string        `xml:"view,attr,omitempty"`             // Optional. A list from an enumerated value set of: { author | candidate | proctor | scorer | testConstructor | tutor }. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_view
	Interpretation        string        `xml:"interpretation,attr,omitempty"`   // Optional. Human interpretation of the variable's value. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_interpretation
	LongInterpretation    string        `xml:"longInterpretation,attr"`         // Optional. URI to extended interpretation of the variable's value. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_longInterpretation
	NormalMaximum         float64       `xml:"normalMaximum,attr,omitempty"`    // Optional. Non-Negative. Maximum Magnitude. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_normalMaximum
	NormalMinimum         float64       `xml:"normalMinimum,attr,omitempty"`    // Optional. Minimum value of numeric outcome variables. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_normalMinimum
	MasteryValue          float64       `xml:"masteryValue,attr,omitempty"`     // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_masteryValue
	ExternalScored        string        `xml:"externalScored,attr,omitempty"`   // Optional. "externalMachine" | "human". Identifies what scores the outcome. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_externalScored
	VariableIdentifierRef string        `xml:"variableIdentifierRef,omitempty"` // Optional. Identifier for an external variable that will be used to provide the external scoring value. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_OutcomeDeclaration.Attr_variableIdentifierRef
	DefaultValue          *DefaultValue `xml:"defaultValue,omitempty"`          // Optional. Default: NULL value. default outcome value to be used when no matching tabel entry is found. If omitted, the NULL value is used. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_DefaultValue
	// TODO: Implement OutcomeDeclaration's LookupTable attribute
	//LookupTable           LookupTable  `xml:"lookupTable,omitempty"`           // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_OutcomeDeclaration_lookupTable
}

type RubricBlock struct {
	XMLName  xml.Name `xml:"rubricBlock" json:"-"`
	Xml      string   `xml:",innerxml"`
	Use      string   `xml:"use,attr,omitempty"`
	View     string   `xml:"view,attr"`
	Id       string   `xml:"id,attr,omitempty"`
	Class    string   `xml:"class,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Label    string   `xml:"label,attr,omitempty"`
	Base     string   `xml:"base,attr,omitempty"`
	Dir      string   `xml:"dir,attr,omitempty"` //  ltr | rtl | auto, Default auto
	Role     string   `xml:"role,attr,omitempty"`
	// TODO: Implement Aria characteristics
	// TODO: Finish Rubric Block
}

// Precondition. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_LogicSingle
type PreCondition struct {
	XMLName xml.Name `xml:"preCondition" json:"-"`
	Logic   string   `xml:"innerxml"`
}

// StyleSheet implements imsglobal StyleSheet Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Derived_StyleSheet
type StyleSheet struct {
	XMLName xml.Name `xml:"stylesheet" json:"-"`
	Href    string   `xml:"href,attr"`            // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_StyleSheet.Attr_href
	Type    string   `xml:"type,attr"`            // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_StyleSheet.Attr_type
	Media   string   `xml:"media,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_StyleSheet.Attr_media
	Title   string   `xml:"title,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedCharacteristic_StyleSheet.Attr_title
}

// ApipAccessibility implements imsglobal APIPAccessibility Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Imported_APIPAccessibility
type ApipAccessibility struct {
	XMLName xml.Name `xml:"apipAccessibility" json:"-"`
	// TODO: Finished ApipAccessbility
}

// Shape implements the imsglobal Shape Class. Only one of the following attributes should be specified to indicate the shape and the string contains the shape's properties. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Enumerated_Shape
type Shape struct {
	XMLName xml.Name `xml:"shape" json:"-"`
	Circle  string   `xml:"circle,attr"`
	Default string   `xml:"default,attr"` // The default shape refers to the entire area of the associated image.
	Ellipse string   `xml:"ellipse,attr"`
	Poly    string   `xml:"poly,attr"`
	Rect    string   `xml:"rect,attr"`
}

// Coords  implements the imsglobal Coords Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Derived_Coords
type Coords struct {
	XMLName xml.Name `xml:"coords" json:"-"`
	Pattern string   `xml:"pattern"` // Required. Default:(([0-9]+%?[,]){2}([0-9]+%?))|(([0-9]+%?[,]){3}([0-9]+%?))|(([0-9]+%?[,]){2}(([0-9]+%?[,]){2})+([0-9]+%?[,])([0-9]+%?)). https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DerivedAttribute_Coords_pattern
}

// Expression Group. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Abstract_ExpressionGroup
type ExpressionGroup struct {
	XMLName xml.Name `xml:"expressionGroup" json:"-"`
	Xml     string   `xml:",innerxml"`
}

var (
	ResponseTemplates map[string]string // Key is the URL, string in the XML
)

func init() {
	ResponseTemplates = make(map[string]string, 0)
	ResponseTemplates["http://www.imsglobal.org/question/qti_v2p2/rptemplates/match_correct"] = responseTemplateMatchCorrect
	ResponseTemplates["https://www.imsglobal.org/question/qti_v2p2/rptemplates/match_correct"] = responseTemplateMatchCorrect
	ResponseTemplates["http://www.imsglobal.org/question/qtiv2p2/rptemplates/map_response.xml"] = responseTemplateMatchCorrect
	ResponseTemplates["https://www.imsglobal.org/question/qtiv2p2/rptemplates/map_response.xml"] = responseTemplateMatchCorrect
	ResponseTemplates["http://www.imsglobal.org/question/qtiv2p2/rptemplates/map_response_point.xml"] = responseTemplateMapResponsePoint
	ResponseTemplates["https://www.imsglobal.org/question/qtiv2p2/rptemplates/map_response_point.xml"] = responseTemplateMapResponsePoint
}

var responseTemplateMatchCorrect = `<responseProcessing xsi:schemaLocation="http://www.imsglobal.org/xsd/imsqti_v2p2 http://www.imsglobal.org/xsd/qti/qtiv2p2/imsqti_v2p2.xsd" xmlns="http://www.imsglobal.org/xsd/imsqti_v2p2" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"> 
    <responseCondition>
        <responseIf>
            <match>
                <variable identifier="RESPONSE"/>
                <correct identifier="RESPONSE"/>
            </match>
            <setOutcomeValue identifier="SCORE">
                <baseValue baseType="float">1</baseValue>
            </setOutcomeValue>
        </responseIf>
        <responseElse>
            <setOutcomeValue identifier="SCORE">
                <baseValue baseType="float">0</baseValue>
            </setOutcomeValue>
        </responseElse>
    </responseCondition>
</responseProcessing>
`
var responseTemplateMapResponse = `
<responseProcessing xsi:schemaLocation="http://www.imsglobal.org/xsd/imsqti_v2p2 http://www.imsglobal.org/xsd/qti/qtiv2p2/imsqti_v2p2.xsd">
    <responseCondition>
        <responseIf>
            <isNull>
                <variable identifier="RESPONSE"/>
            </isNull>
            <setOutcomeValue identifier="SCORE">
                <baseValue baseType="float">0.0</baseValue>
            </setOutcomeValue>
        </responseIf>
        <responseElse>
            <setOutcomeValue identifier="SCORE">
                <mapResponse identifier="RESPONSE"/>
            </setOutcomeValue>
        </responseElse>
    </responseCondition>
</responseProcessing>`

var responseTemplateMapResponsePoint = `
<responseProcessing xsi:schemaLocation="http://www.imsglobal.org/xsd/imsqti_v2p2 http://www.imsglobal.org/xsd/qti/qtiv2p2/imsqti_v2p2.xsd">
    <responseCondition>
        <responseIf>
            <isNull>
                <variable identifier="RESPONSE"/>
            </isNull>
            <setOutcomeValue identifier="SCORE">
                <baseValue baseType="float">0</baseValue>
            </setOutcomeValue>
        </responseIf>
        <responseElse>
            <setOutcomeValue identifier="SCORE">
                <mapResponsePoint identifier="RESPONSE"/>
            </setOutcomeValue>
        </responseElse>
    </responseCondition>
</responseProcessing>`
