package main

import "fmt"

func main() {
	cards := []string{"Test ne con", "Them mot test nua ne"}
	cards = append(cards, "Five of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
