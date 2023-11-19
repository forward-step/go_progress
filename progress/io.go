package progress

import "errors"

// io.Writer
func (p *Progress) Write(b []byte) (n int, err error) {
	if p.isEnd() {
		err = errors.New("progress is end , cloud not write")
		return
	}

	n = len(b)
	p.Add(int64(n))
	return
}

// io.Reader
func (p *Progress) Read(b []byte) (n int, err error) {
	if p.isEnd() {
		err = errors.New("progress is end , cloud not read")
		return
	}

	n = len(b)
	p.Add(int64(n))
	return
}

// io.Closer
func (p *Progress) Close() (err error) {
	if p.isEnd() {
		err = errors.New("close progress is fail")
		return
	}
	if p.isDone() {
		p.Success()
	} else {
		p.Fail()
	}
	return
}
