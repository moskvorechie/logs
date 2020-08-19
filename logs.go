package logs

import (
	"github.com/kei2100/rotate"
	"github.com/rs/zerolog"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"time"
)

type Log struct {
	logger zerolog.Logger
	w      *rotate.Writer
}

type Config struct {
	App      string
	FilePath string
	Clear    bool
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func New(cfg *Config) (log Log, err error) {

	// Init
	log.logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Write to file and console if path exist
	if len(cfg.FilePath) > 0 {

		// Truncate file
		if cfg.Clear {
			_ = os.Remove(cfg.FilePath)
			files, err := filepath.Glob(cfg.FilePath + "*")
			if err != nil {
				panic(err)
			}
			for _, f := range files {
				_ = os.Remove(f)
			}
		}

		// Paths
		logPath := path.Dir(cfg.FilePath)
		fileName := path.Base(cfg.FilePath)

		const bytes3 int64 = 5 * 1024 * 1024
		log.w, err = rotate.NewWriter(logPath, fileName, rotate.WithSizeBasedPolicy(bytes3))
		if err != nil {
			panic(err)
		}

		// Log to file and console
		log.logger = zerolog.New(log.w).With().Timestamp().Logger()
	}

	// Add datetime hook
	dtHook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, msg string) {
		e.Str("datetime", time.Now().Format("02.01.2006 15:04:05.999999999"))
	})
	log.logger = log.logger.Hook(dtHook)

	// Add app if exist
	if len(cfg.App) > 0 {
		log.logger = log.logger.With().Str("app", cfg.App).Logger()
	}

	return
}

func (l *Log) SetLevel(level zerolog.Level) {
	l.logger = l.logger.Level(level)
}

func (l *Log) Close() {
	if l.w != nil {
		_ = l.w.Close()
	}
	l.w = nil
	l.logger = zerolog.Logger{}
}

func (l *Log) SetCustomLogger(logger zerolog.Logger) {
	l.logger = logger
}

func (l *Log) Logger() *zerolog.Logger {
	return &l.logger
}

func (l *Log) Debug(text string) {
	l.logger.Debug().Caller(1).Msg(text)
}

func (l *Log) DebugF(format string, v ...interface{}) {
	l.logger.Debug().Msgf(format, v...)
}

func (l *Log) Info(text string) {
	l.logger.Info().Msg(text)
}

func (l *Log) InfoF(format string, v ...interface{}) {
	l.logger.Info().Msgf(format, v...)
}

func (l *Log) Warn(text string) {
	l.logger.Warn().Msg(text)
}

func (l *Log) WarnF(format string, v ...interface{}) {
	l.logger.Warn().Msgf(format, v...)
}

func (l *Log) Error(text string) {
	l.logger.Error().Msg(text)
}

func (l *Log) ErrorF(format string, v ...interface{}) {
	l.logger.Error().Msgf(format, v...)
}

func (l *Log) Fatal(text string) {
	l.logger.Fatal().Msg(text)
}

func (l *Log) FatalF(format string, v ...interface{}) {
	l.logger.Fatal().Msgf(format, v...)
}

func (l *Log) LogError(err error) {
	l.logger.Error().Msgf("Error stack: %s", string(debug.Stack()))
	l.logger.Error().Msgf("Error: %+v", err)
}

func (l *Log) FatalError(err error) {
	l.logger.Error().Msgf("Fatal stack: \n" + string(debug.Stack()))
	l.logger.Fatal().Msgf("Error: %+v", err)
}
