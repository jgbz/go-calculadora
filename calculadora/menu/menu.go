package menu

import (
	"fmt"
	"strings"
	"strconv"
	"calculadora/soma"
	"calculadora/subtracao"
	"calculadora/multiplicacao"
	"calculadora/divisao"
)

func Menu(){
	var escolha string
	
	fmt.Println("Calculadora")		
	fmt.Println("Insira a operacao")
	fmt.Scan(&escolha);
	listSimbolos := make([] string, 0);
	listValores := make([] string, 0);
	
	for _, simbolo := range escolha{		
		if (simbolo == 43 || simbolo == 45 || simbolo == 42 || simbolo == 47) {
			listSimbolos = append(listSimbolos, strconv.QuoteRune(simbolo));
			listValores = append(listValores, "-");
		} else {
			listValores = append(listValores, strconv.QuoteRune(simbolo));
		}
	}
	
	
	//~ for i, simbolo := range listSimbolos{
		//~ fmt.Println(simbolo);
		//~ fmt.Println(i);
	//~ }
	//~ fmt.Println(listValores);
	//~ fmt.Println(len(listSimbolos));
	escolhaSlice := strings.Split(escolha, "");
	valorIte := 0.0
	proximaOperacao := ""
	for _ , simbolo := range listSimbolos{
		indice := strings.IndexAny(escolha, simbolo);
		escolhaSlice2 := escolhaSlice[ :indice]
		
		
		
		
		
		fmt.Println("_____________");
		fmt.Println(escolhaSlice2);
		fmt.Println("_____________");
		
		valorEscolhaSlice, _ := strconv.ParseFloat(strings.Join(escolhaSlice2, ""), 64);
		escolhaSlice = escolhaSlice[indice+1:]
		
		fmt.Println("_____________");
		fmt.Println(valorEscolhaSlice);
		fmt.Println("_____________");
		
		//~ valorIte,  = strconv.ParseFloat(strings.Join(escolhaSlice2, ""), 64);
		
		fmt.Println("_____________");
		fmt.Println(escolhaSlice);
		fmt.Println("_____________");
		
		switch {
		case strings.Contains(simbolo, "+") == true:
			valorIte = soma.Somar(valorIte, valorEscolhaSlice);
		case strings.Contains(simbolo, "-") == true:
			valorIte = subtracao.Subtrair(valorIte, valorEscolhaSlice);
		case strings.Contains(simbolo, "*") == true:
			valorIte = multiplicacao.Multiplicar(valorIte, valorEscolhaSlice);
		case strings.Contains(simbolo, "/") == true:
			valorIte = divisao.Dividir(valorIte, valorEscolhaSlice);
		}			
	}
	
	fmt.Println(valorIte);
	
	
	//~ for i, _ := range listValores{
		//~ if (listValores[i-1] != " "){
			//~ listValores[i-1] = listValores[i-1]+listValores[i]
			//~ listValores[i] = " "
			
		//~ }
	//~ }
	//~ fmt.Println(len(listValores));
	
	//~ fmt.Println(listSimbolos);
	//~ fmt.Println(listValores);
		
	
	//~ fmt.Println("");
	Menu();	
}

func separar(a, b string) (v1, v2 string){
	fmt.Println(a);
	v := strings.Split(a, b);
	v1 = v[0]; v2 = v[1];
	
	return v1, v2
}
