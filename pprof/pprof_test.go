package rand

import (
	"os"
	"runtime/pprof"
	"testing"
)

func TestPProf(t *testing.T) {
	p := pprof.Lookup("heap")
	p.WriteTo(os.Stdout, 3)
}