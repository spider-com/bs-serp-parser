package serp

import (
	"io"
)

type ParseJSON func(r io.Reader) ([]byte, error)