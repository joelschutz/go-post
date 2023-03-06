# New file

## ASTs
Abstract Syntax Tree é o tipico tema que reramente aparece no nosso dia a dia com desenvolvimento
mas que no dia que precisamos dela, vira um problema sem tamanho. Eu vou mostrar que nao tem
motivo para ter medo e que talvez devesse ate dar um pouco mais de carinho para essas ferramentas.
Podem te ajudar com tarefas que nem imaginava que poderia utilizar e resolver a sua vida na
hora de programar diariamente.

Primeiramente, oque sao essas arvore$ De onde vem$ o que fazem$ onde dormem$ A ultima é
facil: elas nunca dormem. Quando as outras, vamos ser breves pq eu imagino que o exemplo vai
deixar tudo mais claro.
- Oque.
	- AST nada mais é que uma forma de representar codigo de forma semantica, ou seja, ao inves
	de ver o codigo como uma string de caracteres. A arvore contem todo os elementos presentes
	no script: variaveis, funcoes, importacoes, valores, loops, etc. Ela tbm organiza esses
	elememtos de forma hierarquica: valor pertence a variavel, que pertence a funcao, que pertence a classe...

- Onde.
	- AST sao uma ferramenta importante em qualquer linguagem, é utilizando essas arvores que
	os compiladres e interpretadores "leem" o seu codigo antes de traduzi-lo para a maquina.Toda
	linguagem possui pelo menos uma definicao(podem haver mais) de sua AST pois no fim é ela
	quem torna possivel a linguagem ser executada.
	- Nao podemos esquecer que ha mais etapas alem da AST para compilar um programa, o que
	vamos ver é que com essa ferramentas conseguimos "ler" o codigo da mesma forma que o
	compilar faz

- Como.
	- A forma de lidar com essas estruturas é "caminhando" por elas, voce cria um visitante
	que ira passar por cada nó da arvore e fazer alguma operacao. Cada no contem informacoes
	sobre sua posicao no arquivo, conteudo, tipo, identificadores(nomes basicamente) e seus
	filhos, caso existam.
	- Um uso comum dessas estruturas são nos linters, quando criamos um plugin para alguma
	ferramenta dessas precisamos sempre comecar por um visitante que atravessa a arvore acusando
	infracoes as regras definidas por ele.

## Demo
Esse é um daaqueles projetos que eu não sei se quem veio primeiro foi o problema ou a solução,
mas se certa forma ele é os dois ao mesmo tempo. A ideiai é a seguinte: eu gosto de comentar
o meu codigo, mas raramente uso eles para escrever documentacao ou artigos como esse. Se eu
tivesse alguma ferramenta que pudesse extrair esses comentarios e codigos mais importantes para
um arquivo .md, eu poderia facilmente adiciona-los ao restante do material. Isso que essa demo
faz.

Esse artigo em si esta sendo escrito no codigo fonte desse programa, e se estiver lendo é por
que o programa funciona. Mas não so isso, olhe esse struct:

```go
type MDParser struct {
	cells        [][2]int // (cellType, extIndex)
	txtSegments  []*ast.CommentGroup
	declSegments []ast.Decl
	pins         []ast.Node
	File         []byte
}

```

O programa consegue trazer extrair do codigo fonte utilizando apenas um comentario ˋ// PINˋ
antes do bloco de codigo.Não é incrivel$ Sugiro que seja o arquivo [parser.go](./internal/parser.go)
e veja que esta tudo la. Pode inclusive tentar voce mesmo:

ˋˋˋbash

go run parser.go ./parser.go

ˋˋˋ

