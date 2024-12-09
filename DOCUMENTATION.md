# Rodar

### Pré-requisitos

- Instalar NPM e NPX (https://www.npmjs.com/get-npm)
- Instalar Docker and Docker Compose (https://www.docker.com/)
- Instalar Besu (https://besu.hyperledger.org/private-networks/get-started/install/binary-distribution)
- Instalar Go (https://golang.org/dl/)
- Instalar git, para clonar o repository https://github.com/luan441/goledger-challenge-besu

### Setup

1. Clonar repository

```bash
git clone https://github.com/luan441/goledger-challenge-besu
```

2. Entrar no diretório do besu, instalar Hardhat rodar e iniciar a rede besu

Para entrar no diretório besu rode o seguinte comando

```bash
cd besu
```

Agora para instalar o Hardhat

```bash
npm install --save-dev hardhat
```

E por fim para iniciar a rede besu, basta rodar o seguinte comando

```bash
./startDev.sh
```

Pegue o codigo hexdecimal gerado na ultima linha pelo como `./startDev.sh` e substitua ele nas seguintes linhas do arquivo `internal/besu/interaction.go`

- Linha: 105
- Linha: 202

3. Iniciar o banco de dados

Volte para o diretório raiz e inicie os containers docker

```bash
cd ..
docker compose up -d
```

4. Inciar servidor GO

```bash
go run cmd/server/main.go
```

### Iniciar

Com o setup feito sempre que quiser iniciar a aplicaçcão basta rodar os seguintes comando em ordem:

```bash
cd besu
./startDev.sh
cd ..
docker compose up -d
go run cmd/server/main.go
```

# Arquitetura da aplicação

A arquitetura foi pensada de maneira simples, apenas para evitar arquivos grandes que dificutam o entendimento.

No arquivo `cmd/server/main.go` fica o start da API com as definições das rotas.

Na pasta `internal` fica toda a logica do app.

- No diretório `besu` fica a logica de executar o contrato, gravando na rede besu um valor inteiro, e também a função responsável por recuperar o valor da rede.

- Já a pasta `database` contém o arquivo que contém a conexão com o banco de dados.

- Em seguida na pasta `entiry` possui a `struct` que representa o modelo dos dados no banco de dados.

- Na pasta `handlers` contém as funções que lidaram como dada rota registrada, nelas contém tratamento de erros como as manipulações do banco de dados e dá rede, os handlers são análogo as controllers na design partner MVC, coloquei esse nome pois a biblioteca http do GO na hora de definir as rotas ele usa a função Handle, quiz deixar algo mais proximo disso.

- Por fim temos a pasta `repository` que implementa o padrão de mesmo nome para minipulação do banco de dados, como pode ver obtei por receber a conexão com o banco de dados nos handlers para facilitar nas mensagens de erro referente ao banco de dados.

Também tem uma pasta `scripts` nela tem o sql responsável pela criação da tabela, não há necessidade de executá-lo pois eu configurei o container do banco de dados para fazer isso.
