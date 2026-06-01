package my_module

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"sync"

	"://github.com"
	"://github.com/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(ApiGuardModule{})
}

type ApiGuardModule struct {
	// Memória ram de busca rápida O(1)
	apiKeysMemory map[string]struct{}
	mu            sync.RWMutex
}

func (ApiGuardModule) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.api_guard",
		New: func() caddy.Module { return new(ApiGuardModule) },
	}
}

// Provision executa 1 vez no início do Caddy. É o formato mais rápido para buscas futuras.
func (m *ApiGuardModule) Provision(ctx caddy.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.apiKeysMemory = make(map[string]struct{})

	// Abre o arquivo com 1 chave por linha
	file, err := os.Open("api_keys.txt")
	if err != nil {
		// Se o arquivo não existir, criamos o mapa vazio (ou pode retornar o erro)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key := strings.TrimSpace(scanner.Text())
		if key != "" {
			m.apiKeysMemory[key] = struct{}{} // Alocação de memória zero para o valor
		}
	}

	return scanner.Err()
}

func (m *ApiGuardModule) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Captura o Header solicitado
	apiKeyInput := r.Header.Get("API_KEY")

	m.mu.RLock()
	_, exists := m.apiKeysMemory[apiKeyInput]
	m.mu.RUnlock()

	// Se a chave não existir ou o header vier vazio, barra imediatamente com 401 Unauthorized
	if apiKeyInput == "" || !exists {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized: Invalid or missing API_KEY"))
		return nil // Corta o fluxo aqui, não chama o 'next'
	}

	// Se passar, continua o fluxo para o próximo handler do Caddy
	return next.ServeHTTP(w, r)
}

// Interface guards
var (
	_ caddy.Module                = (*ApiGuardModule)(nil)
	_ caddy.Provisioner           = (*ApiGuardModule)(nil)
	_ caddyhttp.MiddlewareHandler = (*ApiGuardModule)(nil)
)
