package operadora

import "github.com/robsonkumagai/cursodego/variaveis/pacotes/prefixo"
import "strconv"

//NomeOperadora é uma variável pública
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora é uma variável de outro pacote
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + "" + NomeOperadora
