package assessmentitem

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const (
	illegal token = iota
	ident
	ctrl_eof
	delim_lthan
	delim_gthan
	delim_ws
	delim_nl
	delim_tab
	delim_cr
	delim_dquote
	delim_equals
	delim_fslash
	delim_bslash

	node_ident
	node_choiceInteraction
	node_textEntryInteraction
	node_prompt
	node_simpleChoice
	node_eof
	node_a
	node_img
	node_src
	node_math
)

func isControl(ch byte) bool {
	return ch == '<' || ch == '>' || ch == '\n' || ch == ' ' || ch == '/' || ch == '\t' || ch == '=' || ch == '"' || ch == '\\'
}
func isIdent(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '.' || ch == '-' || ch == '+' || ch == ':' || ch == '#' || ch == ',' || ch == '_' || ch == '&' || ch == ';' || ch == '*' || ch == '@' || ch == '!' || ch == '~' || ch == '?'
}

type itemBodyScanner struct {
	r *bufio.Reader
}

func newQtiScanner(r io.Reader) *itemBodyScanner {
	return &itemBodyScanner{r: bufio.NewReader(r)}
}

func (s *itemBodyScanner) unread() error {
	err := s.r.UnreadByte()
	return err
}

func (s *itemBodyScanner) scan() (tok token, lit []byte, err error) {
	ch, err := s.r.ReadByte()
	if err != nil {
		return ctrl_eof, []byte(""), err
	}
	if isControl(ch) {
		s.unread()
		return s.scanDeliminators()
	} else {
		s.unread()
		return s.scanIdent()
	}
}

func (s *itemBodyScanner) scanDeliminators() (tok token, lit []byte, err error) {
	var buf bytes.Buffer
	ch, err := s.r.ReadByte()
	if err != nil {
		return ctrl_eof, []byte(""), err
	}
	buf.WriteByte(ch)

	switch strings.ToUpper(buf.String()) {
	case "<":
		return delim_lthan, []byte(buf.String()), nil
	case ">":
		return delim_gthan, []byte(buf.String()), nil
	case " ":
		return delim_ws, []byte(buf.String()), nil
	case "\n":
		return delim_nl, []byte(buf.String()), nil
	case "\t":
		return delim_tab, []byte(buf.String()), nil
	case "\r":
		return delim_cr, []byte(buf.String()), nil
	case "\"":
		return delim_dquote, []byte(buf.String()), nil
	case "=":
		return delim_equals, []byte(buf.String()), nil
	case "/":
		return delim_fslash, []byte(buf.String()), nil
	case "\\":
		return delim_bslash, []byte(buf.String()), nil
	}

	return node_ident, []byte(buf.String()), nil
}

func (s *itemBodyScanner) scanIdent() (tok token, lit []byte, err error) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	ch, err := s.r.ReadByte()
	if err != nil {
		return ctrl_eof, []byte(""), err
	}

	buf.WriteByte(ch)

	//Read every subsequent whitespace character into the buffer.
	//Non-whitespace characters and ctrl_eof will cause the loop to exit.
	for {
		if ch, err := s.r.ReadByte(); err != nil {
			s.unread()
			break
		} else if isControl(ch) {
			s.unread()
			break
		} else {
			buf.WriteByte(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch buf.String() {
	case "src":
		return node_src, []byte(buf.String()), nil
	case "img":
		return node_img, []byte(buf.String()), nil
	case "a":
		return node_a, []byte(buf.String()), nil
	case "choiceInteraction":
		return node_choiceInteraction, []byte(buf.String()), nil
	case "prompt":
		return node_prompt, []byte(buf.String()), nil
	case "simpleChoice":
		return node_simpleChoice, []byte(buf.String()), nil
	case "math":
		fallthrough
	case "m:math":
		return node_math, []byte(buf.String()), nil
	case "textEntryInteraction":
		return node_textEntryInteraction, []byte(buf.String()), nil
	case "EOF":
		return node_eof, []byte(buf.String()), nil
	}

	// Otherwise return as a regular identifier.
	return ident, []byte(buf.String()), nil
}
