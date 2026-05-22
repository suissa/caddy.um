package my_module

import (
	"net/http"

	"://github.com"
	"://github.com/modules/caddyhttp"
)

// init automatically registers the module into Caddy when the server starts.
func init() {
	caddy.RegisterModule(MyGeneratedModule{})
}

// MyGeneratedModule defines the structure of your plugin.
// Your compiler can add fields here if your language needs state or configuration.
type MyGeneratedModule struct{}

// CaddyModule returns the module metadata for Caddy's ecosystem.
// The ID must follow the pattern: http.handlers.your_module_name
func (MyGeneratedModule) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.generated_module",
		New: func() caddy.Module { return new(MyGeneratedModule) },
	}
}

// Provision is optional. Use it to allocate resources or validate initial configurations.
func (m *MyGeneratedModule) Provision(ctx caddy.Context) error {
	return nil
}

// ServeHTTP handles the incoming web requests.
// Any web request hitting the configured route will execute the logic inside this method.
func (m MyGeneratedModule) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// SET THE RESPONSE HEADER
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// YOUR GENERATED BUSINESS LOGIC GOES HERE
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Native Go compiled directly into Caddy!"))

	// Return nil if you want to terminate the response loop here.
	// If this was an authentication middleware, you would call `return next.ServeHTTP(w, r)`.
	return nil 
}

// Interface guards ensure at compile time that the module implements all required Caddy interfaces.
var (
	_ caddy.Module                = (*MyGeneratedModule)(nil)
	_ caddy.Provisioner           = (*MyGeneratedModule)(nil)
	_ caddyhttp.MiddlewareHandler = (*MyGeneratedModule)(nil)
)
