package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"comunty/ms-auth/conf"
	"comunty/ms-auth/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Iniciando microservicio auth")
	level := []log.Level{log.InfoLevel, log.DebugLevel, log.ErrorLevel, log.TraceLevel}
	log.SetLevel(level[conf.GlobalConf.LogLevel])
	log.Info("logLevel: " + log.GetLevel().String())
	engine := gin.Default()

	groups := engine.Group("/ms-auth")
	routes.Auth(groups)

	log.Info("ms-auth iniciado")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", limit(Gzip(engine))))
}

// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !conf.GlobalConf.GzipResponse || !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler.ServeHTTP(gzw, r)
	})
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json, text/plain")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Content-Type, Apikey, "+
				"Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods,  Origin, x-ibm-client-id, X-Authorization-Type, X-System-Type")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			return
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Strict-Transport-Security", "max-age=86400; includeSubDomains")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			next.ServeHTTP(w, r)
		}
	})
}
