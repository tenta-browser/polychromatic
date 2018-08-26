/**
 * Polychromatic
 *
 *    Copyright 2018 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * eventual.go: Wrapper for logs which may be interesting later, but aren't interesting now
 */

package polychromatic

import (
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type eventualLogEntry struct {
	date      time.Time
	level     logrus.Level
	hasFormat bool
	format    string
	data      []interface{}
	str       string
}

// Type EventualLogger provides a buffer for log entries whose usefulnes cannot be
// determined at creation time (but only eventually)
type EventualLogger struct {
	entries []*eventualLogEntry
	w       *sync.Mutex
}

// Creates a new eventual logger
func NewEventualLogger() *EventualLogger {
	return &EventualLogger{
		entries: make([]*eventualLogEntry, 0),
		w:       &sync.Mutex{},
	}
}

// Queuef buffers the selected format string and args to be written later. It does not
// perform any kind of immediate evaluation on the arguments, so points which may be
// different later should probably be converted to a string when calling
func (l *EventualLogger) Queuef(level logrus.Level, format string, args ...interface{}) {
	e := &eventualLogEntry{
		date:      time.Now(),
		level:     level,
		hasFormat: true,
		format:    format,
		data:      args,
		str:       "",
	}
	defer l.w.Unlock()
	l.w.Lock()
	l.entries = append(l.entries, e)
}

// Queue buffers teh selected message to be written later
func (l *EventualLogger) Queue(level logrus.Level, message string) {
	e := &eventualLogEntry{
		date:      time.Now(),
		level:     level,
		hasFormat: false,
		format:    "",
		data:      nil,
		str:       message,
	}
	defer l.w.Unlock()
	l.w.Lock()
	l.entries = append(l.entries, e)
}

// Flush writes out everything from the buffer
func (l *EventualLogger) Flush(target *logrus.Entry) {
	defer l.w.Unlock()
	l.w.Lock()
	for _, e := range l.entries {
		t := target.WithTime(e.date)
		if e.hasFormat {
			switch e.level {
			case logrus.DebugLevel:
				t.Debugf(e.format, e.data...)
			case logrus.InfoLevel:
				t.Infof(e.format, e.data...)
			case logrus.WarnLevel:
				t.Warnf(e.format, e.data...)
			case logrus.ErrorLevel:
				t.Errorf(e.format, e.data...)
			default:
				t.WithField("original-level", e.level.String()).Printf(e.format, e.data...)
			}
		} else {
			switch e.level {
			case logrus.DebugLevel:
				t.Debug(e.str)
			case logrus.InfoLevel:
				t.Info(e.str)
			case logrus.WarnLevel:
				t.Warn(e.str)
			case logrus.ErrorLevel:
				t.Error(e.str)
			default:
				t.WithField("original-level", e.level.String()).Print(e.str)
			}
		}
	}
	l.entries = make([]*eventualLogEntry, 0)
}
