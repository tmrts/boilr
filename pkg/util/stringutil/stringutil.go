package stringutil

import "io"

// String wraps string with io.ReadWriter
type String interface {
	io.ReadWriter
	String() string
}

// NewString creates a String
func NewString(contents string) String {
	s := str(contents)
	return &s
}

type str string

func (s *str) Read(buf []byte) (int, error) {
	ss := string(*s)
	for i := range buf {
		buf[i] = byte(ss[i])
	}

	return len(buf), nil
}

func (s *str) Write(buf []byte) (int, error) {
	*s = str(string(*s) + string(buf))
	return len(buf), nil
}

func (s *str) String() string {
	return string(*s)
}
