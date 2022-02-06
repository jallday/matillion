package responsewriter

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	status int
	body   string
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)

}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.body = string(b)
	if w.status == 0 {
		w.status = 200
	}
	return n, err
}

// StatusCode - fetches the status code
func (w *ResponseWriter) StatusCode() int {
	return w.status
}

// Body - fetches the body
func (w *ResponseWriter) Body() string {
	return w.body
}
