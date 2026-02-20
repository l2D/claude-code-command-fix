package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/l2D/claude-code-command-fix/internal/clipboard"
	"github.com/l2D/claude-code-command-fix/internal/formatter"
	"github.com/l2D/claude-code-command-fix/internal/version"
)

const separator = "════════════════════════════════════════════════════════════"

// ClipboardFunc is the function signature for clipboard operations.
type ClipboardFunc func(string) error

// formatAndDisplay formats the input, prints the result, and copies to clipboard.
//
//nolint:errcheck // fmt writes to stdout/buffer; errors are not actionable in a CLI
func formatAndDisplay(w io.Writer, input string, clipFn ClipboardFunc) {
	formatted := formatter.FormatCommand(input)
	if formatted == "" {
		fmt.Fprintln(w, "No command text provided.")
		return
	}

	fmt.Fprintln(w, separator)
	fmt.Fprintln(w, formatted)
	fmt.Fprintln(w, separator)

	if err := clipFn(formatted); err != nil {
		fmt.Fprintf(w, "Could not copy to clipboard: %v\n", err)
		fmt.Fprintln(w, "Please copy the command above manually.")
	} else {
		fmt.Fprintln(w, "Copied to clipboard!")
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "TEMPORARY FIX: This tool addresses Claude Code issue #4686")
	fmt.Fprintln(w, "https://github.com/anthropics/claude-code/issues/4686")
}

// interactiveMode reads multi-line input until an empty line is entered.
//
//nolint:errcheck // fmt writes to stdout/buffer; errors are not actionable in a CLI
func interactiveMode(r io.Reader, w io.Writer, clipFn ClipboardFunc) {
	fmt.Fprintln(w, "Paste the command below (press Enter twice to process):")
	fmt.Fprintln(w)

	scanner := bufio.NewScanner(r)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	input := strings.Join(lines, "\n")
	formatAndDisplay(w, input, clipFn)
}

// Run is the main entry point for the CLI.
func Run() {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "--version" {
		fmt.Printf("claude-code-command-fix %s (commit: %s, built: %s)\n",
			version.Version, version.CommitSHA, version.BuildTime)
		return
	}

	if len(args) > 0 {
		input := strings.Join(args, " ")
		formatAndDisplay(os.Stdout, input, clipboard.Copy)
		return
	}

	interactiveMode(os.Stdin, os.Stdout, clipboard.Copy)
}
