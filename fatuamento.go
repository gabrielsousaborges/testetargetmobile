package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Faturamento(faturamentoDia []float64) (float64, float64, int, error) {
   
    
	var faturamentoValido []float64
	for _, valor := range faturamentoDia {
		if valor > 0 {
			faturamentoValido = append(faturamentoValido, valor)
		}
	}
    
    diasMaior := 0

	// Calcular média
	var soma float64
	for _, valor := range faturamentoValido {
		soma += valor
	}
	media := soma / float64(len(faturamentoValido))

    menorValor := faturamentoValido[0]
	maiorValor := faturamentoValido[0]
	for _, valor := range faturamentoValido {
        
		if valor < menorValor {
			menorValor = valor
		}
        
		if valor > maiorValor {
			maiorValor = valor
		}
	}

	
	for _, valor := range faturamentoValido {
		if valor > media {
			diasMaior++
		}
	}

	return menorValor, maiorValor, diasMaior, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite os valores separados por virgula: ")
	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
    
    var valores []float64
	for _, valorString := range strings.Split(entrada, ",") {
		valor, err := strconv.ParseFloat(valorString, 64)
		if err != nil {
			fmt.Println("Erro ao converter valor:", valorString)
			return
		}
		valores = append(valores, valor)
	}
    
	if len(valores) == 0 {
		fmt.Println("Nenhum valor de faturamento fornecido.")
		return
	}

	menor, maior, diaMaior, err := Faturamento(valores)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Menor valor de faturamento: %v\n", menor)
	fmt.Printf("Maior valor de faturamento: %v\n", maior)
	fmt.Printf("Número de dias com faturamento acima da média: %v\n", diaMaior)
}
