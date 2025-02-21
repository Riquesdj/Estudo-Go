package enderecos

import "strings"

// TipoDeEndereco  verifica se um endereço tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	// strings.split() vai separar todos os itens em um slice sempre que ele detectar o que está depois da virgula
	// Nesse caso seria (a cada " " na string endereco ele irá separar as palavras)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	//RUA DOS BOBOS COMPLEMENTO ALGUMA COISA
	//["RUA", "DOS", "BOBOS", "COMPLEMENTO", "ALGUMA", COISA]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo Inválido"
}
