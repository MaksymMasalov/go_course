package main

/*
Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)
*/

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) (string, error) {
	var result strings.Builder
	runeSlice := []rune(s)
	length := len(runeSlice)
	i := 0

	for i < length {
		if unicode.IsDigit(runeSlice[i]) {
			return "", fmt.Errorf("некорректная строка")
		}

		result.WriteRune(runeSlice[i])
		i++

		if i < length && unicode.IsDigit(runeSlice[i]) {
			count, _ := strconv.Atoi(string(runeSlice[i]))
			result.WriteString(strings.Repeat(string(runeSlice[i-1]), count))
			i++
		}
	}

	return result.String(), nil
}

func main() {
	stringsToUnpack := []string{
		"a4bc2d5e",
		"abcd",
		"45",
	}

	for _, s := range stringsToUnpack {
		result, err := unpackString(s)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}
