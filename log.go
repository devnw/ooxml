package office

import (
	"log"
)

// Log is used to log content from within the library.  The intent is to use
// logging sparingly, preferring to return an error.  At the very least this
// allows redirecting logs to somewhere more appropriate than stdout.
var Log = log.Printf

// DisableLogging sets the Log function to a no-op so that any log messages are
// silently discarded.
func DisableLogging() {
	Log = func(string, ...interface{}) {}
}
