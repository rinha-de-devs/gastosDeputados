package scrapping_adapter

import (
	"deputySpending/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type scrappingClient struct {
	deputies []domain.Deputy
}

func New() *scrappingClient {
	return &scrappingClient{
		deputies: []domain.Deputy{},
	}
}

func (scrapping *scrappingClient) SearchDeputySlice() ([]domain.Deputy, error) {

	response, err := http.Get("https://www.camara.leg.br/transparencia/gastos-parlamentares")
	if err != nil {
		fmt.Printf("FAIL TO PERFORM REQUEST %d %s",
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

				rgx := regexp.MustCompile(`\S+\s\S+`)
				nome := rgx.FindString(selection.Text())

				rgx = regexp.MustCompile(`\(([^)]+)\)`)
				submatch := rgx.FindStringSubmatch(selection.Text())
				partidoEstado := submatch[1]
				partido := partidoEstado[0:2]
				estado := partidoEstado[len(partidoEstado)-2:]

				tempDeputy := domain.Deputy{
					Nome:    nome,
					Partido: partido,
					Estado:  estado,
					ID:      selection.AttrOr("value", "")}

				scrapping.deputies = append(scrapping.deputies, tempDeputy)
			}
		})
	})

	return scrapping.deputies, nil
}

func (scrapping *scrappingClient) ScrappingDeputies(deputies []domain.Deputy) ([]domain.Deputy, error) {

	var wg sync.WaitGroup

	wg.Add(len(deputies))

	for index := range deputies {

		time.Sleep(1 * time.Second)
		go func(index int) {
			defer wg.Done()

			fmt.Printf("Progresso: %d de %d\n", index, len(deputies))

			url := fmt.Sprintf("https://www.camara.leg.br/transparencia/gastos-parlamentares?legislatura=&ano=2021&mes=&por=deputado&deputado=%s&uf=&partido=", deputies[index].ID)

			response, err := http.Get(url)
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

			cota, err := scrapping.pegaCota(*doc)
			if err != nil {
				log.Fatal(err, "deputado: ", deputies[index].Nome)
			}

			deputies[index].Cota = cota

			gabineteGasto, err := scrapping.pegaVerbaDeGabineteGasto(*doc)
			if err != nil {
				log.Fatal(err, "deputado: ", deputies[index].Nome)
			}

			deputies[index].VerbaDeGabineteGasto = gabineteGasto.verbaGasta
			deputies[index].PorcentagemGasto = gabineteGasto.porcentagemGasta

			gabineteDisponivel, err := scrapping.pegaVerbaDeGabineteDisponivel(*doc)
			if err != nil {
				log.Fatal(err)
			}

			deputies[index].VerbaDeGabineteDisponivel = gabineteDisponivel.verbaDisponivel
			deputies[index].PorcentagemDisponivel = gabineteDisponivel.porcentagemDisponivel

			marshal, err := json.MarshalIndent(deputies[index], "", "")
			if err != nil {
				log.Fatalln(err.Error())
			}

			fmt.Printf("%s\n", marshal)

		}(index)

	}

	wg.Wait()

	return deputies, nil
}

func (scrapping *scrappingClient) pegaCota(document goquery.Document) (string, error) {
	cota := document.Find("#cota > div > div.l-cota__row > div:nth-child(1) > div > div.l-card.l-cota-resumo > div > div > section > p.gastos__resumo-texto.gastos__resumo-texto--destaque > span").Text()

	if len(cota) == 0 {
		cota = "0"
	}

	return cota, nil
}

func (scrapping *scrappingClient) pegaVerbaDeGabineteGasto(document goquery.Document) (struct {
	verbaGasta       string
	porcentagemGasta string
}, error) {
	verbaGasta := document.Find("#js-percentual-gasto > tbody > tr:nth-child(1) > td:nth-child(2)").Text()

	if len(verbaGasta) == 0 {
		verbaGasta = "0"
	}

	porcentagemGasta := document.Find("#js-percentual-gasto > tbody > tr:nth-child(1) > td:nth-child(3)").Text()

	if len(porcentagemGasta) == 0 {
		porcentagemGasta = "0"
	}

	return struct {
		verbaGasta       string
		porcentagemGasta string
	}{verbaGasta, porcentagemGasta}, nil
}

func (scrapping *scrappingClient) pegaVerbaDeGabineteDisponivel(document goquery.Document) (struct {
	verbaDisponivel       string
	porcentagemDisponivel string
}, error) {
	verbaDisponivel := document.Find("#js-percentual-gasto > tbody > tr:nth-child(2) > td:nth-child(2)").Text()

	if len(verbaDisponivel) == 0 {
		verbaDisponivel = "0"
	}

	porcentagemDisponivel := document.Find("#js-percentual-gasto > tbody > tr:nth-child(2) > td:nth-child(3)").Text()

	if len(porcentagemDisponivel) == 0 {
		porcentagemDisponivel = "0"
	}

	return struct {
		verbaDisponivel       string
		porcentagemDisponivel string
	}{verbaDisponivel, porcentagemDisponivel}, nil
}
