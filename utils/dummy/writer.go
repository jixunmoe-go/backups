package dummy

type Writer struct{}

func (w *Writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}
