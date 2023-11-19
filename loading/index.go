package loading

type Loading struct {
	index   int
	current int
}

func New(i int) *Loading {
	if i < 0 || i >= len(spinners) {
		panic("loading.New(i int): out of spinners")
	}
	return &Loading{
		index: i,
	}
}

func (l *Loading) Next() string {
	spinner := spinners[l.index]
	length := len(spinner)
	now := l.current % length
	l.current++
	return spinner[now]
}
