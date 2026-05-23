package my_module

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

// init automatically registers the module and the Caddyfile directive
func init() {
	caddy.RegisterModule(MyGeneratedModuleMandatoryCode{})
	
	// Registra a diretiva que será lida no Caddyfile (ex: my_module)
	httpcaddyfile.RegisterHandlerDirective("my_module", parseCaddyfile)
}

// MyGeneratedModuleMandatoryCode defines the structure of your plugin.
type MyGeneratedModuleMandatoryCode struct {
	// Propriedades geradas a partir do arquivo .um
	ApiKeysList string `json:"api_keys_list,omitempty"`
}

// CaddyModule returns the module metadata.
func (MyGeneratedModuleMandatoryCode) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.generated_module_mandatory_code",
		New: func() caddy.Module { return new(MyGeneratedModuleMandatoryCode) },
	}
}

// Provision is optional. Use it to allocate resources or set DEFAULT values.
func (m *MyGeneratedModuleMandatoryCode) Provision(ctx caddy.Context) error {
	// O seu compilador injeta a lógica de defaults aqui.
	// Se o Caddyfile não passou "api_keys_list=...", usamos o "default is" do .um
	if m.ApiKeysList == "" {
		m.ApiKeysList = "~/api_keys.duckdb" 
	}
	
	// Aqui você também inicializaria a conexão com o DuckDB, por exemplo.
	return nil
}

// UnmarshalCaddyfile parses the Caddyfile arguments in "key=value" format.
func (m *MyGeneratedModuleMandatoryCode) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		// Pega todos os argumentos na mesma linha após o nome da diretiva
		args := d.RemainingArgs()
		
		for _, arg := range args {
			// Divide o argumento na primeira ocorrência do '='
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				return d.Errf("invalid argument format, expected key=value: %s", arg)
			}
			
			key := parts[0]
			val := parts[1]

			// O compilador gera este switch baseado nas propriedades do .um
			switch key {
			case "api_keys_list":
				m.ApiKeysList = val
			// case "outra_propriedade":
			// 	m.OutraProp = val
			default:
				return d.Errf("unknown property: %s", key)
			}
		}
	}
	return nil
}

// parseCaddyfile is the helper function that links the Caddyfile to the struct
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m MyGeneratedModuleMandatoryCode
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// ServeHTTP handles the incoming web requests.
func (m MyGeneratedModuleMandatoryCode) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Lógica do seu módulo
	// m.ApiKeysList estará preenchida corretamente (via Caddyfile ou default)
	
	fmt.Printf("Using API Keys from: %s\n", m.ApiKeysList)

	return next.ServeHTTP(w, r) 
}

// Interface guards ensure at compile time that the module implements all required Caddy interfaces.
var (
	_ caddy.Module                = (*MyGeneratedModuleMandatoryCode)(nil)
	_ caddy.Provisioner           = (*MyGeneratedModuleMandatoryCode)(nil)
	_ caddyhttp.MiddlewareHandler = (*MyGeneratedModuleMandatoryCode)(nil)
	_ caddyfile.Unmarshaler       = (*MyGeneratedModuleMandatoryCode)(nil)
)
