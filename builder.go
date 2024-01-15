package errors

import "fmt"

type builder struct {
	level Level
}

func (b *builder) New(msg string) error {
	return &fundamental{
		level: b.level,
		msg:   msg,
		stack: callers(),
	}
}

func (b *builder) Errorf(format string, args ...any) error {
	return &fundamental{
		level: b.level,
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

func (b *builder) Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   msg,
	}
	return &withStack{
		level: b.level,
		error: err,
		stack: callers(),
	}
}

func (b *builder) Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
	return &withStack{
		level: b.level,
		error: err,
		stack: callers(),
	}
}

type Fundamental interface {
	New(msg string) error
	Errorf(format string, args ...any) error
	Wrap(err error, msg string) error
	Wrapf(err error, format string, args ...any) error
}

func Internal() Fundamental { return &builder{LevelInternal} }

func Bad() Fundamental { return &builder{LevelBad} }

func Auth() Fundamental { return &builder{LevelAuth} }
