package assessmenttest

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/grokify/go-qti"
)

// AssessmentSection groups together individual item references or sub sections. GetSectionPart() must be run to populate the items and sub-sections. Implements https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Root_AssessmentSection
type AssessmentSection struct {
	XMLName      xml.Name `xml:"assessmentSection"`
	Xml          string   `xml:",innerxml"`                   // Used during processing to generate the interface object list since AssessmentSection is a container for other elements.
	Identifier   string   `xml:"identifier,attr"`             // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentSection.Attr_identifier
	Required     bool     `xml:"required,attr,omitempty"`     // Optional, Default False, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentSection.Attr_required
	Fixed        bool     `xml:"fixed,attr,omitempty"`        // Optional, Default False, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentSection.Attr_fixed
	Title        string   `xml:"title,attr"`                  // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentSection.Attr_title
	Visible      bool     `xml:"visible,attr"`                // Required. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentSection.Attr_visible
	KeepTogether bool     `xml:"keepTogether,attr,omitempty"` // Optional, Default True, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_AssessmentSection.Attr_keepTogether

	PreCondition       []*qti.PreCondition `xml:"preCondition",omitempty`
	BranchRule         []*BranchRule       `xml:"branchRule,omitempty"`
	ItemSessionControl *ItemSessionControl `xml:"itemSessionControl,omitempty"`
	TimeLimits         *TimeLimits         `xml:"timeLimits,omitempty"`
	Selection          *Selection          `xml:"selection,omitempty"`
	Ordering           *Ordering           `xml:"ordering,omitempty"`
	RubricBlock        []*qti.RubricBlock  `xml:"rubricBlock,omitempty"`
	SectionPart        []interface{}       // Abstract Container, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootAttribute_AssessmentSection_sectionPart
}

// SetDefault Sets the default values according to the imsglobal spec
func (as *AssessmentSection) SetDefaults() {
	as.KeepTogether = true
}

// UnmarshalXML is used to handle the SectionPart Container
func (as *AssessmentSection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	fmt.Println("Running Custom assessmentTest.testPart.assessmentSection UnmarshalXML")
	// Grab ItemResult attributes first
	as.XMLName = start.Name

	// Set Defaults
	as.KeepTogether = true

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "identifier":
			as.Identifier = attr.Value
		case "required":
			err := parseXmlBool(attr.Value, &as.Required)
			if err != nil {
				return err
			}
		case "fixed":
			err := parseXmlBool(attr.Value, &as.Fixed)
			if err != nil {
				return err
			}
		case "title":
			as.Title = attr.Value
		case "visible":
			err := parseXmlBool(attr.Value, &as.Visible)
			if err != nil {
				return err
			}
		case "keepTogether":
			err := parseXmlBool(attr.Value, &as.KeepTogether)
			if err != nil {
				return err
			}
		default:
			fmt.Printf("Unkown Tag Attribute %s[@%s]", start.Name.Local, attr.Name.Local)
		}
	}
	// Get the Next token and build the ItemVariables
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}

		switch se := t.(type) {
		case xml.EndElement:
			return nil
		case xml.StartElement:

			switch se.Name.Local {
			case "preCondition":
				if as.PreCondition == nil {
					as.PreCondition = make([]*qti.PreCondition, 0)
				}
				var r qti.PreCondition
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.PreCondition = append(as.PreCondition, &r)
			case "branchRule":
				if as.BranchRule == nil {
					as.BranchRule = make([]*BranchRule, 0)
				}
				var r BranchRule
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.BranchRule = append(as.BranchRule, &r)
			case "itemSessionControl":
				var r ItemSessionControl
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.ItemSessionControl = &r
			case "timeLimits":
				var r TimeLimits
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.TimeLimits = &r
			case "selection":
				var r Selection
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.Selection = &r
			case "ordering":
				var r Ordering
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.Ordering = &r
			case "rubricBlock":
				if as.RubricBlock == nil {
					as.RubricBlock = make([]*qti.RubricBlock, 0)
				}
				var r qti.RubricBlock
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.RubricBlock = append(as.RubricBlock, &r)
			case "include":
				fmt.Println("assessmentTest.testPart.assessmentSection.SectionPart.include not supported.")
			case "assessmentItemRef":
				if as.SectionPart == nil {
					as.SectionPart = make([]interface{}, 0)
				}
				var r AssessmentItemRef
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.SectionPart = append(as.SectionPart, &r)
			case "assessmentSection":
				if as.SectionPart == nil {
					as.SectionPart = make([]interface{}, 0)
				}
				var r AssessmentSection
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.SectionPart = append(as.SectionPart, &r)
			case "assessmentSectionRef":
				if as.SectionPart == nil {
					as.SectionPart = make([]interface{}, 0)
				}
				var r AssessmentSectionRef
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				as.SectionPart = append(as.SectionPart, &r)
			default:
				fmt.Printf("Unknown Tag Name: %s\n", se.Name.Local)
			}
		}
	}
	return nil
}

// Selection implements imsglobal Selection Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_Selection
type Selection struct {
	XMLName         xml.Name `xml:"selection"`
	Select          int      `xml:"select,attr"`                    // Required. Number of Elements to be selected. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Selection.Attr_select
	WithReplacement bool     `xml:"withReplacement,attr,omitempty"` // Optional. Default False. Allow items to be selected multiple times. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Selection.Attr_withReplacement
	//TODO Make extension support
}

func (s *Selection) SetDefaults() {
	s.WithReplacement = false
}

// Ordering implements the imsglobal class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Data_Ordering
type Ordering struct {
	XMLName    xml.Name `xml:"ordering"`
	Shuffle    bool     `xml:"shuffle,attr,omitempty"`    // Optional, Default False, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Ordering.Attr_shuffle
	Extension  string   `xml:"extension,attr,omitempty"`  // Optional, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Ordering.Attr_extension
	Extensions []string `xml:"extensions,attr,omitempty"` // Optional, https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#DataCharacteristic_Ordering.Attr_extension
}

func parseXmlBool(src string, dst *bool) error {
	if len(src) == 0 {
		*dst = false
		return nil
	}

	value, err := strconv.ParseBool(strings.TrimSpace(string(src)))
	if err != nil {
		return err
	}

	*dst = value
	return nil
}
