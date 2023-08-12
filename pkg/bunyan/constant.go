package bunyan

type LogLevel string

const LogLevelFatal LogLevel = "fatal"
const LogLevelError LogLevel = "error"
const LogLevelWarn LogLevel = "warn"
const LogLevelInfo LogLevel = "info"
const LogLevelDebug LogLevel = "debug"
const LogLevelTrace LogLevel = "trace"

type TimeDisplayFormat string

const TimeDisplayFormatUtc TimeDisplayFormat = "utc"
const TimeDisplayFormatLocal TimeDisplayFormat = "local"

type OutputMode string

const OutputModeLong OutputMode = "long"
const OutputModeJson OutputMode = "json"
const OutputModeJsonN OutputMode = "json-N"
const OutputModeBunyan OutputMode = "bunyan"
const OutputModeShort OutputMode = "short"
const OutputModeSimple OutputMode = "simple"
