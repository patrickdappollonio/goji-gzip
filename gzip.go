package gzip

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type compressResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *compressResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *compressResponseWriter) Write(b []byte) (int, error) {
	h := w.ResponseWriter.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", http.DetectContentType(b))
	}

	return w.Writer.Write(b)
}

func Compress(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	L:
		for _, enc := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
			switch strings.TrimSpace(enc) {
			case "gzip":
				w.Header().Set("Content-Encoding", "gzip")
				w.Header().Add("Vary", "Accept-Encoding")

				gw := gzip.NewWriter(w)
				defer gw.Close()

				w = &compressResponseWriter{
					Writer:         gw,
					ResponseWriter: w,
				}
				break L
			case "deflate":
				w.Header().Set("Content-Encoding", "deflate")
				w.Header().Add("Vary", "Accept-Encoding")

				fw, _ := flate.NewWriter(w, flate.DefaultCompression)
				defer fw.Close()

				w = &compressResponseWriter{
					Writer:         fw,
					ResponseWriter: w,
				}
				break L
			}
		}

		h.ServeHTTP(w, r)
	})
}
