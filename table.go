package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var tableData = []map[string]string{}
	var collumns = []string{}
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите колонку: ")
		collumn, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		collumn = strings.TrimSpace(collumn)
		if collumn == "0" {
			break
		}
		collumns = append(collumns, collumn)
	}

	for {
		fmt.Print("\n\nДобавление записи в таблицу\n\n")
		data := make(map[string]string)
		for i := 0; i < len(collumns); i++ {
			collumn := strings.TrimSpace(collumns[i])
			fmt.Printf("Введите %s: ", collumn)
			inputData, err := inputReader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			inputData = strings.TrimSpace(inputData)
			if inputData == "0" {
				break
			}
			data[collumn] = inputData

			if utf8.RuneCountInString(data[collumn]) > len(collumns[i]) {
				collumns[i] += strings.Repeat(" ", utf8.RuneCountInString(data[collumn])-len(collumns[i]))
			}
		}
		tableData = append(tableData, data)
		fmt.Print("Добавить еще запись в таблицу?: ")

		inputData, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		inputData = strings.TrimSpace(inputData)
		if inputData == "0" {
			break
		}
	}

	fmt.Println("")
	fmt.Println("Ваша таблица: ")

	var line string
	for i := 0; i < len(collumns); i++ {
		if i+1 == len(collumns) {
			line += fmt.Sprintf("+%s+", strings.Repeat("-", len(collumns[i])+2))
		} else {
			line += fmt.Sprintf("+%s", strings.Repeat("-", len(collumns[i])+2))
		}
	}

	fmt.Println(line)
	for i := 0; i < len(collumns); i++ {
		if i+1 == len(collumns) {
			fmt.Printf("| %s |\n", collumns[i])
		} else {
			fmt.Printf("| %s ", collumns[i])
		}
	}
	fmt.Println(line)

	for i := 0; i < len(tableData); i++ {
		tr := ""
		for j := 0; j < len(collumns); j++ {
			collumn := strings.TrimSpace(collumns[j])
			if utf8.RuneCountInString(tableData[i][collumn]) < len(collumns[j]) {
				tableData[i][collumn] += strings.Repeat(" ", len(collumns[j])-utf8.RuneCountInString(tableData[i][collumn]))
			}

			if j+1 == len(collumns) {
				tr += fmt.Sprintf("| %s |\n", tableData[i][collumn])
			} else {
				tr += fmt.Sprintf("| %s ", tableData[i][collumn])
			}
		}
		fmt.Print(tr)
		fmt.Println(line)
	}
}
