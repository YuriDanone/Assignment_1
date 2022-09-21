package main

import "fmt"

type Observable interface {
	subscribe()
	unsubscribe()
	sendAll()
}

type Observer interface {
	handleEvent(v []string)
}

type JobSite struct {
	subscribers []Person
	vacancies   []string
}

type Person struct {
	name string
}

func (p *Person) handleEvent(v []string) {
	fmt.Printf("Hi dear %v\n", p.name)
	fmt.Println("Vacancies updated:")
	for _, value := range v {
		fmt.Printf("%v\n", value)
	}
	fmt.Println("\n")
}

func (js *JobSite) sendAll() {
	for _, value := range js.subscribers {
		value.handleEvent(js.vacancies)
	}
}

func (js *JobSite) addVacancy(vacancy string) {
	js.vacancies = append(js.vacancies, vacancy)
	js.sendAll()
}

func (js *JobSite) removeVacancy(vacancy string) {
	tempVacancies := make([]string, 0)
	for _, value := range js.vacancies {
		if value != vacancy {
			tempVacancies = append(tempVacancies, value)
		}
	}
	js.vacancies = tempVacancies
	js.sendAll()
}

func (js *JobSite) subscribe(p Person) {
	js.subscribers = append(js.subscribers, p)

}

func (js *JobSite) unsubscribe(p Person) {
	tempSubscribers := make([]Person, 0)
	for _, value := range js.subscribers {
		if value != p {
			tempSubscribers = append(tempSubscribers, value)
		}
	}
	js.subscribers = tempSubscribers

}

func main() {
	hhKz := JobSite{
		subscribers: []Person{},
		vacancies:   []string{},
	}
	person1 := Person{"Dastan"}
	person2 := Person{"Adil"}
	hhKz.subscribe(person1)
	hhKz.subscribe(person2)
	hhKz.addVacancy("PHP Developer")
	hhKz.addVacancy("Node.js Developer")
}
