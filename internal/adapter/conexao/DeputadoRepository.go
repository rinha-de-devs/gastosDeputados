package conexao

import (
	"deputySpending/internal/domain"
	"deputySpending/internal/ports/conexao"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type DeputadoRepository struct {
	conn *conexao.Conexao
}

func (d DeputadoRepository) BuscaDeputado(fn func() []domain.DeputadoResponse) {

	deputados := fn()
	var wg sync.WaitGroup

	wg.Add(len(deputados))

	for indice, dep := range deputados {

		time.Sleep(1 * time.Second)
		go func(id string, nome string, partido string, estado string) {
			defer wg.Done()

			fmt.Printf("Progresso: %d de %d\n", indice, len(deputados))

			url := fmt.Sprintf("https://www.camara.leg.br/transparencia/gastos-parlamentares?legislatura=&ano=2021&mes=&por=deputado&deputado=%s&uf=&partido=", id)

			response, err := http.Get(url)

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

			var deputado domain.Deputado
			deputado.Nome = nome

			cota, err := d.pegaCota(*doc)
			if err != nil {
				log.Fatal(err, "deputado: ", nome)
			}

			deputado.Cota = cota

			gabineteGasto, err := d.pegaVerbaDeGabineteGasto(*doc)
			if err != nil {
				log.Fatal(err, "deputado: ", nome)
			}

			deputado.VerbaDeGabineteGasto = gabineteGasto.verbaGasta
			deputado.PorcentagemGasto = gabineteGasto.porcentagemGasta

			gabineteDisponivel, err := d.pegaVerbaDeGabineteDisponivel(*doc)
			if err != nil {
				log.Fatal(err)
			}

			deputado.VerbaDeGabineteDisponivel = gabineteDisponivel.verbaDisponivel
			deputado.PorcentagemDisponivel = gabineteDisponivel.porcentagemDisponivel

			deputado.Partido = partido
			deputado.Estado = estado

			marshal, err := json.MarshalIndent(deputado, "", "")
			if err != nil {
				log.Fatalln(err.Error())
			}

			fmt.Printf("%s\n", marshal)

		}(dep.ID, dep.Nome, dep.Partido, dep.Estado)

	}

	wg.Wait()

}

func (d DeputadoRepository) BuscaDeputados() []domain.DeputadoResponse {
	var deputados []domain.DeputadoResponse

	response, err := http.Get("https://www.camara.leg.br/transparencia/gastos-parlamentares")
	if err != nil {
		fmt.Printf("FALHA AO EXECUTAR REQUISICAO %d %s",
			response.StatusCode, response.Status)
		panic(err.Error())
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#deputado").Each(func(i int, s *goquery.Selection) {
		s.Find("option").Each(func(i int, selection *goquery.Selection) {
			if len(selection.AttrOr("value", "")) != 0 {

				rgx := regexp.MustCompile("\\S+\\s\\S+")
				nome := rgx.FindString(selection.Text())

				rgx = regexp.MustCompile("\\(([^)]+)\\)")
				submatch := rgx.FindStringSubmatch(selection.Text())
				partidoEstado := submatch[1]
				partido := partidoEstado[0:2]
				estado := partidoEstado[len(partidoEstado)-2:]

				deputados = []domain.DeputadoResponse{
					{
						Nome: nome,
						Partido: partido,
						Estado: estado,
						ID: selection.AttrOr("value", "")},
				}
			}
		})
	})

	return deputados

}

func (d DeputadoRepository) pegaNome(document goquery.Document) (string, error) {
	nome := document.Find("#main-content > section.gastos-form > div.gastos-form__resumo-resposta > div > p > span:nth-child(1)").Text()

	if len(nome) == 0 {
		return "", errors.New("Nome não encontrado")
	}

	return nome, nil
}

func (d DeputadoRepository) pegaCota(document goquery.Document) (string, error) {
	cota := document.Find("#cota > div > div.l-cota__row > div:nth-child(1) > div > div.l-card.l-cota-resumo > div > div > section > p.gastos__resumo-texto.gastos__resumo-texto--destaque > span").Text()

	if len(cota) == 0 {
		return "", errors.New("Cota não encontrada")
	}

	return cota, nil
}

func (d DeputadoRepository) pegaVerbaDeGabineteGasto(document goquery.Document) (struct {
	verbaGasta       string
	porcentagemGasta string
}, error) {
	verbaGasta := document.Find("#js-percentual-gasto > tbody > tr:nth-child(1) > td:nth-child(2)").Text()

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
	}{verbaGasta, porcentagemGasta}, nil
}

func (d DeputadoRepository) pegaVerbaDeGabineteDisponivel(document goquery.Document) (struct {
	verbaDisponivel       string
	porcentagemDisponivel string
}, error) {
	verbaDisponivel := document.Find("#js-percentual-gasto > tbody > tr:nth-child(2) > td:nth-child(2)").Text()

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
