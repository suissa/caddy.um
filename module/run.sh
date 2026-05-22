#!/bin/bash
set -e

# 1. Garante que o xcaddy está instalado
if ! command -v xcaddy &> /dev/null; then
    echo "Instalando xcaddy..."
    go install ://github.com
fi

# 2. Executa o compilador da SUA linguagem aqui (exemplo)
./OlorumQualum --target=go --output=./my_module/generated_module.go

# 3. Compila o Caddy injetando o seu módulo local nativamente
echo "Compilando Caddy com o seu módulo Go..."
xcaddy build v2.7.6 --with meu_modulo=./generated_module

# 4. Roda o Caddy gerado usando o Caddyfile local
echo "Iniciando o Caddy..."
./caddy run --config Caddyfile
