package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var globalRom = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}                                                          // Словарь: римские цифры => арабские цифры
var orderRom = []string{"M", "D", "C", "L", "X", "V", "I"} // Порядок римских цифр для словаря

func main() {

	fmt.Println("Введите операцию:")
	reader := bufio.NewReader(os.Stdin)

	inputData, _ := reader.ReadString('\n')
	inputData = strings.TrimSpace(inputData)

	text := strings.Split(inputData, " ") // Получаем массив разделяя строку
	if len(text) < 3 {
		panic("Строка не является математической операцией")
	} else if len(text) > 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	} else if (strings.Count(text[0], "I") > 3 || strings.Count(text[2], "I") > 3) || (strings.Count(text[0], "V") > 1 || strings.Count(text[2], "V") > 1) {
		panic("Неверный формат введенного римского числа")
	}

	typeOperands := panicAlarm(text[0], text[2])

	if typeOperands == "int" { // Если оба число
		fmt.Println(operations(text))
	} else if typeOperands == "stringRom" { // Если оба римские цифры
		result := operationRom(text)
		fmt.Println(result)
	} else {
		panic("Неизвестное выражение")
	}
}

func operations(lst []string) int {

	panicAlarm(lst[0], lst[2])
	operand1, _ := strconv.Atoi(lst[0])
	operand2, _ := strconv.Atoi(lst[2])

	switch lst[1] {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "/":
		return operand1 / operand2
	case "*":
		return operand1 * operand2
	default:
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
} // Проводим операцию в зависимости от операнда

func panicAlarm(x string, y string) string {
	operandRom := func(x string) bool {
		for _, value := range orderRom {
			for _, result := range x {
				if string(result) == value {
					return false
				}
			}
		}
		return true
	}

	type1 := typeOperand(x)
	type2 := typeOperand(y)

	if type1 == "float64" || type2 == "float64" {
		panic("Калькулятор умеет работать только с целыми числами")
	} else if type1 == "int" && type2 == "int" {
		x, _ := strconv.Atoi(x)
		y, _ := strconv.Atoi(y)
		if x > 10 || y > 10 || x < 1 || y < 1 {
			panic("Калькулятор принимает на вход только числа от 1 до 10 включительно")
		}
		return "int"
	} else if (type1 == "string" && type2 == "int") || (type1 == "int" && type2 == "string") {
		panic("Используются одновременно разные системы счисления")
	} else if operandRom(x) || operandRom(y) {
		panic("Калькулятор умеет работать только с арабскими или римскими цифрами")
	} else if type1 == "string" && type2 == "string" {
		for _, value := range x {
			value := string(value)

			rt := globalRom[value]
			if rt == 0 {
				panic("Неверный формат введенного римского числа")
			}
		}
		for _, value := range y {
			value := string(value)

			rt := globalRom[value]
			if rt == 0 {
				panic("Неверный формат введенного римского числа")
			}
		}

		return "stringRom"
	}
	return ""
} // Проверка условий

func typeOperand(operand string) string {
	isInt := func(x string) bool {
		_, err := strconv.Atoi(x)
		return err == nil
	}

	isFloat := func(x string) bool {
		_, err := strconv.ParseFloat(x, 64)
		return err == nil
	}

	isFloatRom := func(x string) bool {
		for _, value := range x {
			if string(value) == "." {
				return true
			}
		}
		return false
	}

	if isInt(operand) {
		return "int"
	} else if isFloat(operand) {
		return "float64"
	} else if isFloatRom(operand) {
		return "float64"
	}
	return "string"

} // Получения типа оператора

func transformationRom(x int) string {
	str := ""
	if x == 0 { // Базовый случай
		return ""
	}

	for _, value := range orderRom {
		if globalRom[value]-1 == x {
			str += "I"
			str += value
			x -= globalRom[value] - 1
			break
		} else if x-globalRom[value] >= 0 {
			str += value
			x -= globalRom[value]
			break
		}
	}

	return str + transformationRom(x)
} // Рекурсия: арабские числа => римские числа

func transformationAr(x string) int {
	summ := 0
	for key, value := range x {
		value := string(value)
		if key < len(x)-1 {
			if globalRom[value] < globalRom[string(x[key+1])] {
				summ -= globalRom[value]
				continue
			} else if globalRom[value] > globalRom[string(x[key+1])] {
				summ += globalRom[value]
				continue
			} else {
				summ += globalRom[value]
			}
		} else {
			summ += globalRom[value]
		}
	}
	return summ
} // римские числа => арабские числа

func operationRom(lst []string) string {
	res1 := transformationAr(lst[0])
	res2 := transformationAr(lst[2])
	newLst := []string{strconv.Itoa(res1), lst[1], strconv.Itoa(res2)}

	resOperation := operations(newLst)
	if resOperation <= 0 {
		panic("Результат работы с римскими числами должен быть положительным числом")
	}
	return transformationRom(resOperation)
} // Операции над римскими цифрами
