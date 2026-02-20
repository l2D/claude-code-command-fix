package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func mockClipboardOK(_ string) error { return nil }

func mockClipboardFail(_ string) error { return fmt.Errorf("clipboard unavailable") }

func TestFormatAndDisplay(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		clipFn     ClipboardFunc
		wantSubstr []string
	}{
		{
			name:   "formats and copies",
			input:  "ls -la\n  | grep test",
			clipFn: mockClipboardOK,
			wantSubstr: []string{
				separator,
				"ls -la | grep test",
				"Copied to clipboard!",
				"#4686",
			},
		},
		{
			name:   "empty input",
			input:  "",
			clipFn: mockClipboardOK,
			wantSubstr: []string{
				"No command text provided.",
			},
		},
		{
			name:   "clipboard failure",
			input:  "echo hello",
			clipFn: mockClipboardFail,
			wantSubstr: []string{
				"echo hello",
				"Could not copy to clipboard",
				"copy the command above manually",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			formatAndDisplay(&buf, tt.input, false, tt.clipFn)
			got := buf.String()

			for _, want := range tt.wantSubstr {
				if !strings.Contains(got, want) {
					t.Errorf("output missing %q\ngot:\n%s", want, got)
				}
			}
		})
	}
}

func TestInteractiveMode(t *testing.T) {
	input := "command1   &&    command2\n\n"
	var buf bytes.Buffer

	interactiveMode(strings.NewReader(input), &buf, false, mockClipboardOK)
	got := buf.String()

	if !strings.Contains(got, "command1 && command2") {
		t.Errorf("expected formatted command in output, got:\n%s", got)
	}
	if !strings.Contains(got, "Copied to clipboard!") {
		t.Errorf("expected clipboard success message, got:\n%s", got)
	}
}
