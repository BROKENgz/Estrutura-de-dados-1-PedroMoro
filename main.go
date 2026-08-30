package main

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
