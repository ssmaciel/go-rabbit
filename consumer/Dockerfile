# Definindo imagem de base
FROM golang:1.15

# Definindo diretório de trabalho
WORKDIR /go/src/go-rabbit

# Copiando diretório para a imagem
# Como é o WORKDIR poderia também ser: ADD . .
ADD . .

# Executando comando de build
RUN GOOS=linux go build

CMD "./consumer" 
