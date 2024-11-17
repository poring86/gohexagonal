FROM golang:1.21

# Configura o diretório de trabalho no contêiner
WORKDIR /go/src

# Adiciona o binário ao PATH
ENV PATH="/go/bin:${PATH}"

# Inicializa um módulo Go e instala as ferramentas necessárias
RUN go mod init temp-module && \
  go install github.com/spf13/cobra-cli@latest && \
  go install github.com/golang/mock/mockgen@v1.5.0

# Instala o SQLite3
RUN apt-get update && apt-get install -y sqlite3

# Ajusta permissões para o usuário www-data
RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache

# Muda para o usuário www-data
USER www-data

# Define um comando padrão para o contêiner
CMD ["tail", "-f", "/dev/null"]
