# go-with-tests

Meu percurso pelo livro **[Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)**, de Chris James — aprendendo Go pela prática de **TDD**, escrevendo o teste antes da implementação em cada tópico.

Não é uma biblioteca nem um produto: é um repositório de estudo, com cada conceito da fundação da linguagem isolado em sua própria pasta, sempre com implementação (`*.go`) e teste (`*_test.go`).

## Tópicos praticados

| Pasta | Conceito de Go |
|---|---|
| `HelloWorld` | primeiro teste, subtestes, iteração no ciclo red-green-refactor |
| `Integers` | testes de exemplo e documentação executável |
| `Iteration` | benchmarks e loops |
| `ArraysAndSlices` | `range`, cobertura de teste, `reflect.DeepEqual` |
| `StructsMethodsInterfaces` | métodos, interfaces e polimorfismo |
| `Pointers` | ponteiros, tipos próprios e erros |
| `Maps` | mapas e erros customizados |
| `DependencyInjection` | injeção de dependência sem framework |
| `Mocking` | mocks escritos à mão e o que eles ensinam sobre design |

## Como rodar

Os módulos ficam em `GoFundamentals/`:

```bash
cd GoFundamentals
go test ./...
```

## Por que TDD

O livro ensina Go, mas o hábito que fica é o ciclo: escrever o teste que falha, fazer passar com o mínimo, refatorar com a rede de segurança pronta. É a disciplina que eu levo para os projetos de verdade, não só para o exercício.
