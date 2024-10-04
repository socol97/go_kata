package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var number1_string string
	var number2_string string
	var number1 int
	var number2 int
	var symbol string
	var isRoman bool
	fmt.Println("Введите первое число")
	_, err := fmt.Scan(&number1_string)
	if err != nil {
		log.Fatal(fmt.Printf("scan error: %w", err))
	}
	value, exists := RomanNumbers[number1_string]
	if exists {
		number1 = value
		isRoman = true
	}
	fmt.Println("Введите второе число")
	_, err = fmt.Scan(&number2_string)
	if err != nil {
		log.Fatal(fmt.Sprintf("scan error: %w", err))
	}
	if isRoman {
		value, exists := RomanNumbers[number2_string]
		if exists {
			number2 = value
		} else {
			fmt.Println("Второе число тоже дожно являться римским")
		}
	}
	fmt.Println("Введите операцию(+,-,*,/)")
	_, err = fmt.Scan(&symbol)
	if err != nil {
		log.Fatal(fmt.Sprintf("scan error: %w", err))
	}
	if symbol != "*" && symbol != "+" && symbol != "-" && symbol != "/" {
		log.Fatal("Операция не найдена")
	}
	if !isRoman {
		number1, err = strconv.Atoi(number1_string)
		if err != nil {
			value, exists := RomanNumbers[number1_string]
			if exists {
				number1 = value
			} else {
				log.Fatal(fmt.Sprintf("scan error: %w", err))
			}
		}
	}
	if number1 >= 10 && number1 > 0 {
		log.Fatal("Число не в диапазоне 1-10")
	}
	if !isRoman {
		number2, err = strconv.Atoi(number2_string)
		if err != nil {
			value, exists := RomanNumbers[number1_string]
			if exists {
				number2 = value
			} else {
				log.Fatal(fmt.Sprintf("scan error: %w", err))
			}
		}
	}
	if number2 >= 10 && number2 > 0 {
		log.Fatal("Число не в диапазоне 1-10")
	}
	var result string
	switch symbol {
	case "+":
		if isRoman {
			if number1+number2 < 1 {
				log.Fatal("результат(Римские числа) меньше нуля...")
			} else {
				result = ReverseRomanNumbers[number1+number2]
			}
		} else {
			result = strconv.Itoa(number1 + number2)
		}
		fmt.Println(fmt.Sprintf("%s + %s = %s", number1_string, number2_string, result))
	case "-":
		if isRoman {
			if number1-number2 < 1 {
				log.Fatal("результат(Римские числа) меньше нуля...")
			} else {
				result = ReverseRomanNumbers[number1-number2]
			}
		} else {
			result = strconv.Itoa(number1 - number2)
		}
		fmt.Println(fmt.Sprintf("%s - %s = %s", number1_string, number2_string, result))
	case "/":
		if isRoman {
			if number1/number2 < 1 {
				log.Fatal("результат(Римские числа) меньше нуля...")
			} else {
				result = ReverseRomanNumbers[number1/number2]
			}
		} else {
			result = strconv.Itoa(number1 / number2)
		}
		fmt.Println(fmt.Sprintf("%s / %s = %s", number1_string, number2_string, result))
	case "*":
		if isRoman {
			if number1*number2 < 1 {
				log.Fatal("результат(Римские числа) меньше нуля...")
			} else {
				result = ReverseRomanNumbers[number1+number2]
			}
		} else {
			result = strconv.Itoa(number1 * number2)
		}
		fmt.Println(fmt.Sprintf("%s * %s = %s", number1_string, number2_string, result))
	default:
		fmt.Println("Неизвестная операция")
	}

}
