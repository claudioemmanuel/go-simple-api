
  
## Tecnologia
#### `Back-end`
- [Go](https://go.dev)
#### Para executar a aplicação será necessário o [Docker](https://www.docker.com)
#### Para os testes na API REST será necessário o [Postman](https://www.postman.com)

## Preparação do projeto

Clone o projeto
```bash
$ git clone https://github.com/claudioemmanuel/devbook-api.git
```
Iniciação do banco de dados
```bash
$ docker-compose up --build
```
Iniciar o projeto executando o comando
```bash
$ go run main.go
```
Se tudo correr bem, o projeto deverá estar rodando em **127.0.0.1:3000**

## Testes
Para execução dos testes unitários, após o projeto estar rodando, basta entrar com o comando  

```bash
$ go test ./pkg/tests
```
e você deverá ver um retorno como ```ok  command-line-arguments  0.221s``` indicando que todos os testes rodaram com sucesso.

## Postman
Para os testes há um arquivo com os payloads de exemplo no diretório **docs/payloads.json**
Basta importá-lo no **Postman** para utilizar.

## 📙 Licença
> Com base nos termos de [MIT LICENSE](https://opensource.org/licenses/MIT)

##### Feito por Claudio Emmanuel com ❤️