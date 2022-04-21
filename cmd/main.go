package main

import "deputySpending/internal/adapter/conexao"

func main() {
	conn := conexao.DeputadoRepository{}

	conn.BuscaDeputado(conn.BuscaDeputados)

}
