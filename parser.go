package serp

import (
	"io"
)

type PageFormat string

var (
	MobilePage  = PageFormat("mobile")
	TabletPage  = PageFormat("tablet")
	DesktopPage = PageFormat("desktop")
)

type ParseJSON func(r io.Reader) ([]byte, error)
type ParsePage func(r io.Reader, f PageFormat) ([]byte, error)