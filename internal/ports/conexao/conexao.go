package conexao

import "deputySpending/internal/domain"

type Conexao interface {

	BuscaDeputado(fn func() []domain.DeputadoResponse)
}
