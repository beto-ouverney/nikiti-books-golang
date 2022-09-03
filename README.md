# Nikiti Books GO LANG with SOLID, MONGODB e MVC #

É uma API que controla os livros de um usuário

### O Desafio

Uma API que controla os livros de um usuário, lista um ou mais livros, exclui, adiciona e atualiza.

**O usuário será capaz de**

  - endpoint GET /books
    * Lista todos os livros do usuário
    * Se o usuário não tiver livros, deve retornar o Status 200 com o body vazio
   
  Exemplo de retorno:
   ```json
  [
	{
		"title": "The Lord of the Rings",
		"author": "J. R. R. Tolkien",
		"category": [
			"Fantasy",
			"Adventure"
		],
		"synopsis": "The Lord of the Rings is an epic high-fantasynovel by English author and scholar J. R. R. Tolkien."
	},
	{
		"title": "The Hobbit",
		"author": "J. R. R. Tolkien",
		"category": [
			"Fantasy",
			"Adventure"
		],
		"synopsis": "The Hobbit is a children's fantasy novel by English author J. R. R. Tolkien."
    }
]
```

  - endpoint GET /books/:title
    * Lista o livro com o título passado na URL com status 200
    * Se o livro não existir, retorna o Status 404 com a mensagem
  
- ```json
  {
    "message": "Book not found"
  }
  ```

  - endpoint POST /books
   * Adiciona um livro ao banco de dados
   * O endpoint retorna o Status 201 com o body vazio

  
  - Validações do endpoint
  * O campo title é obrigatório
  * O campo author é obrigatório e deve ter no mínimo 3 caracteres
  * O campo category é obrigatório e deve ter no mínimo 1 categoria
  * O campo synopsis é obrigatório e deve ter no mínimo 30 caracteres
  
- Caso a validação do campo title falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	  "message": "title field is invalid"
   }
  ```
- Caso a validação do campo author falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	   "message": "author field is invalid, must be more than 3 characters"
   }
  ```
- Caso a validação do campo category falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
       "message": "category field is invalid, must be more than 1 category"
   }
  ```
- Caso a validação do campo synopsis falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
       "message": "synopsis field is invalid, must be more than 30 characters"
   }
  ```  
  - endpoint PUT /books/:title
    * Atualiza o livro com o título passado na URL, o title é o indice que é usado para atualizar o livro.
    * O endpoint retorna o Status 200 com o body vazio

  - Validações do endpoint
  * O campo title é obrigatório
  * O campo author é obrigatório e deve ter no mínimo 3 caracteres
  * O campo category é obrigatório e deve ter no mínimo 1 categoria
  * O campo synopsis é obrigatório e deve ter no mínimo 30 caracteres

  <details>
  <summary><strong>Retornos</strong></summary><br />
 - Caso a validação do campo title falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	  "message": "title field is invalid"
   }
  ```
- Caso a validação do campo author falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	   "message": "author field is invalid, must be more than 3 characters"
   }
  ```
- Caso a validação do campo category falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
       "message": "category field is invalid, must be more than 1 category"
   }
  ```
- Caso a validação do campo synopsis falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
       "message": "synopsis field is invalid, must be more than 30 characters"
   }
  ```  
</details>

   - endpoint DELETE /books/:title
    * Deleta o livro com o título passado na URL e retorna o Status 204 com o body vazio
    * Se o livro não existir, retorna o Status 404 com a mensagem:
    ```json
    {
      "message": "Book not found"
    }
    ```
   
-
## O Desenvolvimento

Foi utilizado o método TDD, para o desenvolvimento. Foi escolhido chi como router por
este possuir o melhor benchmark de performance, e o banco de dados escolhido foi o MongoDB. Todo ele dockerizado.

-> Justificativa do banco de dados escolhido
  * No caso dado que é apenas uma entidade ali de Livro, não teríamos muitos pontos para se escolher MongoDB ou MySQL.
    Pois, poderíamos pensar que relacionamentos poderiam fazer com que as queries ficassem pesadas, mas, no caso de uma API simples de uma entidade não se teria tanta diferença.
    Mas, no geral, temos alguns pontos.O MongoDB foi o escolhido porque teria de retornar muitas vezes o livro com o seu autor, neste sentido o MongoDB é melhor, pois, estaria tudo em um documento apenas.
    Pois, entra no sentido que se fosse MySQL faríamos um JOIN, e comumente, ele é uma operação dita de razoável processamento, muitos JOIN’s, significam um peso muito grande de processamento. Já o MongoDB não teria esse peso dos JOIN’s.

### Ferramentas usadas

- [Golang](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Chi](https://github.com/go-chi/chi)
- [Docker](https://www.docker.com/)

## Uso

- Clone o repositório
```bash
git clone git@github.com:beto-ouverney/nikiti-books-golang.git
```
- Entre na pasta do projeto

```bash
cd nikiti-books-golang
```

- Utilize o comando abaixo para subir o container do MongoDB com os bancos
```bash
docker-compose -f docker-compose.dev.yml up -d --build
```
- exitem dois bancos, um para uso continuo e outro somente para testes,
- para que não haja conflito entre os dados de teste e os dados de uso continuo.
- Por padrão ele vem setado para o banco de teste, para mudar deve entrar na pasta config e comentar as constantes MONGO_CONNECT e PORT e descomentar as de produção

- Para rodar o projeto

```bash
go run main.go
```

## Test

* O Banco de testes deve estar ativo e as variaveis de ambientes de testes também, senão nao rodará todos os testes somente os mockados.
    - Nos testes dos handlers, o o banco de dados é populado no inico de cada teste e depois é feito um drop.
- Para rodar os testes
```bash
     go test -v ./... 
```

- Para rodar os testes com cores para diferenciar os erros mais facilmente
```bash
  go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
  ```

## Author

- LinkedIn - [Alberto Ouverney Paz](https://www.linkedin.com/in/beto-ouverney-paz/)
