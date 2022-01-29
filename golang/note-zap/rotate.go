package logx

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultMaxSize    = 1024 // 1024MB
	DefaultMaxAge     = 7
	DefaultMaxBackups = 10
)

// RotateOptions for logger rotate
type RotateOptions struct {
	MaxSize   int
	MaxAge    int
	MaxBackup int
	Compress  bool
	LocalTime bool
}

func newDefaultRotateOptions() RotateOptions {
	return RotateOptions{
		MaxSize:   DefaultMaxSize,
		MaxAge:    DefaultMaxAge,
		MaxBackup: DefaultMaxBackups,
		Compress:  true,  // default compress log
		LocalTime: false, // default use UTC time
	}
}

type RotateOption func(*RotateOptions)

func newRotateOptions(opts ...RotateOption) RotateOptions {
	opt := newDefaultRotateOptions()

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func newLumberLogWriter(filename string, opts ...RotateOption) *lumberjack.Logger {
	option := newRotateOptions(opts...)

	hook := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    option.MaxSize,
		MaxBackups: option.MaxBackup,
		MaxAge:     option.MaxAge,
		LocalTime:  option.LocalTime,
		Compress:   option.Compress,
	}

	return hook
}

func MaxSize(size int) RotateOption {
	return func(o *RotateOptions) {
		o.MaxSize = size
	}
}

func MaxAge(age int) RotateOption {
	return func(o *RotateOptions) {
		o.MaxAge = age
	}
}

func MaxBackup(backup int) RotateOption {
	return func(o *RotateOptions) {
		o.MaxBackup = backup
	}
}

func Compress(compress bool) RotateOption {
	return func(o *RotateOptions) {
		o.Compress = compress
	}
}

func LocalTime(localtime bool) RotateOption {
	return func(o *RotateOptions) {
		o.LocalTime = localtime
	}
}
