package main

import (
	"deputySpending/internal/adapter/conexao"
	"encoding/json"
	"fmt"
	"log"
)

// 1 passo modelo de domínio <check>
// 2 passo conexão com a página <check>
// 3 passo printar os dados <check>
// 4 passo sanatização de dados


func main()  {

	deputado := conexao.BuscaDeputado()

	fmt.Printf("%+v\n", deputado)

	marshal, err := json.MarshalIndent(deputado, "", "")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%s\n", marshal)

}