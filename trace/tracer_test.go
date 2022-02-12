package trace

import (
	"bytes"
	"testing"
)

func Test_Tracer(t *testing.T) {
	t.Run("Testing Tracer", func(t *testing.T) {
		var buf bytes.Buffer
		tracer := NewTracer(&buf)
		if tracer == nil {
			t.Fatal("Return from NewTracer should not be nil")
		}
		msg := "Hello trace package"
		tracer.Trace(msg)
		if buf.String() != msg+"\n" {
			t.Errorf("want :%s, got: %s", msg, buf.String())
		}
	})
}
