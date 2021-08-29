package server

import "net/http"

type Config struct {
	Addr string
	Port string
}

// Run will run our http server
func Serve(c Config) error {

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", Healthz)

	return http.ListenAndServe(c.Addr+":"+c.Port, mux)
}
