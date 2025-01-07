package main

import (
	"Tuan3/entity"
	"fmt"
)

func ShowListStudent() {
	student := entity.Student{
		ClassName: "DHKTPM18ATT",
		Point:     20.5,
		Person: entity.Person{
			PersonID: "ID93812",
			FullName: "Nguyen Anh Tung",
			Age:      20,
		},
	}
	fmt.Println(student)
}