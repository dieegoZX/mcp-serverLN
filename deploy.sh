#!/bin/bash

echo "🔄 Iniciando o deploy do MCP Server para o GitHub..."

# Adiciona todos os arquivos modificados
git add .

# Pede uma mensagem de commit
read -p "✏️  Escreva a mensagem do commit: " commit_message

# Faz o commit
git commit -m "$commit_message"

# Sobe para o GitHub
git push origin main

echo "✅ Deploy concluído com sucesso! MCP atualizado no GitHub!"
