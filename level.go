package errors

import "fmt"

type Level int

// Names for common levels.
const (
	LevelInternal Level = iota
	LevelBad
	LevelAuth
)

func (l Level) String() string {
	str := func(base string) string {
		return fmt.Sprintf("%s%+d", base, l)
	}

	switch l {
	case LevelBad:
		return str("BAD")
	case LevelAuth:
		return str("AUTH")
	default:
		return str("INTERNAL")
	}
}
