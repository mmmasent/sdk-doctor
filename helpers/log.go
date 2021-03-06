package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

type Logger struct {
	warns  []string
	errors []string
}

func timeLogStr() string {
	t := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d.%03d",
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/int(time.Millisecond))
}

func (l *Logger) NewLine() {
	fmt.Printf("\n")
}

func (l *Logger) Log(format string, args ...interface{}) {
	line := fmt.Sprintf(format, args...)
	fmt.Printf("%s INFO ▶ %s\n", timeLogStr(), line)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	line := fmt.Sprintf(format, args...)
	fmt.Printf("%s WARN ▶ %s\n", timeLogStr(), line)
	l.warns = append(l.warns, line)
}

func (l *Logger) Error(format string, args ...interface{}) {
	line := fmt.Sprintf(format, args...)
	fmt.Printf("%s ERRO ▶ %s\n", timeLogStr(), line)
	l.errors = append(l.errors, line)
}

func (l Logger) PrintSummary() {
	fmt.Printf("Summary:\n")

	for _, line := range l.warns {
		fmt.Printf("%s %s\n", color.YellowString("[WARN]"), line)
	}
	for _, line := range l.errors {
		fmt.Printf("%s %s\n", color.RedString("[ERRO]"), line)
	}

	fmt.Printf("\n")
	if len(l.warns) > 0 || len(l.errors) > 0 {
		fmt.Printf("Found multiple issues, see listing above.\n")
	} else {
		fmt.Printf("Nothing of importance to note!  Nice job!\n")
	}
}
