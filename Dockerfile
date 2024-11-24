FROM golang:1.22

# Configura o diretório de trabalho no contêiner
WORKDIR /go/src

# Adiciona o binário ao PATH (global)
ENV PATH="/go/bin:${PATH}"

# Inicializa um módulo Go e instala as ferramentas necessárias como root
RUN go mod init temp-module && \
  go install github.com/spf13/cobra-cli@latest && \
  go install github.com/golang/mock/mockgen@v1.5.0

# Instala o SQLite3
RUN apt-get update && apt-get install -y sqlite3

# Cria diretórios e ajusta permissões para o usuário www-data
RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache /go/src /go/bin /go/pkg
RUN chown -R www-data:www-data /go /var/www/.cache

# Alterna para o usuário www-data
USER www-data

# Define o comando padrão do contêiner
CMD ["tail", "-f", "/dev/null"]
