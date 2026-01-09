// Конвертер температур: Цельсий → Фаренгейт.
// Формула: `F = C * 9/5 + 32`. Ввод из аргументов, вывод — с одним знаком после запятой.

// Пример:
// ```
// $ go run . 36.6
// 97.9 °F
// ```

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	c, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Ошибка: некорректное число: %s\n", os.Args[1])
		return
	}
	fmt.Printf("%.1f\n", c*9/5+32)

}