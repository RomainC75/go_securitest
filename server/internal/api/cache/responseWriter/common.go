package responsewriter

import (
	"net/http"
)

type FakeResponseWriter struct {
	w *http.ResponseWriter
	b **[]byte
}

func NewResponseWriter(w *http.ResponseWriter) FakeResponseWriter {
	arr := &[]byte{}
	fw := FakeResponseWriter{
		w: w,
		b: &arr,
	}

	return fw
}

func (fw FakeResponseWriter) Write(b []byte) (int, error) {
	*fw.b = &b
	return 1, nil
}

func (fw FakeResponseWriter) Header() http.Header {
	return (*fw.w).Header()
}

func (fw FakeResponseWriter) WriteHeader(statusCode int) {
	(*fw.w).WriteHeader(statusCode)
}

func (fw FakeResponseWriter) Send() {
	(*fw.w).Write(**fw.b)
}

func (fw FakeResponseWriter) Get() []byte {
	return **fw.b
}
