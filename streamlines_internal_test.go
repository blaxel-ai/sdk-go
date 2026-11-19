package blaxel

import (
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

func scanAll(s *streamLineScanner) []string {
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func TestProcessStreamLineScannerLines(t *testing.T) {
	long := strings.Repeat("x", 300_000)
	// OneByteReader makes every read short, as a slow network does.
	s := newStreamLineScanner(iotest.OneByteReader(strings.NewReader("a\r\n\n" + long + "\nlast")))
	lines := scanAll(s)
	if s.Err() != nil {
		t.Fatalf("Err: %v", s.Err())
	}
	if len(lines) != 4 || lines[0] != "a" || lines[1] != "" || lines[2] != long || lines[3] != "last" {
		t.Fatalf("got %d lines", len(lines))
	}
}

func TestProcessStreamLineScannerStopsAtTheLineCap(t *testing.T) {
	s := newStreamLineScanner(strings.NewReader("short\n" + strings.Repeat("x", 200) + "\nafter\n"))
	s.max = 100
	lines := scanAll(s)
	if len(lines) != 1 || lines[0] != "short" {
		t.Fatalf("lines=%q", lines)
	}
	if s.Err() == nil || !strings.Contains(s.Err().Error(), "longer than 100 bytes") {
		t.Fatalf("Err=%v", s.Err())
	}
	if s.Scan() {
		t.Fatal("Scan returned true after an error")
	}
}

func TestProcessStreamLineScannerReturnsReadErrors(t *testing.T) {
	s := newStreamLineScanner(io.MultiReader(strings.NewReader("one\npartial"), iotest.ErrReader(io.ErrUnexpectedEOF)))
	lines := scanAll(s)
	if len(lines) != 1 || lines[0] != "one" || !errors.Is(s.Err(), io.ErrUnexpectedEOF) {
		t.Fatalf("lines=%q Err=%v", lines, s.Err())
	}
}
