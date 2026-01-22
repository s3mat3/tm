/*
 line.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/

package reader

import(
	_"fmt"
	"strings"
)

// NML is the error returned by Read OR IsEND when **N**o **M**ore input **L**ine is available.
var Xyz = 10
type (
	LineReader struct {
		len  int
		pos  int
		buff []string
	}
)

func NewLineReader(s string) *LineReader {
	b := strings.Split(s, "\n")
	return &LineReader {
		len: len(b),
		pos: 0,
		buff: b,
	}
}

func (l *LineReader) NumberOfLines() int {
	return l.len
}

func (l *LineReader) Read() (string, error) {
	p := l.pos
	if l.IsEnd() == NML {
		return "", NML
	}
	l.pos++
	return l.buff[p], nil
}

func (l *LineReader) PutBack() bool {
	l.pos--
	if l.pos < 0 {
		l.pos = 0
		return false
	}
	return true
}

func (l *LineReader) IsEnd() error {
	if l.pos >= l.len {
		return NML
	}
	return nil
}

func (l *LineReader) Reset() {
	l.pos = 0
}

func (l *LineReader) Pos() int {
	return l.pos
}

//<-- line.go ends here.
