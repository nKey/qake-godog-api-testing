#language: pt
Funcionalidade: get comment
  In order to know jsonplaceholder
  As an API user
  Eu need to be able to request body

  Cenario: should get comment json
    Quando I send "GET" request to "https://jsonplaceholder.typicode.com/comments?name=alias%20odio%20sit"
    Então the response code should be 200
    E the response should match json:
      """
      [
        {
          "postId": 1,
          "id": 4,
          "name": "alias odio sit",
          "email": "Lew@alysha.tv",
          "body": "non et atque\noccaecati deserunt quas accusantium unde odit nobis qui voluptatem\nquia voluptas consequuntur itaque dolor\net qui rerum deleniti ut occaecati"
        }
      ]
      """
