# Utiliza uma imagem base de Go
FROM golang:1.22.3

# Instala o CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia todos os arquivos do projeto para o contêiner
COPY . .


ENTRYPOINT [ "./entrypoint.sh" ]