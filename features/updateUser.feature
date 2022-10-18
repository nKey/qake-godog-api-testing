      #language: pt
      Funcionalidade: put user
      Atualização das informações do usuários

      Cenario: Como um usuário gostaria de atualizar minhas informações pessoais no sistema
      Dado Enviarei as seguintes informacoes atualizadas
      """
      "{\"postId\": 5, \"name\": \Maykon Douglas\", \"email\": \"Kariane@jadyn.tv\", "body": \"fuga eos qui dolor rerum\ninventore corporis exercitationem\ncorporis cupiditate et deserunt recusandae est sed quis culpa\neum maiores corporis et\"}"
      """
      Quando Vou enviar "PUT" uma requisicao para "https://jsonplaceholder.typicode.com/users/5"
      Então Verificar o status code 200
      E Por fim verificar o corpo da requisicao criada
      """
      {
        "id": 5
      }
      """

