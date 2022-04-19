package conexaoMockada

type ConexaoMockada struct {
}

func (conexaoHttp *ConexaoMockada) BuscaDeputado() string {

	return "Retornou o deputado"
}
