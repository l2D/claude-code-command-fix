# Claude Code Command Fix

[![GitHub Issues](https://img.shields.io/badge/Related%20Issue-%23%204686-blue?logo=github)](https://github.com/anthropics/claude-code/issues/4686)
[![Community Solution](https://img.shields.io/badge/Community-Solution-green)](#)
[![Temporary Fix](https://img.shields.io/badge/Status-Temporary%20Fix-orange)](#when-to-uninstall)

> **TEMPORARY FIX** — This tool will be unnecessary once Anthropic fixes the copy-paste issue in Claude Code.

Community solution for [Claude Code Issue #4686](https://github.com/anthropics/claude-code/issues/4686) — helping 100+ developers resolve copy-paste formatting problems.

## Problem

Claude Code has a formatting bug where long terminal commands get line-wrapped when copied from code blocks, causing them to fail when pasted into Terminal.

**GitHub Issue:** [#4686 — Copy-paste from Claude Code introduces extra spaces/characters in code blocks](https://github.com/anthropics/claude-code/issues/4686)

### Example of the Problem

**What you get when copying from Claude Code:**

```bash
sudo launchctl unload
  /System/Library/LaunchDaemons/com.apple.backupd-auto.plist && sudo
  launchctl load /System/Library/LaunchDaemons/com.apple.backupd-auto.plist
```

**What you need (single line):**

```bash
sudo launchctl unload /System/Library/LaunchDaemons/com.apple.backupd-auto.plist && sudo launchctl load /System/Library/LaunchDaemons/com.apple.backupd-auto.plist
```

## Solution

This tool fixes the formatting issues and copies the corrected command to your clipboard.

## Installation

### Homebrew (Recommended)

```bash
brew install l2D/tap/claude-code-command-fix
```

This installs both `claude-fix` and `claude-command-fix` commands.

### Go Install

```bash
go install github.com/l2D/claude-code-command-fix/cmd/claude-fix@latest
go install github.com/l2D/claude-code-command-fix/cmd/claude-command-fix@latest
```

### Download Binary

Grab a pre-built binary from the [Releases](https://github.com/l2D/claude-code-command-fix/releases) page for your OS and architecture.

### Build from Source

```bash
git clone https://github.com/l2D/claude-code-command-fix.git
cd claude-code-command-fix
make build
```

Binaries are output to `bin/claude-fix` and `bin/claude-command-fix`.

## Usage

```
Usage: claude-fix [OPTIONS] [COMMAND...]

Options:
  -s, --single-line  Collapse backslash line continuations into a single line
  -h, --help         Show this help message
      --version      Show version information
```

### Command Line

```bash
claude-fix "your broken command here"
```

Also available as `claude-command-fix`, `cc-fix`, or `ccfix`.

### Interactive Mode

```bash
claude-fix
```

Then paste your broken command and press Enter twice.

### Multi-line Commands

Commands with `\` line continuations are auto-detected and preserved with consistent indentation:

```bash
claude-fix "docker run \
      --name myapp \
  -p 8080:80 \
          -v /data:/data \
    nginx:latest"
```

**Output:**

```
════════════════════════════════════════════════════════════
docker run \
  --name myapp \
  -p 8080:80 \
  -v /data:/data \
  nginx:latest
════════════════════════════════════════════════════════════
Copied to clipboard!
```

Use `-s` to collapse into a single line instead:

```bash
claude-fix -s "docker run \
  --name myapp \
  -p 8080:80 \
  nginx:latest"
```

**Output:**

```
════════════════════════════════════════════════════════════
docker run --name myapp -p 8080:80 nginx:latest
════════════════════════════════════════════════════════════
Copied to clipboard!
```

### Single-line Commands

Commands without `\` continuations are always collapsed to a single line:

```bash
claude-fix "sudo launchctl unload
  /System/Library/LaunchDaemons/com.apple.backupd-auto.plist && sudo
  launchctl load /System/Library/LaunchDaemons/com.apple.backupd-auto.plist"
```

**Output:**

```
════════════════════════════════════════════════════════════
sudo launchctl unload /System/Library/LaunchDaemons/com.apple.backupd-auto.plist && sudo launchctl load /System/Library/LaunchDaemons/com.apple.backupd-auto.plist
════════════════════════════════════════════════════════════
Copied to clipboard!
```

## What It Fixes

- Removes unwanted line breaks
- Fixes extra whitespace and indentation
- Preserves proper operator spacing (`&&`, `|`, `;`)
- Auto-detects and preserves `\` line continuations with consistent indentation
- Copies corrected command to clipboard automatically
- Works on macOS (`pbcopy`), Linux (`xclip` / `wl-copy`), and Windows (`clip.exe`)

## When to Uninstall

**Remove this tool when:**

- Anthropic releases a fix for the copy-paste issue
- Claude Code adds a "Copy" button to code blocks
- The formatting issue is resolved in a future update

```bash
# Homebrew
brew uninstall claude-code-command-fix

# Go
rm "$(which claude-fix)" "$(which claude-command-fix)"
```

## Related Issues

- **Primary Issue:** [Claude Code #4686](https://github.com/anthropics/claude-code/issues/4686) — Copy-paste formatting problems
- **Impact:** Critical configurations fail silently, causing difficult-to-debug errors
- **Affects:** Terminal commands, connection strings, configuration files

## Contributing

This is a temporary community fix. Once Anthropic resolves the underlying issue, this repository will be archived.

If you want to help:

1. Test the tool with different command types
2. Report bugs or edge cases
3. Improve cross-platform clipboard support

## Credits

Inspired by [freyjay/claude-code-command-fix](https://github.com/freyjay/claude-code-command-fix) (Python). This is a Go rewrite distributed as a static binary.

## License

MIT License — See [LICENSE](LICENSE) file for details.

---

**Remember:** This is a temporary workaround. The real solution is for Anthropic to fix the copy-paste formatting in Claude Code itself.
