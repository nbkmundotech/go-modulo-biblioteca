// Package consolidar tem funcoes para converter um mapa de chaves
// e valores para uma query string.
package consolidar

import "fmt"

// MapaParaString consolida as chaves e valores de um mapa e cria uma string com
// o seguinte formato:
//		"?chave1=valor1&chave2=valor2"
//
// O mapa de entrada segue o seguinte formato:
//		{"chave1": "valor1", "chave2": "valor2"}
func MapaParaString(mapa map[string]string) string {
	var queryString = "?"
	var indice = 0
	for chave, valor := range mapa {
		if indice != 0 {
			queryString += "&"
		}
		queryString += fmt.Sprintf("%v=%v", chave, valor)
		indice++
	}

	return queryString
}
