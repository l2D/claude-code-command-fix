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
func formatAndDisplay(w io.Writer, input string, singleLine bool, clipFn ClipboardFunc) {
	formatted := formatter.FormatCommand(input, singleLine)
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
func interactiveMode(r io.Reader, w io.Writer, singleLine bool, clipFn ClipboardFunc) {
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
	formatAndDisplay(w, input, singleLine, clipFn)
}

// Run is the main entry point for the CLI.
func Run() {
	args := os.Args[1:]

	var singleLine bool
	var remaining []string

	for _, arg := range args {
		switch arg {
		case "--help", "-h":
			printUsage(os.Stdout)
			return
		case "--version":
			fmt.Printf("claude-code-command-fix %s (commit: %s, built: %s)\n",
				version.Version, version.CommitSHA, version.BuildTime)
			return
		case "--single-line", "-s":
			singleLine = true
		default:
			remaining = append(remaining, arg)
		}
	}

	if len(remaining) > 0 {
		input := strings.Join(remaining, " ")
		formatAndDisplay(os.Stdout, input, singleLine, clipboard.Copy)
		return
	}

	interactiveMode(os.Stdin, os.Stdout, singleLine, clipboard.Copy)
}

//nolint:errcheck // fmt writes to stdout; errors are not actionable in a CLI
func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Usage: claude-fix [OPTIONS] [COMMAND...]")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Fix formatting issues in terminal commands copied from Claude Code.")
	fmt.Fprintln(w, "If no arguments are given, starts interactive mode (paste, then press Enter twice).")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Options:")
	fmt.Fprintln(w, "  -s, --single-line  Collapse backslash line continuations into a single line")
	fmt.Fprintln(w, "  -h, --help         Show this help message")
	fmt.Fprintln(w, "      --version      Show version information")
}
