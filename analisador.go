package analisador

import "strings"

// ?busca=termo&ordem=desc
//   => map[string]string{ "busca": "termo", "ordem": "desc" }

func Analisar(queryString string) map[string]string {
	var pares []string = strings.Split(queryString[1:], "&")
	var mapa = map[string]string{}
	for i := 0; i < len(pares); i++ {
		var chaveValor []string = strings.Split(pares[i], "=")
		var chave string = chaveValor[0]
		var valor string = chaveValor[1]

		mapa[chave] = valor
	}
	return mapa
}
