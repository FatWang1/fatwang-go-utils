package logger

type noop struct{}

func NewNoop() *noop {
	return &noop{}
}

func (n *noop) Info(_ string) {
	return
}
