package main

import "fmt"

type Class_info struct {
	Kind            string
	Name_Object     string
	Surname_Teacher string
}

func Counter_practik(list []Class_info, Surname string) int {
	counter := 0

	for i := 0; i < len(list); i++ {
		el := list[i]
		if el.Kind == "Практика" && el.Surname_Teacher == Surname {
			counter++
		}
	}
	return counter
}

func main() {
	var list []Class_info
	Surname := "Гончарук"

	list = append(list, Class_info{"Парктика", "Матан", "Зиновьев"})
	list = append(list, Class_info{"Лекция", "Матан", "Зиновьев"})
	list = append(list, Class_info{"Практика", "Матан", "Зиновьев"})
	list = append(list, Class_info{"Практика", "Матан", "Осипова"})
	list = append(list, Class_info{"Лекция", "Матан", "Осипова"})
	list = append(list, Class_info{"Лекция", "ОАИП", "Дудко"})
	list = append(list, Class_info{"Практика", "ОАИП", "Дудко"})
	list = append(list, Class_info{"Практика", "ОАИП", "Гончарук"})
	list = append(list, Class_info{"Практика", "ОАИП", "Гончарук"})
	list = append(list, Class_info{"Практика", "ОАИП", "Гончарук"})
	list = append(list, Class_info{"Практика", "Матан", "Соболева"})
	list = append(list, Class_info{"Практика", "Ангем", "Сухонос"})

	fmt.Println(Counter_practik(list, Surname))

}
