package domain

type Deputado struct {
	Nome                       string `json:"nome,omitempty"`
	Partido                    string `json:"partido,omitempty"`
	Estado                     string `json:"estado,omitempty"`
	Cota                       string `json:"cota,omitempty"`
	VerbaDeGabineteDisponoivel string `json:"verba_de_gabinete_disponoivel,omitempty"`
	PorcentagemDisponoivel     string `json:"porcentagem_disponoivel,omitempty"`
	VerbaDeGabineteGasto       string `json:"verba_de_gabinete_gasto,omitempty"`
	PorcentagemGasto           string `json:"porcentagem_gasto,omitempty"`
}
