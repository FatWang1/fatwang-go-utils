package logger

type noop struct{}

func NewNoop() *noop {
	return &noop{}
}
