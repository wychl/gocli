package middleware

import (
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/wychl/gocli/command"
)

// Logger 日志中间件
type Logger struct {
	command.Command
	logger log.Logger
}

func NewLogger(cmd command.Command) *Logger {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "name", cmd.Name())
	return &Logger{cmd, logger}
}

func (l Logger) Run(cfg interface{}) error {
	l.logger.Log("msg", "starting")
	start := time.Now()
	l.logger.Log("msg", "running")
	if err := l.Run(cfg); err != nil {
		l.logger.Log("err", err)
	}
	l.logger.Log("msg", "end", "duration", time.Since(start))
	return nil
}
