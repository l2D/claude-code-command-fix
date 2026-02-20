package formatter

import "testing"

func TestFormatCommand(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "basic formatting with pipe",
			in:   "ls -la\n  | grep test",
			want: "ls -la | grep test",
		},
		{
			name: "complex command with &&",
			in:   "sudo launchctl unload\n  /System/Library/LaunchDaemons/com.apple.backupd-auto.plist\n  && sudo launchctl load\n  /System/Library/LaunchDaemons/com.apple.backupd-auto.plist",
			want: "sudo launchctl unload /System/Library/LaunchDaemons/com.apple.backupd-auto.plist && sudo launchctl load /System/Library/LaunchDaemons/com.apple.backupd-auto.plist",
		},
		{
			name: "multiple operators",
			in:   "command1   &&    command2 |  command3;   command4",
			want: "command1 && command2 | command3; command4",
		},
		{
			name: "empty string",
			in:   "",
			want: "",
		},
		{
			name: "whitespace only",
			in:   "   \n  \n  ",
			want: "",
		},
		{
			name: "single line unchanged",
			in:   "ls -la /home/user",
			want: "ls -la /home/user",
		},
		{
			name: "excessive whitespace",
			in:   "command1     arg1    arg2        arg3",
			want: "command1 arg1 arg2 arg3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatCommand(tt.in)
			if got != tt.want {
				t.Errorf("FormatCommand(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
