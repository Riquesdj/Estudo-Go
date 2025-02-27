// TESTE DE UNIDADE
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

func TestTipoDeEndereco(t *testing.T) {
	//Test+NomeDaFuncao(t *testing.T)

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua dos anjos", "Rua"},
		{"Rodovia do imigrantes", "Rodovia"},
		//{"Praça das rosas", "Tipo Inválido"},
		{"Estrada do campo limpo", "Estrada"},
		{"RUA dos cajus", "Rua"},
		//{" ", "Tipo Inválido"},
		{"AVENIDA REBOUÇAS", "Avenida"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.tipoEsperado {
			t.Errorf("O tipo de recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.tipoEsperado)
		}
	}

	/*
		enderecoParaTeste := "Rua Paulista"
		tipoDeEnderecoEsperado := "Avenida"
		tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	*/
	/*
		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			t.Errorf("O Tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				tipoDeEnderecoEsperado,
				tipoDeEnderecoRecebido)

		}*/

}
