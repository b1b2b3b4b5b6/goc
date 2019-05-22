package dumptool

import (
	"io"

	"github.com/davecgh/go-spew/spew"
)

func Dump(args ...interface{}) {
	spew.Dump(args)
}

func Sdump(args ...interface{}) string {
	return spew.Sdump(args)
}

func Fdump(w io.Writer, args ...interface{}) {
	spew.Fdump(w, args)
}
