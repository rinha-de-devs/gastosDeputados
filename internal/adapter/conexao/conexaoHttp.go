package conexao

import (
	"deputySpending/internal/domain"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func BuscaDeputado()domain.Deputado{

	site := "https://www.camara.leg.br/transparencia/gastos-parlamentares?legislatura=56&ano=2021&mes=&por=deputado&deputado=204554&uf=&partido="

	response, err := http.Get(site)

	defer response.Body.Close()

	if err != nil {
		fmt.Printf("FALHA AO EXECUTAR REQUISICAO %d %s",
			response.StatusCode, response.Status)
		panic(err.Error())
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	nome, err := pegaNome(*doc)
	if err != nil{
		log.Fatal(err)
	}

	var deputado domain.Deputado
	deputado.Nome = nome

	cota, err := pegaCota(*doc)
	if err != nil{
		log.Fatal(err)
	}

	deputado.Cota = cota

	gabineteGasto, err := pegaVerbaDeGabineteGasto(*doc)
	if err != nil {
		log.Fatal(err)
	}

	deputado.VerbaDeGabineteGasto = gabineteGasto.verbaGasta
	deputado.PorcentagemGasto = gabineteGasto.porcentagemGasta

	gabineteDisponível, err := pegaVerbaDeGabineteDisponível(*doc)
	if err != nil {
		log.Fatal(err)
	}

	deputado.VerbaDeGabineteDisponoivel = gabineteDisponível.verbaDisponivel
	deputado.PorcentagemDisponoivel = gabineteDisponível.porcentagemDisponivel

	deputado.Partido = "Nao implementado"
	deputado.Estado = "Nao implementado"

	return deputado

}

func pegaNome(document goquery.Document)(string, error){
	nome := document.Find("#main-content > section.gastos-form > div.gastos-form__resumo-resposta > div > p > span:nth-child(1)").Text()

	if len(nome) == 0 {
		return "", errors.New("Nome não encontrado")
	}

	return nome, nil
}

func pegaCota(document goquery.Document)(string, error){
	cota := document.Find("#cota > div > div.l-cota__row > div:nth-child(1) > div > div.l-card.l-cota-resumo > div > div > section > p.gastos__resumo-texto.gastos__resumo-texto--destaque > span").Text()

	if len(cota) == 0 {
		return "", errors.New("Cota não encontrada")
	}

	return cota, nil
}

func pegaVerbaDeGabineteGasto(document goquery.Document)(struct{verbaGasta string
										                        porcentagemGasta string}, error){
	verbaGasta := document.Find("#js-percentual-gasto > tbody > tr:nth-child(2) > td:nth-child(2)").Text()

	if len(verbaGasta) == 0 {
		return struct {
			verbaGasta       string
			porcentagemGasta string
		}{}, errors.New("Verba Gasta não encontrada")
	}

	porcentagemGasta := document.Find("#js-percentual-gasto > tbody > tr:nth-child(1) > td:nth-child(3)").Text()

	if len(porcentagemGasta) == 0 {
		return struct {
			verbaGasta       string
			porcentagemGasta string
		}{}, errors.New("Porcentagem Gasta não encontrada")
	}

	return struct {
		verbaGasta       string
		porcentagemGasta string
	}{verbaGasta, porcentagemGasta }, nil
}

func pegaVerbaDeGabineteDisponível(document goquery.Document)(struct{verbaDisponivel string
	porcentagemDisponivel string}, error){
	verbaDisponivel := document.Find("#js-percentual-gasto > tbody > tr:nth-child(1) > td:nth-child(2)").Text()

	if len(verbaDisponivel) == 0 {
		return struct {
			verbaDisponivel       string
			porcentagemDisponivel string
		}{}, errors.New("Verba Disponivel não encontrada")
	}

	porcentagemDisponivel := document.Find("#js-percentual-gasto > tbody > tr:nth-child(2) > td:nth-child(3)").Text()

	if len(porcentagemDisponivel) == 0 {
		return struct {
			verbaDisponivel       string
			porcentagemDisponivel string
		}{}, errors.New("Porcentagem Disponivel não encontrada")
	}

	return struct {
		verbaDisponivel       string
		porcentagemDisponivel string
	}{verbaDisponivel, porcentagemDisponivel}, nil
}

