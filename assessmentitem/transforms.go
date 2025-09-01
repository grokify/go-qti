package assessmentitem

import (
	"bytes"
	"encoding/xml"
	"path"
)

func (itm *ItemBody) ToWebComponent() ([]byte, error) {
	finalBytes := make([]byte, len(itm.Xml))
	parser := newItemBodyParser(bytes.NewReader([]byte(itm.Xml)))
	err := parser.parse()
	if err != nil {
		return nil, err
	}

	//var prevTok *tokVal
	for parser.stack.Size() > 0 {
		tokval := parser.stack.Pop().(*tokVal)
		switch tokval.tok {
		case node_choiceInteraction:
			tokval.val = []byte("qti-choiceInteraction")
		case node_prompt:
			tokval.val = []byte("qti-prompt")
		case node_simpleChoice:
			tokval.val = []byte("qti-simpleChoice")
		case node_textEntryInteraction:
			tokval.val = []byte("qti-textEntryInteraction")
		}

		finalBytes = append(finalBytes, tokval.val...)
	}

	return finalBytes, nil
}

type ItemWebComponent struct {
	XMLName       xml.Name `xml:"qti-itemBody"`
	Identifier    string   `xml:"identifier,attr"`
	Title         string   `xml:"title,attr"`
	Label         string   `xml:"label,attr"`
	Adaptive      bool     `xml:"adaptive,attr"`
	TimeDependent bool     `xml:"timeDependent,attr"`
	Xml           string   `xml:",innerxml"`
}

type ImageTag struct {
	Img xml.Name `xml:"img"`
	Src string   `xml:"src,attr"`
}

func (itm *ItemBody) UpdateImagePath(basePath string) ([]byte, error) {
	finalBytes := make([]byte, len(itm.Xml))
	parser := newItemBodyParser(bytes.NewReader([]byte(itm.Xml)))
	err := parser.parse()
	if err != nil {
		return nil, err
	}

	//var prevTok *tokVal
	for parser.stack.Size() > 0 {
		tokval := parser.stack.Pop().(*tokVal)
		switch tokval.tok {
		case node_img:
			// If Image tag, start building the image tag contents.
			imageBytes := make([]byte, 0)
			imageBytes = append(imageBytes, tokval.val...)
			for {
				imgVal := parser.stack.Pop().(*tokVal)
				// If this is an src tag, start processing the contents
				if imgVal.tok == node_src {
					imageBytes = append(imageBytes, imgVal.val...)

					// Find the first set of quotations for the value of src=
					for {
						preIdent := parser.stack.Pop().(*tokVal)
						if preIdent.tok == delim_dquote {
							imageBytes = append(imageBytes, preIdent.val...)
							break
						} else {
							imageBytes = append(imageBytes, preIdent.val...)
						}
					}

					// Build the original path until a non-escape double qoute is encountered.
					pathBytes := make([]byte, 0)
					for {
						ttok := parser.stack.Pop().(*tokVal)
						var prevTok *tokVal // used to detect if \ was the previous character for escaped double quote
						if ttok.tok == delim_dquote {
							if prevTok != nil && prevTok.tok == delim_bslash {
								// Keep building with the escaped quote
								pathBytes = append(pathBytes, ttok.val...)
							} else {
								// Put the non escaped double quote back on the stack.
								parser.stack.Push(ttok)
								break
							}
						} else {
							pathBytes = append(pathBytes, ttok.val...)
						}
						prevTok = ttok
					}

					fileName := path.Base(string(pathBytes))
					final := path.Join(basePath, fileName)
					imageBytes = append(imageBytes, final...)
					break
				} else {
					imageBytes = append(imageBytes, imgVal.val...)
				}
			}
			finalBytes = append(finalBytes, imageBytes...)
		default:
			// wasn't an image tag so just keep on going.
			finalBytes = append(finalBytes, tokval.val...)
		}

	}

	return finalBytes, nil
}
