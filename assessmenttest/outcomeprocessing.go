package assessmenttest

import (
	"encoding/xml"
	"fmt"
	"github.com/stmath/go-xmldom"
)

// ResponseProcessing implements imsglobal ResponseProcessing Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Root_ResponseProcessing
type OutcomeProcessing struct {
	XMLName             xml.Name       `xml:"outcomeProcessing"`
	Template            string         `xml:"template,attr,omitempty"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_ResponseProcessing.Attr_template
	TemplateLocation    string         `xml:"templateLocation,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_ResponseProcessing.Attr_templateLocation
	OutcomeRuleGroupXML string         `xml:",innerxml"`                       //
	OutcomeRuleGroup    []*xmldom.Node `xml:"-" json:"-"`
}
type OutcomeRuleGroup struct {
	XML string `xml:",innerxml"`
}

func (op *OutcomeProcessing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	fmt.Println("Running Custom assessmentTest.outcomeProcessing UnmarshalXML")
	// Grab ItemResult attributes first
	op.XMLName = start.Name

	var org OutcomeRuleGroup
	err := d.DecodeElement(&org, &start)
	if err != nil {
		return err
	}
	op.OutcomeRuleGroupXML = org.XML
	dom, err := xmldom.ParseXML("<outcomeProcessing>" + op.OutcomeRuleGroupXML + "</outcomeProcessing>")
	if err != nil {
		return err
	}
	op.OutcomeRuleGroup = dom.Root.Children

	return nil
}
