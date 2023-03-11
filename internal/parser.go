/*
Copyright © 2023 JOEL SCHUTZ <JOELSSCHUTZ@YAHOO.COM.BR>

*/
package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

const (
	TEXT = iota
	PIN
)

/*POST
## ASTs
Abstract Syntax Tree é o tipico tema que reramente aparece no nosso dia a dia com desenvolvimento
mas que no dia que precisamos dela, vira um problema sem tamanho. Eu vou mostrar que nao tem
motivo para ter medo e que talvez devesse ate dar um pouco mais de carinho para essas ferramentas.
Podem te ajudar com tarefas que nem imaginava que poderia utilizar e resolver a sua vida na
hora de programar diariamente.

Primeiramente, oque sao essas arvore? De onde vem? o que fazem? onde dormem? A ultima é
facil: elas nunca dormem. Quando as outras, vamos ser breves pq eu imagino que o exemplo vai
deixar tudo mais claro.
- Oque?
	- AST nada mais é que uma forma de representar codigo de forma semantica, ou seja, ao inves
	de ver o codigo como uma string de caracteres. A arvore contem todo os elementos presentes
	no script: variaveis, funcoes, importacoes, valores, loops, etc. Ela tbm organiza esses
	elememtos de forma hierarquica: valor pertence a variavel, que pertence a funcao, que pertence a classe...

- Onde?
	- AST sao uma ferramenta importante em qualquer linguagem, é utilizando essas arvores que
	os compiladres e interpretadores "leem" o seu codigo antes de traduzi-lo para a maquina.Toda
	linguagem possui pelo menos uma definicao(podem haver mais) de sua AST pois no fim é ela
	quem torna possivel a linguagem ser executada.
	- Nao podemos esquecer que ha mais etapas alem da AST para compilar um programa, o que
	vamos ver é que com essa ferramentas conseguimos "ler" o codigo da mesma forma que o
	compilar faz

- Como?
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
*/

// PIN
type MDParser struct {
	cells        [][2]int            // Representa um bloco de texto com estrutura: (cellType, extIndex)
	txtSegments  []*ast.CommentGroup // Array de commentarios marcados com `POST`
	declSegments []ast.Decl          // Array de declaracoes marcados com `PIN`
	pins         []ast.Node          // Array com a localizacao dos marcadores `PIN`
	File         []byte              // Conteudo do arquivo alvo
}

/*POST
O programa consegue extrair do codigo fonte utilizando apenas um comentario ˋ// PINˋ
antes do bloco de codigo.Não é incrivel? Sugiro que abra o arquivo [parser.go](./internal/parser.go)
e veja que esta tudo la. Pode inclusive tentar voce mesmo:

ˋˋˋbash

go build main.go
./main ./internal/parser.go # Mac/Linux
./main.exe ./internal/parser.go # Windows

ˋˋˋ

O funcionamento por enquanto é bastante simples, a classe MDParser implementa apenas 3 métodos:
*/

// PIN
func (p *MDParser) parseComments(c []*ast.CommentGroup) error {
	for _, tk := range c {
		if strings.HasPrefix(tk.Text(), "PIN") {
			p.pins = append(p.pins, tk)
			p.cells = append(p.cells, [2]int{PIN, len(p.pins) - 1})
		} else if strings.HasPrefix(tk.Text(), "POST") {
			p.txtSegments = append(p.txtSegments, tk)
			p.cells = append(p.cells, [2]int{TEXT, len(p.txtSegments) - 1})
		}
	}

	return nil
}

/*POST
Em parseComments os commentarios extraidos do arquivo sao classificados entre `PIN` e `POST`,
em seguida adicionados aos arrays correspondentes
*/

// PIN
func (p *MDParser) parseDeclarations(decl []ast.Decl) error {
	for _, tk := range decl {
		for _, v := range p.pins {
			if v.End()+1 == tk.Pos() {
				p.declSegments = append(p.declSegments, tk)
			}
		}
	}

	return nil
}

/*POST
Em parseDeclarations extraimos as declaracoes que foram marcadas com `PIN` que identificamos no
metodo anterior.
*/

// PIN
func (p MDParser) Flush(title string) string {
	s := fmt.Sprintf("# %s\n\n", title)

	for _, cell := range p.cells {
		switch cell[0] {
		case PIN:
			s += fmt.Sprintf("```go\n%s```\n\n", string(p.File[p.pins[cell[1]].End():p.declSegments[cell[1]].End()]))
		case TEXT:
			s += strings.TrimPrefix(p.txtSegments[cell[1]].Text(), "POST\n") + "\n"
		}
	}
	return s
}

/*POST
Por fim, o metodo Flush gera e retorna um arquivo Markdown com todo o conteudo extraido. Essa
string pode ser salva em um novo arquivo, renderizado na tela, o que for necessario

E para amarrar tudo, existe um construtor que realiza todo o processo dado o endereco de um arquivo
alvo:
*/

// PIN
func NewMDParserFromFile(targetFile string) (*MDParser, error) {
	// Criamos a AST do arquivo
	fs := token.NewFileSet()
	fTree, err := parser.ParseFile(fs, targetFile, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Extraimos o conteudo do arquivo
	buf, err := ioutil.ReadFile(targetFile)
	if err != nil {
		return nil, err
	}

	// Devolvemos um ponteiro para o objeto ja parseado
	p := &MDParser{File: buf}
	p.parseComments(fTree.Comments)
	p.parseDeclarations(fTree.Decls)
	return p, nil
}
