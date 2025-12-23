package main

import (
	"fmt"
)

func main() {
	mapField := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	var noliki bool
	var num string
	fmt.Println("Игра Крестики-Нолики")
	fmt.Println("Начинают крестики, пишите в консоль цифру ячейки, которую хотите занять. Удачи!")
	for i := 1; i <= 9; i++ {
		switch i % 2 {
		case 0:
			noliki = true
			fmt.Println("Ходят нолики")
			printMap(mapField)
			fmt.Println("Введите цифру вашего хода:")
			num = getPlayerCellNumber(mapField)
			mapField = makeMove(mapField, num, noliki)
			if hasWinner(mapField) {
				printMap(mapField)
				fmt.Println("Ого! Нолики победили!")
				return
			}
		default:
			noliki = false
			fmt.Println("Ходят крестики")
			printMap(mapField)
			fmt.Println("Введите цифру вашего хода:")
			num = getPlayerCellNumber(mapField)
			mapField = makeMove(mapField, num, noliki)
			if hasWinner(mapField) {
				printMap(mapField)
				fmt.Println("Ура! Крестики победили!")
				return
			}
		}
	}
	if !hasWinner(mapField) {
		printMap(mapField)
		fmt.Print("Ничья! Победила дружба!")
	}
}

func printMap(pole [3][3]string) {
	const (
		Reset = "\033[0m"
		Red   = "\033[31m"
		Blue  = "\033[34m"
	)
	fmt.Println("-------------")
	for i := range pole {
		fmt.Print("| ")
		for j := range pole {
			switch pole[i][j] {
			case "O":
				fmt.Print(Blue, pole[i][j], Reset, " | ")
			case "X":
				fmt.Print(Red, pole[i][j], Reset, " | ")
			default:
				fmt.Print(Reset, pole[i][j], " | ")
			}
		}
		fmt.Println()
		fmt.Println("-------------")
	}
}

func isMoveCorrect(mapField [3][3]string, cellNumber string) bool {
	for i := range mapField {
		for j := range mapField {
			if mapField[i][j] == cellNumber {
				return true
			}
		}
	}
	return false
}

func getPlayerCellNumber(mapField [3][3]string) string {
	var a string
	validDigits := map[string]bool{
		"1": true, "2": true, "3": true,
		"4": true, "5": true, "6": true,
		"7": true, "8": true, "9": true,
	}

	for {
		fmt.Scan(&a)

		if !validDigits[a] {
			fmt.Println("Неверный ввод. Пожалуйста, введите цифру от 1 до 9.")
			continue
		}

		if !isMoveCorrect(mapField, a) {
			fmt.Println("Неверный ввод. Пожалуйста, введите цифру пустой ячейки.")
			continue
		}

		break
	}
	return a
}

func makeMove(mapField [3][3]string, cellNumber string, isZeroNow bool) [3][3]string {
	for i := range mapField {
		for j := range mapField {
			if mapField[i][j] == cellNumber {
				switch isZeroNow {
				case true:
					mapField[i][j] = "O"
				case false:
					mapField[i][j] = "X"
				}
			}
		}
	}
	return mapField
}

func hasWinner(mapField [3][3]string) bool {
	if mapField[0][0] == mapField[0][1] && mapField[0][0] == mapField[0][2] ||
		mapField[1][0] == mapField[1][1] && mapField[1][0] == mapField[1][2] ||
		mapField[2][0] == mapField[2][1] && mapField[2][0] == mapField[2][2] ||
		mapField[0][0] == mapField[1][0] && mapField[0][0] == mapField[2][0] ||
		mapField[0][1] == mapField[1][1] && mapField[0][1] == mapField[2][1] ||
		mapField[0][2] == mapField[1][2] && mapField[0][2] == mapField[2][2] ||
		mapField[0][0] == mapField[1][1] && mapField[0][0] == mapField[2][2] ||
		mapField[0][2] == mapField[1][1] && mapField[0][2] == mapField[2][0] {
		return true
	}
	return false
}
