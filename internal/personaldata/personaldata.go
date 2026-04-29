package personaldata

import "fmt"

type Personal struct {
	Name           string
	Weight, Height float64
}

func (p Personal) Print() {
	fmt.Printf(`Имя: %s
Вес: %.2f
Рост: %.2f
`, p.Name, p.Weight, p.Height)
}
