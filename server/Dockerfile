# Utiliza uma imagem base de Go
FROM golang:1.22.3

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

RUN apt-get update && apt-get install -y

#intall git
RUN apt-get install git -y

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download



# Instala o CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest



# Copia todos os arquivos do projeto para o contêiner
COPY . .

COPY docker/entrypoint.sh .
RUN chmod +x entrypoint.sh



ENTRYPOINT [ "entrypoint.sh" ]
