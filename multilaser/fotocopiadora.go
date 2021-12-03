package fotocopiadora

// import "./impressora"
import "meuModulo.com/multilaser/impressora"
// import "./scanner"
import "meuModulo.com/multilaser/scanner"

func ScanAndPrint() {
	impressao.Imprimir()
	scanner.Scanear()
}