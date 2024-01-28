package interprog

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestGeom(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("InterProg", func() {
		g.It("sample inter prog", func() {
			g.Timeout(1 * time.Hour)
			var t0 = time.Now()
			var prog = NewInterProg("progress")
			for i := 0; i < 10000; i++ {
				time.Sleep(1 * time.Millisecond)
				prog.Update()
			}
			prog.Done()
			var delta = time.Since(t0)
			g.Assert(delta.Seconds() > 10).IsTrue()
		})

	})
}
