package rand

import (
	"os"
	"runtime/pprof"
	"testing"
)

// 参考文章 {@link : https://juejin.im/entry/5ac9cf3a518825556534c76e}
func TestPProf(t *testing.T) {
	p := pprof.Lookup("heap")
	p.WriteTo(os.Stdout, 3)
}