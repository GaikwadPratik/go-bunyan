package bunyan

type ParsedArgs struct {
	Args       []string
	Level      LogLevel
	Color      bool
	Pagination bool
	JsonIndent int
	Strict     bool
	Pids       []int
	PidsType   bool
	Timeformat TimeDisplayFormat
	Version    bool
	OutputMode OutputMode
}
