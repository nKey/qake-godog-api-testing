      #language: pt
      Funcionalidade: post user
      Cadastro ddas informações de um usuário usuários

      Cenario: Como um usuário gostaria de me cadastrar no sistema, preenchendo todas as informações
      Dado Enviarei as seguintes informacoes
      """
      "{\"postId\": 5, \"name\": \"Douglas Queiroz2\", \"email\": \"Kariane@jadyn.tv\", "body": \"fuga eos qui dolor rerum\ninventore corporis exercitationem\ncorporis cupiditate et deserunt recusandae est sed quis culpa\neum maiores corporis et\"}"
      """
      Quando Vou enviar "POST" uma requisicao para "https://jsonplaceholder.typicode.com/users"
      Então Verificar o status code 201
      E Por fim verificar o corpo da requisicao criada
      """
      {
        "id": 11
      }
      """

