package main

type Funcionario interface {
	ProcessarPagamento()
}

type FuncionarioCLT struct {
	salario float64
}

type FuncionarioPJ struct {
	salario float64
}

func (f *FuncionarioCLT) ProcessarPagamento() {
	f.salario -= f.salario / 10
}

func (f *FuncionarioPJ) ProcessarPagamento() {
	f.salario -= f.salario / 10
}

func main() {
	var fnclt Funcionario
	var fnpj Funcionario

	clt := &FuncionarioCLT{1000}
	pj := &FuncionarioPJ{1000}

	fnclt = clt
	fnclt.ProcessarPagamento()

	fnpj = pj
	fnpj.ProcessarPagamento()

}
