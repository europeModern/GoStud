package main

import "fmt"

type PepeSchnele struct {
	Speed int
	Charisma int
	Wisdom int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed: speed, Charisma: charisma, Wisdom: wisdom,
	}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func (p *PepeSchnele) String() string {
	return fmt.Sprintf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
		p.Speed, p.Charisma, p.Wisdom, p.GetRating())
}

func main() {
	fmt.Println("ЗАДАЧА #1: Структура Пепе")
	
	pepe1 := NewPepeSchnele(10, 8, 9)
	pepe2 := NewPepeSchnele(7, 10, 12)
	
	fmt.Println(pepe1)
	fmt.Println(pepe2)
	
	rating1 := pepe1.GetRating()
	rating2 := pepe2.GetRating()
	
	if rating1 > rating2 {
		fmt.Printf("\nПервый Пепе Шнеле имеет более высокий рейтинг (%d > %d)\n", rating1, rating2)
	} else if rating2 > rating1 {
		fmt.Printf("\nВторой Пепе Шнеле имеет более высокий рейтинг (%d > %d)\n", rating2, rating1)
	} else {
		fmt.Printf("\nОба Пепе Шнеле имеют одинаковый рейтинг (%d)\n", rating1)
	}
}
