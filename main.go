package main

import "fmt"

type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio *no
}

// Exercício 2
func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor}
	novo.proximo = l.inicio
	l.inicio = novo
}

func (l *lista) adicionarFim(valor int) {
	novo := &no{valor: valor}

	if l.inicio == nil {
		l.inicio = novo
		return
	}

	atual := l.inicio
	for atual.proximo != nil {
		atual = atual.proximo
	}
	atual.proximo = novo
}

// Exercício 3
func (l *lista) adicionarPosicao(valor, posicao int) bool {
	if posicao < 0 {
		return false
	}

	if posicao == 0 {
		l.adicionarInicio(valor)
		return true
	}

	anterior := l.inicio
	for i := 0; i < posicao-1 && anterior != nil; i++ {
		anterior = anterior.proximo
	}

	if anterior == nil {
		return false
	}

	novo := &no{valor: valor}
	novo.proximo = anterior.proximo
	anterior.proximo = novo
	return true
}

// Exercício 4
func (l *lista) removerInicio() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	removido := l.inicio
	l.inicio = removido.proximo
	return removido.valor, true
}

func (l *lista) removerFim() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	if l.inicio.proximo == nil {
		valor := l.inicio.valor
		l.inicio = nil
		return valor, true
	}

	anterior := l.inicio
	for anterior.proximo.proximo != nil {
		anterior = anterior.proximo
	}

	removido := anterior.proximo
	anterior.proximo = nil
	return removido.valor, true
}

// Exercício 5
func (l *lista) removerPosicao(posicao int) (int, bool) {
	if posicao < 0 || l.inicio == nil {
		return 0, false
	}

	if posicao == 0 {
		return l.removerInicio()
	}

	anterior := l.inicio
	for i := 0; i < posicao-1 && anterior != nil; i++ {
		anterior = anterior.proximo
	}

	if anterior == nil || anterior.proximo == nil {
		return 0, false
	}

	removido := anterior.proximo
	anterior.proximo = removido.proximo
	return removido.valor, true
}

// Exercício 6
func (l *lista) posicao(valorProcurado int) (int, bool) {
	atual := l.inicio
	posicao := 0

	for atual != nil {
		if atual.valor == valorProcurado {
			return posicao, true
		}
		atual = atual.proximo
		posicao++
	}

	return 0, false
}

// Exercício 7
func (l *lista) valorNaPosicao(posicaoProcurada int) (int, bool) {
	if posicaoProcurada < 0 {
		return 0, false
	}

	atual := l.inicio
	posicao := 0

	for atual != nil {
		if posicao == posicaoProcurada {
			return atual.valor, true
		}
		atual = atual.proximo
		posicao++
	}

	return 0, false
}

// Exercício 8
func (l *lista) tamanho() int {
	total := 0

	for atual := l.inicio; atual != nil; atual = atual.proximo {
		total++
	}

	return total
}

func (l *lista) imprimir() {
	for atual := l.inicio; atual != nil; atual = atual.proximo {
		fmt.Printf("%d -> ", atual.valor)
	}
	fmt.Println("nil")
}

func main() {
	var l lista
	var opcao, valor, posicao int

	for {
		fmt.Println("\n--- LISTA ENCADEADA ---")
		fmt.Println("1 - Adicionar no início")
		fmt.Println("2 - Adicionar no fim")
		fmt.Println("3 - Adicionar em uma posição")
		fmt.Println("4 - Remover do início")
		fmt.Println("5 - Remover do fim")
		fmt.Println("6 - Remover de uma posição")
		fmt.Println("7 - Procurar posição de um valor")
		fmt.Println("8 - Consultar valor em uma posição")
		fmt.Println("9 - Mostrar lista e tamanho")
		fmt.Println("0 - Sair")
		fmt.Print("Opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarInicio(valor)

		case 2:
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarFim(valor)

		case 3:
			fmt.Print("Valor e posição: ")
			fmt.Scan(&valor, &posicao)
			if !l.adicionarPosicao(valor, posicao) {
				fmt.Println("Posição inválida.")
			}

		case 4:
			valor, ok := l.removerInicio()
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Lista vazia.")
			}

		case 5:
			valor, ok := l.removerFim()
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Lista vazia.")
			}

		case 6:
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.removerPosicao(posicao)
			if ok {
				fmt.Println("Removido:", valor)
			} else {
				fmt.Println("Posição inválida ou lista vazia.")
			}

		case 7:
			fmt.Print("Valor procurado: ")
			fmt.Scan(&valor)
			posicao, ok := l.posicao(valor)
			if ok {
				fmt.Println("Encontrado na posição:", posicao)
			} else {
				fmt.Println("Valor não encontrado.")
			}

		case 8:
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.valorNaPosicao(posicao)
			if ok {
				fmt.Println("Valor:", valor)
			} else {
				fmt.Println("Posição inválida.")
			}

		case 9:
			l.imprimir()
			fmt.Println("Tamanho:", l.tamanho())

		case 0:
			return

		default:
			fmt.Println("Opção inválida.")
		}
	}
}
