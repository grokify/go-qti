package assessmentitem

import "bytes"

func (itm *ItemBody) MathML() ([]string, error) {
	fin := make([]string, 0)

	finalBytes := make([]byte, len(itm.Xml))
	parser := newItemBodyParser(bytes.NewReader([]byte(itm.Xml)))
	err := parser.parse()
	if err != nil {
		return nil, err
	}

	for parser.stack.Size() > 0 {
		tokval := parser.stack.Pop().(*tokVal)
		switch tokval.tok {
		case node_math:
			var mxml string
			mxml += "<" + string(tokval.val)
			var prevTok *tokVal
			// First collect the entire math section
			for parser.stack.Size() > 0 {
				mtok := parser.stack.Pop().(*tokVal)
				if mtok.tok == node_math && prevTok != nil && prevTok.tok == delim_fslash {
					mxml += string(mtok.val)
					for parser.stack.Size() > 0 {
						closeTok := parser.stack.Pop().(*tokVal)
						if closeTok.tok != delim_gthan {
							mxml += string(closeTok.val)
						} else {
							mxml += string(closeTok.val)
							break
						}
					}

					fin = append(fin, mxml)
					break
				} else {
					mxml += string(mtok.val)
					//fmt.Printf("%s\n", mxml)
				}
				prevTok = mtok
			}
		default:
			// wasn't a math tag so keep going
			finalBytes = append(finalBytes, tokval.val...)
		}
	}

	return fin, nil
}
