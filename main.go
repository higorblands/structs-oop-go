package main

import (
	"fmt"
	"math"
	"reflect"
)

// struct objetos complexos e semelhanças com OOP

// struct to student
type Student struct {
	name string
}

// struct to course
type Course struct {
	description string
	students    []Student
}

func (course *Course) register(student Student) {
	course.students = append(course.students, student)
}

// EadCourse This is another course type
type EadCourse struct {
	course  Course
	website string
}

// NewFloat float with steroids

type NewFloat float64

func (newFloat NewFloat) roundBy(places float64) NewFloat {
	newPlaces := math.Pow(10.00, places)
	return NewFloat(math.Round(float64(newFloat)*newPlaces) / newPlaces)

}

// json struct complex private
type jsonStruct struct {
	json []string
}

// json Type Public
type Json jsonStruct

func (json Json) toJson(text string) Json {
	newJson := Json{append(json.json, "{", text, "}")}
	return newJson
}

func main() {

	var student Student
	student.name = "Higor"
	fmt.Print(student.name, " - ", reflect.TypeOf(student), "\n")

	dotInMemory := &student
	fmt.Print(dotInMemory, " - ", reflect.TypeOf(dotInMemory), "\n")
	fmt.Print(dotInMemory.name, " - ", reflect.TypeOf(student), "\n")

	var otherStudent Student
	otherStudent.name = "Aldry"
	fmt.Print(dotInMemory.name, " - ", otherStudent.name, "\n")

	person := Student{"Goku"}
	fmt.Print(person.name, " - ", reflect.TypeOf(person), "\n")

	//Função
	engineering := Course{"Engineering", make([]Student, 0)}
	engineering.register(person)
	engineering.register(otherStudent)
	fmt.Println(engineering)

	//Herança
	ead := EadCourse{
		Course{"New EAD", make([]Student, 0)},
		"https://myead.com"}
	ead.course.register(student)
	fmt.Println(ead)

	// Métodos em outros tipos
	var number NewFloat = 5.0293019384
	fmt.Println(number.roundBy(2))
	fmt.Println(number.roundBy(3))

	// Criando toJson
	myJson := Json{make([]string, 0)}
	myJson = myJson.toJson("\"key\": \"value\"")
	myJson = myJson.toJson("\"key2\": \"value2\"")
	fmt.Println(myJson.json)
}