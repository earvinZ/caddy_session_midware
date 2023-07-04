package simpleheaderadder

import (
	"github.com/caddyserver/caddy/v2"

	"net/http"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"

	"fmt"
)

type MyFirstMid struct {
}

func init() {
	caddy.RegisterModule(MyFirstMid{})
}

func (MyFirstMid) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "jude_add_header",
		New: func() caddy.Module { return new(MyFirstMid) },
	}
}

// Implement the HTTP handler interface
func (m MyFirstMid) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// your middleware logic here

	fmt.Fprintf(w, "Hello from my middleware!")
	r.Header.Add("hey_jude", "dont_be_afraid")

	return next.ServeHTTP(w, r)
}
