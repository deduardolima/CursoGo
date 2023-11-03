package main


type Conta struct {
	saldo int
}



func (c *Conta) Simulador(valor int) int{
		c.saldo += valor
		println(c.saldo)
		return c.saldo
}

func main() {
		conta := Conta{saldo:100}
		conta.Simulador(200)
		println(conta.saldo);
	
}
