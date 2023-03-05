package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"io"
	"strings"
)

const (
	TEXT = iota
	PIN
)

type MDParser struct {
	cells        [][2]int // (cellType, extIndex)
	txtSegments  []*ast.CommentGroup
	declSegments []ast.Decl
	pins         []ast.Node
	File         []byte
}

func (p *MDParser) ParseComments(c []*ast.CommentGroup) error {
	for _, tk := range c {
		if strings.HasPrefix(tk.Text(), "PIN") {
			p.pins = append(p.pins, tk)
			p.cells = append(p.cells, [2]int{PIN, len(p.pins) - 1})
		} else if strings.HasPrefix(tk.Text(), "POST") {
			p.txtSegments = append(p.txtSegments, tk)
			p.cells = append(p.cells, [2]int{TEXT, len(p.txtSegments) - 1})
		}
	}

	return nil
}

func (p *MDParser) ParseDeclarations(decl []ast.Decl) error {
	for _, tk := range decl {
		for _, v := range p.pins {
			if v.End()+1 == tk.Pos() {
				p.declSegments = append(p.declSegments, tk)
			}
		}
	}

	return nil
}

func (p MDParser) Flush() io.Reader {
	s := "# New file\n\n"

	for _, cell := range p.cells {
		switch cell[0] {
		case 0:
			s += fmt.Sprintf("```go\n%s\n```\n\n", string(p.File[p.pins[cell[1]].End():p.declSegments[cell[1]].End()]))
		default:
			s += strings.TrimPrefix(p.txtSegments[cell[1]].Text(), "POST\n") + "\n"
		}
	}
	return bytes.NewBufferString(s)
}
