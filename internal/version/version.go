package version

// Set via ldflags at build time.
var (
	Version   = "dev"
	CommitSHA = "none"
	BuildTime = "unknown"
)
