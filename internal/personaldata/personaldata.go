package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес пользователя
	Height float64 // рост пользователя
}

// Print выводит данные структуры на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
