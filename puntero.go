package main

func main() {
	number := 25
	puntero := &number
	cambiarVariable(puntero)
	println(number)
}

func cambiarVariable(number *int) {
	*number = 32
}
