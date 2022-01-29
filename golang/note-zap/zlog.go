package logx

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Strings     = zap.Strings
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any
	ZErr        = zap.Error
)

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel  // -1
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3
	PanicLevel  Level = zap.PanicLevel  // 4. // FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel  Level = zap.FatalLevel  // 5
)

type Field = zap.Field

func (l *Logger) With(fields ...Field) *Logger {
	logger := &Logger{
		zapLogger: l.zapLogger.With(fields...),
	}
	logger.sugarLogger = logger.zapLogger.Sugar()
	return logger
}

func (l *Logger) Named(name string) *Logger {
	logger := &Logger{
		zapLogger: l.zapLogger.Named(name),
	}
	logger.sugarLogger = logger.zapLogger.Sugar()
	return logger
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.zapLogger.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

type Logger struct {
	zapLogger   *zap.Logger
	sugarLogger *zap.SugaredLogger
}

type Option = zap.Option

var (
	WithCaller    = zap.WithCaller
	AddStacktrace = zap.AddStacktrace
	AddCallerSkip = zap.AddCallerSkip
)

func newZapLogger(writer io.Writer, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}

	cfg := zap.NewProductionConfig()

	// cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	// }

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)

	logger := &Logger{
		zapLogger: zap.New(core, opts...),
	}
	logger.sugarLogger = logger.zapLogger.Sugar()

	return logger
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}

func getGeneralLogWriter(filename string) zapcore.WriteSyncer {
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return f
}

func NewLoggerWithOption(filename string, level Level, zapOptions ...Option) *Logger {
	lumberjackWriter := newLumberLogWriter(filename)
	return newZapLogger(lumberjackWriter, level, zapOptions...)
}

// NewLogger creates a new Logger with best practice logger
func NewLogger(filename string, level Level) *Logger {
	return NewLoggerWithOption(filename, level,
		AddStacktrace(DPanicLevel),
		AddCallerSkip(1),
		WithCaller(true))
}
