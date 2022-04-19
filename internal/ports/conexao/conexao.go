package conexao

type Conexao interface {
	BuscaDeputado(f func() []struct {
		nome    string
		partido string
		estado  string
		id      string
	}) string
}
