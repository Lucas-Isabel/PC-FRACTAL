package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estrutura que representa a memória compartilhada
type SharedMemory struct {
	count int // contador de "caras"
	total int // total de lançamentos
}

// Função para simulação analógica (Método de Monte Carlo para estimar Pi)
func analogProcessing(memory *SharedMemory, numSimulations int, done chan bool) {
	// Gerador de números aleatórios global
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Raio do círculo inscrito no quadrado unitário
	radius := 0.5

	// Contador de pontos dentro do círculo
	countInside := 0

	// Simulando pontos aleatórios e contando quantos estão dentro do círculo
	for i := 0; i < numSimulations; i++ {
		// Gerando coordenadas aleatórias no intervalo [-0.5, 0.5]
		x := random.Float64() - 0.5
		y := random.Float64() - 0.5

		// Verificando se o ponto está dentro do círculo
		if x*x+y*y <= radius*radius {
			countInside++
		}
	}

	// Estimativa de Pi usando o método de Monte Carlo
	estimatedPi := 4 * float64(countInside) / float64(numSimulations)

	fmt.Printf("Estimativa de Pi usando %d simulações: %.6f\n", numSimulations, estimatedPi)

	done <- true
}

// Função para processamento discreto (cálculo aritmético da probabilidade de obter pelo menos um número par ao lançar um dado seis vezes)
func discreteProcessing(memory *SharedMemory, numThrows int, done chan bool) {
	// Número de faces do dado
	numFaces := 6

	// Contando números pares
	countEven := 0

	// Calculando probabilidade usando cálculo aritmético
	for i := 1; i <= numFaces; i++ {
		if i%2 == 0 {
			countEven++
		}
	}

	// Probabilidade de não obter nenhum número par em n lançamentos
	probNoEven := pow(float64(numFaces-countEven)/float64(numFaces), float64(numThrows))

	// Probabilidade de obter pelo menos um número par
	probability := 1 - probNoEven

	fmt.Printf("Probabilidade de obter pelo menos um número par em %d lançamentos de dado: %.4f\n", numThrows, probability)

	done <- true
}

// Função auxiliar para calcular potência
func pow(base, exp float64) float64 {
	if exp == 0 {
		return 1
	}
	result := base
	for i := 1; i < int(exp); i++ {
		result *= base
	}
	return result
}

func main() {
	numSimulations := 3000000000
	numThrows := 6

	// Criando memória compartilhada
	memory := &SharedMemory{}

	// Canais para sincronização
	doneAnalog := make(chan bool)
	doneDiscrete := make(chan bool)

	// Executando processamento analógico (simulação) e processamento discreto (cálculo aritmético) em goroutines separadas
	go analogProcessing(memory, numSimulations, doneAnalog)
	go discreteProcessing(memory, numThrows, doneDiscrete)

	// Aguardando ambos os processamentos terminarem
	<-doneAnalog
	<-doneDiscrete
}
