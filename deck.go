package main

import (
	"fmt"
	"math/rand"
	"time"
)

var actualRandNum, actualRandNum1 int

func main() {
	deck := Deck{
		cards: []Card{
			{"Туз", "Пики"},
			{"Двойка", "Пики"},
			{"Тройка", "Пики"},
			{"Четверка", "Пики"},
			{"Пятерка", "Пики"},
			{"Шестерка", "Пики"},
			{"Семерка", "Пики"},
			{"Восьмерка", "Пики"},
			{"Девятка", "Пики"},
			{"Десятка", "Пики"},
			{"Валет", "Пики"},
			{"Дама", "Пики"},
			{"Король", "Пики"},

			{"Туз", "Черви"},
			{"Двойка", "Черви"},
			{"Тройка", "Черви"},
			{"Четверка", "Черви"},
			{"Пятерка", "Черви"},
			{"Шестерка", "Черви"},
			{"Семерка", "Черви"},
			{"Восьмерка", "Черви"},
			{"Девятка", "Черви"},
			{"Десятка", "Черви"},
			{"Валет", "Черви"},
			{"Дама", "Черви"},
			{"Король", "Черви"},

			{"Туз", "Буби"},
			{"Двойка", "Буби"},
			{"Тройка", "Буби"},
			{"Четверка", "Буби"},
			{"Пятерка", "Буби"},
			{"Шестерка", "Буби"},
			{"Семерка", "Буби"},
			{"Восьмерка", "Буби"},
			{"Девятка", "Буби"},
			{"Десятка", "Буби"},
			{"Валет", "Буби"},
			{"Дама", "Буби"},
			{"Король", "Буби"},

			{"Туз", "Треф"},
			{"Двойка", "Треф"},
			{"Тройка", "Треф"},
			{"Четверка", "Треф"},
			{"Пятерка", "Треф"},
			{"Шестерка", "Треф"},
			{"Семерка", "Треф"},
			{"Восьмерка", "Треф"},
			{"Девятка", "Треф"},
			{"Десятка", "Треф"},
			{"Валет", "Треф"},
			{"Дама", "Треф"},
			{"Король", "Треф"},
		},
	}

	for i := 0; i < 80; i++ {
		deck.overhandShuffle()
	}
	for _, c := range deck.cards {
		fmt.Println(c)
	}
	fmt.Println(len(deck.cards))

}

type Card struct {
	value  string
	colour string
}

type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(len(d.cards) - 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) overhandShuffle() {
	rand.Seed(time.Now().UnixNano())
	actualRandNum = len(d.cards)/2 + rand.Intn(len(d.cards)/3) - len(d.cards)/6
	leftHand := d.cards[actualRandNum:len(d.cards)]
	rightHand := d.cards[0:actualRandNum]
	actualRandNum = rand.Intn(3) + 1
	actualRandNum1 = len(rightHand)
	for i := 0; i < actualRandNum && len(rightHand) > actualRandNum1-actualRandNum; i++ {
		leftHand = append(leftHand, rightHand[len(rightHand)-actualRandNum1/actualRandNum:]...)
		rightHand = append(rightHand[:len(rightHand)-actualRandNum1/actualRandNum], rightHand[len(rightHand)-1])
		rightHand = append(rightHand[:len(rightHand)-1], rightHand[len(rightHand):]...)
	}
	leftHand = append(leftHand, rightHand...)
	d.cards = leftHand
}
