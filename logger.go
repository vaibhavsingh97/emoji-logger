// Copyright © 2017 Vaibhav Singh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// This is the set of emoji definition.
const (
	CriticalEmoji = "🚑"
	DebugEmoji    = "🐞"
	InfoEmoji     = "🧐"
	SuccessEmoji  = "✅"
	ErrorEmoji    = "😱"
	WarningEmoji  = "⚠️"
)

// It contains configuration of logger.
var (
	// Level	Numeric value
	// SUCCESS	60
	// CRITICAL	50
	// ERROR	40
	// WARNING	30
	// INFO		20
	// DEBUG	10
	// NOTSET	0
	//
	//
	// Sets the threshold for the logger.
	// Logging messages which are less severe than level will be ignored.
	// The  default level is set to NOTSET (which causes all messages to be
	// processed).
	Level = 0
	// TimeStamps defines if the output has timestamp or not.
	// This is a global option and affects all methods.
	TimeStamps = true
	// Color defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a
	// terminal or not. This is a global option and affects all methods.
	Color = false
	// BackgroundColor defines if the output has Background color or not.
	// It's dynamically set to false or true based on the stdout's file
	// descriptor referring to a terminal or not. This is a global option and
	// affects all methods.
	BackgroundColor = false
)

// Critical formats according to a format specifier and writes to standard output.
func Critical(format string, a ...interface{}) {
	if Level <= 50 {
		a, w := extractLoggerArguments(format, a...)
		s := fmt.Sprintf(emojiLabel(format, CriticalEmoji), a...)
		if Color && BackgroundColor {
			Color = false
		}
		if Color {
			s = color.New(color.FgHiMagenta).Sprintf(s)
		}
		if BackgroundColor {
			c := color.New(color.FgHiWhite)
			s = c.Add(color.BgHiMagenta).Sprintf(s)
		}
		fmt.Fprint(w, s)
	}
}

// Debug formats according to a format specifier and writes to standard output.
func Debug(format string, a ...interface{}) {
	if Level <= 10 {
		a, w := extractLoggerArguments(format, a...)
		s := fmt.Sprintf(emojiLabel(format, DebugEmoji), a...)
		if Color && BackgroundColor {
			Color = false
		}
		if Color {
			s = color.New(color.FgHiWhite).Sprintf(s)
		}
		if BackgroundColor {
			c := color.New(color.FgHiWhite)
			s = c.Add(color.BgBlack).Sprintf(s)
		}
		fmt.Fprint(w, s)
	}
}

// Info formats according to a format specifier and writes to standard output.
func Info(format string, a ...interface{}) {
	if Level <= 20 {
		a, w := extractLoggerArguments(format, a...)
		s := fmt.Sprintf(emojiLabel(format, InfoEmoji), a...)
		if Color && BackgroundColor {
			Color = false
		}
		if Color {
			s = color.New(color.FgHiWhite).Sprintf(s)
		}
		if BackgroundColor {
			c := color.New(color.FgHiWhite)
			s = c.Add(color.BgBlack).Sprintf(s)
		}
		fmt.Fprint(w, s)
	}
}

// Success formats according to a format specifier and writes to standard output.
func Success(format string, a ...interface{}) {
	if Level <= 60 {
		a, w := extractLoggerArguments(format, a...)
		s := fmt.Sprintf(emojiLabel(format, SuccessEmoji), a...)
		if Color && BackgroundColor {
			Color = false
		}
		if Color {
			s = color.New(color.FgHiGreen).Sprintf(s)
		}
		if BackgroundColor {
			c := color.New(color.FgHiWhite)
			s = c.Add(color.BgHiGreen).Sprintf(s)
		}
		fmt.Fprint(w, s)
	}
}

// Warning formats according to a format specifier and writes to standard output.
func Warning(format string, a ...interface{}) {
	if Level <= 30 {
		a, w := extractLoggerArguments(format, a...)
		s := fmt.Sprintf(emojiLabel(format, WarningEmoji), a...)
		if Color && BackgroundColor {
			Color = false
		}
		if Color {
			s = color.New(color.FgHiYellow).Sprintf(s)
		}
		if BackgroundColor {
			c := color.New(color.FgHiBlack)
			s = c.Add(color.BgYellow).Sprintf(s)
		}
		fmt.Fprint(w, s)
	}
}

// Error formats according to a format specifier and writes to standard output.
func Error(format string, a ...interface{}) {
	if Level <= 40 {
		a, w := extractLoggerArguments(format, a...)
		s := fmt.Sprintf(emojiLabel(format, ErrorEmoji), a...)
		if Color && BackgroundColor {
			Color = false
		}
		if Color {
			s = color.New(color.Bold, color.FgHiRed).Sprintf(s)
		}
		if BackgroundColor {
			c := color.New(color.FgHiWhite)
			s = c.Add(color.BgHiRed).Sprintf(s)
		}
		fmt.Fprint(w, s)
	}
}

func extractLoggerArguments(format string, a ...interface{}) ([]interface{}, io.Writer) {
	n := strings.Count(format, "%")
	if n > len(a) {
		panic("string is missing")
	}
	if n < len(a) {
		panic(`%!! is missing`)
	}
	var w io.Writer = os.Stdout
	if length := len(a); length > 0 {
		// extract an io.Writer at the end of a
		if value, ok := a[length-1].(io.Writer); ok {
			w = value
			a = a[0 : length-1]
		}
	}
	return a, w
}

func emojiLabel(format, label string) string {
	if TimeStamps {
		return labelWithTime(format, label)
	} else {
		return labelWithoutTime(format, label)
	}
}

func labelWithTime(format, label string) string {
	t := time.Now()
	rfcTime := t.Format(time.RFC3339)
	if !strings.Contains(format, "\n") {
		format = fmt.Sprintf("%s%s", format, "\n")
	}
	return fmt.Sprintf("%s  %s  %s", rfcTime, label, format)
}

func labelWithoutTime(format, label string) string {
	if !strings.Contains(format, "\n") {
		format = fmt.Sprintf("%s%s", format, "\n")
	}
	return fmt.Sprintf("%s  %s", label, format)
}
