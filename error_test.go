// Go binding for nanomsg

package nanomsg

import (
	"syscall"
	"testing"
)

func TestError(t *testing.T) {
	var err error
	var s *Socket
	if s, err = NewSocket(AF_SP, REP); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Bind("inproc://a"); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Bind("inproc://a"); err == nil {
		t.Fatal("expected failure")
	} else {
		if err != syscall.EADDRINUSE {
			t.Fatal(err)
		}
	}
	if err = s.Close(); err != nil {
		t.Fatal(err)
	}
}
