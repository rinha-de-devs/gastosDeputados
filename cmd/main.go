package main

import "deputySpending/internal/adapter/conexao"

// 1 passo modelo de domínio <check>
// 2 passo conexão com a página <check>
// 3 passo printar os dados <check>
// 4 passo sanatização de dados

func main() {
	conn := conexao.DeputadoRepository{}

	conn.BuscaDeputado(conn.BuscaDeputados)

	//fmt.Printf("%+v\n", deputado)

}
