package assessmentitem

import (
	"io"

	"github.com/alediaferia/stackgo"
)

type tokVal struct {
	tok token
	val []byte
}
type token int

type itemBodyParser struct {
	s   *itemBodyScanner
	buf struct {
		tok token
		lit []byte
		n   int
	}

	stack *stackgo.Stack
}

func newItemBodyParser(r io.Reader) *itemBodyParser {
	return &itemBodyParser{s: newQtiScanner(r)}
}

func (ibp *itemBodyParser) scan() (tok token, lit []byte, err error) {
	// If we have a token on the buffer, then return it.
	if ibp.buf.n != 0 {
		ibp.buf.n = 0
		return ibp.buf.tok, ibp.buf.lit, nil
	}

	// Otherwise read the next token from the scanner.
	tok, lit, err = ibp.s.scan()
	if err != nil {
		return
	}

	// Save it to the buffer in case we unscan later.
	ibp.buf.tok, ibp.buf.lit = tok, lit

	return
}

func (ibp *itemBodyParser) unscan() { ibp.buf.n = 1 }

func (ibp *itemBodyParser) parse() error {
	mode := 0
	flipStack := stackgo.NewStack()

	for mode == 0 {
		tok, lit, err := ibp.scan()
		if err != nil {
			if err != io.EOF {
				return err
			} else {
				mode = 1
			}
		} else {
			token := &tokVal{tok: tok, val: lit}
			flipStack.Push(token)
		}
	}

	ibp.stack = stackgo.NewStackWithCapacity(flipStack.Size() + 1)
	for flipStack.Size() > 1 {
		ibp.stack.Push(flipStack.Pop())
	}

	return nil
}
