package assessmentresult

import (
	"encoding/xml"
	"fmt"
	"io"
)

// ItemResult. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#RootAttribute_AssessmentResult_itemResult
type ItemResult struct {
	XMLName          xml.Name      `xml:"itemResult"`
	InnerXML         string        `xml:",innerxml"`
	Identifier       string        `xml:"identifier,attr"`            // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ItemResult.Attr_identifier
	SequenceIndex    string        `xml:"sequenceIndex,attr"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ItemResult.Attr_sequenceIndex
	Datestamp        string        `xml:"datestamp,attr"`             // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ItemResult.Attr_datestamp
	SessionStatus    string        `xml:"sessionStatus,attr"`         // Required. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataCharacteristic_ItemResult.Attr_sessionStatus
	ItemVariables    []interface{} `xml:",omitempty"`                 // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataAttribute_ItemResult_itemVariable
	CandidateComment string        `xml:"candidateComment,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p1/QTIv2p2-Results-InfoBindModelv1p0p1/imsqtiv2p2_result_v1p0_InfoBindv1p0p1.html#DataAttribute_ItemResult_candidateComment
}

func (tr *ItemResult) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	fmt.Println("Running Custom ItemResult UnmarshalXML")
	// Grab ItemResult attributes first
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "identifier":
			tr.Identifier = attr.Value
		case "datestamp":
			tr.Datestamp = attr.Value
		case "sessionStatus":
			tr.SessionStatus = attr.Value
		case "sequenceIndex":
			tr.SequenceIndex = attr.Value
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
			case "responseVariable":
				if tr.ItemVariables == nil {
					tr.ItemVariables = make([]interface{}, 0)
				}
				var r ResponseVariable
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				tr.ItemVariables = append(tr.ItemVariables, r)
			case "outcomeVariable":
				if tr.ItemVariables == nil {
					tr.ItemVariables = make([]interface{}, 0)
				}
				var r OutcomeVariable
				err := d.DecodeElement(&r, &se)
				if err != nil {
					return err
				}
				tr.ItemVariables = append(tr.ItemVariables, r)
			default:
				fmt.Println("Unknown Tag Name: %s\n", se.Name.Local)
			}
		}
	}
	return nil
}
