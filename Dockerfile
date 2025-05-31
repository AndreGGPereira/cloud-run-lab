FROM golang:1.23

WORKDIR /app

# Copie os arquivos go.mod e go.sum
COPY go.mod ./
COPY go.sum ./

# Baixe todas as dependências do módulo Go
RUN go mod tidy
RUN go mod download

# Copie o código-fonte restante
COPY . .

# Construa o executável
RUN go build -o main 

EXPOSE 8080
# Comando para executar o binário
CMD ["./main"]