📚 GO-BOOKSTORE
Este repositório contém um projeto em Go chamado GO-BOOKSTORE, que demonstra como conectar o Go com o MySQL.

🚀 Como começar
Para começar com este projeto, siga as etapas abaixo:

Clone o repositório:

shell
Copiar
git clone https://github.com/herondi/GO-BOOKSTORE.git
Instale as dependências necessárias. Certifique-se de ter o Go e o MySQL instalados em seu sistema. As dependências específicas do projeto podem ser encontradas no arquivo go.mod.

Configure o banco de dados do MySQL. Você pode encontrar os scripts SQL necessários na pasta db deste repositório.

Configure a conexão com o banco de dados no arquivo config.go.

Compile e execute o projeto:

shell
Copiar
go build
./GO-BOOKSTORE
Acesse a aplicação em seu navegador da web em http://localhost:9010.

🔧 Dependências
O projeto utiliza as seguintes dependências:

"encoding/json"
"fmt"
"net/http"
"strconv"
"github.com/gorilla/mux"
"github.com/herondi/GO-BOOKSTORE/pkg/models"
"github.com/herondi/GO-BOOKSTORE/pkg/utils"
Certifique-se de instalá-las antes de executar o projeto.

🎉 Recursos
Autenticação e autorização de usuários
Gerenciamento de livros (adicionar, editar, excluir)
Gerenciamento de usuários (registro, login, logout)
🤝 Contribuição
Contribuições são bem-vindas! Se você tiver ideias, sugestões ou relatórios de bugs, abra uma issue ou envie um pull request.

📄 Licença
Este projeto está licenciado sob a Licença MIT.

💡 Agradecimentos
Este projeto foi inspirado pela necessidade de demonstrar como integrar o Go e o MySQL em um aplicativo de livraria.
