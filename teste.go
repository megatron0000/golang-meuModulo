package main

import "fmt"
import "meuModulo.com/multilaser/impressora"
import "meuModulo.com/multilaser/scanner"
import "meuModulo.com/multilaser"


func main() {
	fmt.Println("Hello mundo")
	impressao.Imprimir()
	impressao.ImprimirLinha()
	scanner.Scanear()
	fotocopiadora.ScanAndPrint()
}

