package domain

type Deputy struct {
	ID                        string `json:"-"`
	Nome                      string `json:"nome,omitempty"`
	Partido                   string `json:"partido,omitempty"`
	Estado                    string `json:"estado,omitempty"`
	Cota                      string `json:"cota,omitempty"`
	VerbaDeGabineteDisponivel string `json:"verba_de_gabinete_disponivel,omitempty"`
	PorcentagemDisponivel     string `json:"porcentagem_disponivel,omitempty"`
	VerbaDeGabineteGasto      string `json:"verba_de_gabinete_gasto,omitempty"`
	PorcentagemGasto          string `json:"porcentagem_gasto,omitempty"`
}

type DeputadoResponse struct {
	Nome    string
	Partido string
	Estado  string
	ID      string
}