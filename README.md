# Stress test

## Configuração da apliciação

Esta aplicação utiliza docker para sua execução e para facilitar que o ambiente não precise ter GO instalado;
Para executar a aplicação é necessário ter Docker isntalado na máquina ou Go;
Para execução do projeto corretamente devem ser passadas as seguintes informações
--url
--requests
--concurrency

## Docker

Para executar a aplicação utilizando docker basta aplicar os seguintes comando:
docker build -t desafio-stress-test -f ./Dockerfile .
docker run desafio-stress-test --url=http://google.com --requests=100 --concurrency=10

## Go

Para executar a aplicação utilizando Go basta aplicar o seguinte comando:
go run cmd/main.go --url=https://google.com --requests=100 --concurrency=10
