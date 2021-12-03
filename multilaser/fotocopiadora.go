package fotocopiadora

// import "./impressora"
import "github.com/megatron0000/golang-meuModulo/multilaser/impressora"
// import "./scanner"
import "github.com/megatron0000/golang-meuModulo/multilaser/scanner"

func ScanAndPrint() {
	impressao.Imprimir()
	scanner.Scanear()
}