package matematica

//Calculo executa qualquer tipo de cálculo
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador é uma função que múltiplica x vezes o y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor efetua do número a pelo número b
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto efetua do número a pelo número b e mostra o resto da divisão
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
