package log

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// LogFormatter for custom logrus format
type LogFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

// Format logrus
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))

	return []byte(fmt.Sprintf(
		"%s %s %s source=%s\n",
		f.LevelDesc[entry.Level], timestamp, entry.Message, entry.Data["source"])), nil
}
