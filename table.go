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
	counter := 1
	for {
		fmt.Printf("Название %v столбца: ", counter)
		counter++
		collumn, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		collumn = strings.TrimSpace(collumn)
		if collumn == "0" {
			break
		}
		if collumn == "" {
			continue
		}
		collumns = append(collumns, collumn)
	}
	counter = 1
	for {
		fmt.Print("\n")
		data := make(map[string]string)
		var inputData string
		for i := 0; i < len(collumns); i++ {
			collumn := strings.TrimSpace(collumns[i])
			var err error
			fmt.Printf("[%v строка] %q = ", counter, collumn)
			inputData, err = inputReader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			inputData = strings.TrimSpace(inputData)
			if inputData == "0" {
				break
			}
			data[collumn] = inputData

			if utf8.RuneCountInString(data[collumn]) > utf8.RuneCountInString(collumns[i]) {
				collumns[i] += strings.Repeat(" ", utf8.RuneCountInString(data[collumn])-utf8.RuneCountInString(collumns[i]))
			}
		}
		if inputData == "0" {
			break
		}
		tableData = append(tableData, data)
		counter++

	}
	fmt.Println("")
	fmt.Println("Ваша таблица: ")

	var line string
	for i := 0; i < len(collumns); i++ {
		if i+1 == len(collumns) {
			line += fmt.Sprintf("+%s+", strings.Repeat("-", utf8.RuneCountInString(collumns[i])+2))
		} else {
			line += fmt.Sprintf("+%s", strings.Repeat("-", utf8.RuneCountInString(collumns[i])+2))
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
			if utf8.RuneCountInString(tableData[i][collumn]) < utf8.RuneCountInString(collumns[j]) {
				tableData[i][collumn] += strings.Repeat(" ", utf8.RuneCountInString(collumns[j])-utf8.RuneCountInString(tableData[i][collumn]))
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
