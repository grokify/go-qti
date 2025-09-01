package assessmentitem

import (
	"encoding/xml"
	"fmt"
	"github.com/stmath/go-xmldom"
)

// ResponseProcessing implements imsglobal ResponseProcessing Class. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#Root_ResponseProcessing
type ResponseProcessing struct {
	XMLName              xml.Name       `xml:"responseProcessing" json:"-"`
	Template             string         `xml:"template,attr,omitempty"`         // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_ResponseProcessing.Attr_template
	TemplateLocation     string         `xml:"templateLocation,attr,omitempty"` // Optional. https://www.imsglobal.org/question/qtiv2p2p2/QTIv2p2p2-ASI-InformationModelv1p0/imsqtiv2p2p2_asi_v1p0_InfoModelv1p0.html#RootCharacteristic_ResponseProcessing.Attr_templateLocation
	ResponseRuleGroupXML string         `xml:",innerxml"`                       //
	ResponseRuleGroup    []*xmldom.Node `xml:"-" json:"-"`
}
type ResponseRuleGroup struct {
	XML string `xml:",innerxml"`
}

func (rp *ResponseProcessing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	//fmt.Println("Running Custom assessmentItem.responseProcessing UnmarshalXML")
	// Grab ItemResult attributes first
	rp.XMLName = start.Name

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "template":
			rp.Template = attr.Value
		case "templateLocation":
			rp.TemplateLocation = attr.Value
		default:
			fmt.Printf("Unkown Tag Attribute %s[@%s]", start.Name.Local, attr.Name.Local)
		}
	}
	var rrg ResponseRuleGroup
	err := d.DecodeElement(&rrg, &start)
	if err != nil {
		return err
	}

	if rrg.XML != "" {
		rp.ResponseRuleGroupXML = rrg.XML
		dom, err := xmldom.ParseXML("<responseProcessing>" + rp.ResponseRuleGroupXML + "</responseProcessing>")
		if err != nil {
			return err
		}

		rp.ResponseRuleGroup = dom.Root.Children
	}

	return nil
}
