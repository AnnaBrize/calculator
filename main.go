package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var char1, char2, out1, argA, argB int
var lang, okA, okB bool
var argBin, argAin, operin, oper, out2 string

// римские цифры карта
var roma = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	20: "XX", 30: "XXX", 40: "XL", 50: "L",
	60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C"}

// переводим с римского на арабский
func intRoma(m map[int]string, val string) (key int, ok bool) {
	for k, v := range m {
		if v == val {
			key = k
			ok = true
			return
		}
		key = 0
		ok = false
	}
	return
}

// переводим с арабского на римский
func convRoma(numb int) (outNumb string) {
	if numb == 100 {
		outNumb = roma[100]
	} else if numb > 9 {
		ch0 := numb / 10
		ch1 := ch0 * 10
		ch2 := numb % 10
		outNumb = roma[ch1] + roma[ch2]
	} else {
		outNumb = roma[numb]
	}
	return
}

func main() {
	//вводим данные с клавиатуры
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("введити математический пример:")
	text, _ := reader.ReadString('\n')

	t := strings.Split(text, " ")

	if len(t) > 3 || len(t) < 3 {
		error := fmt.Errorf("ERROR: формат математической операции не удовлетворяет заданию\n— два операнда и один оператор (+, -, /, *),\nили строка не является математической операцией.")
		fmt.Println(error)
		return
	} else {
		argAin = strings.TrimSpace(t[0])
		argBin = strings.TrimSpace(t[2])
		operin = strings.TrimSpace(t[1])
	}

	// проверяем систему счисления
	char1, okA := intRoma(roma, argAin)
	char2, okB := intRoma(roma, argBin)

	if okA == true {
		if okB == true {
			lang = false
			argA = char1
			argB = char2
		} else {
			error := fmt.Errorf("ERROR: используются одновременно разные системы счисления.")
			fmt.Println(error)
			return
		}
	} else {
		if okB == true {
			error := fmt.Errorf("ERROR: используются одновременно разные системы счисления.")
			fmt.Println(error)
			return
		} else {
			lang = true
			argA, _ = strconv.Atoi(argAin)
			argB, _ = strconv.Atoi(argBin)
		}
	}

	if argA > 10 || argB > 10 {
		error := fmt.Errorf("ERROR: формат математической операции не удовлетворяет заданию\n-Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более")
		fmt.Println(error)
		return
	}

	//проверяем операцию вычисления
	oper := operin

	switch {
	case oper == "+":
		out1 = argA + argB
	case oper == "-":
		out1 = argA - argB
	case oper == "*":
		out1 = argA * argB
	case oper == "/":
		out1 = argA / argB
	default:
		error := fmt.Errorf("ERROR: такой операции не существует")
		fmt.Println(error)
		return
	}

	//ковертируем в систему счисления и выводим результат
	if lang == true {
		fmt.Println("=", out1)
		return
	} else {
		if out1 <= 0 {
			error := fmt.Errorf("ERROR: в римской системе нет отрицательных чисел и 0")
			fmt.Println(error)
			return
		}
		out2 := convRoma(out1)
		fmt.Println("=", out2)
		return
	}
}
