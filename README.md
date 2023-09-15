# Uso de Estruturas em Go

Este é um exemplo de código em Go que demonstra o uso de estruturas (structs) para criar objetos complexos e ilustrar conceitos como herança e métodos em tipos personalizados. Também inclui um exemplo de como criar um método personalizado para serialização JSON.
Este repositório é a fins de estudo para POO em goLang já que "não existe" :)

## Estruturas de Dados

### Student (Aluno)

A estrutura `Student` representa um aluno com um campo de nome.

### Course (Curso)

A estrutura `Course` representa um curso com uma descrição e uma lista de alunos inscritos. Inclui um método `register` para inscrever novos alunos.

### EadCourse (Curso EAD)

A estrutura `EadCourse` é uma extensão de `Course` que adiciona um campo de site para cursos a distância.

### NewFloat (Novo Float)

O tipo `NewFloat` é uma nova representação de números de ponto flutuante com um método `roundBy` para arredondamento personalizado.

### Json (JSON)

A estrutura `Json` e seu método `toJson` permitem a criação de objetos JSON personalizados, clássico hehehehehe.

## Exemplos de Uso

O código no arquivo `main.go` demonstra o uso das estruturas e tipos personalizados mencionados acima. Ele cria instâncias de alunos, cursos e realiza operações, como inscrição em cursos e arredondamento de números de ponto flutuante.

```go
// Exemplo de uso
var student Student
student.name = "Higor"

var otherStudent Student
otherStudent.name = "Gohan"

person := Student{"Goku"}

engineering := Course{"Engineering", make([]Student, 0)}
engineering.register(person)
engineering.register(otherStudent)

ead := EadCourse{
    Course{"New EAD", make([]Student, 0)},
    "https://myead.com"}

var number NewFloat = 5.0293019384
