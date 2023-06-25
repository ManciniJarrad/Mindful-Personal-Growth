package main

import (
	"fmt"
)

//Define a person type
type Person struct {
	name string
	age int
	career string
	goals []string
}

//Create a person
func CreatePerson(name string, age int, career string, goals []string) Person {
	return Person{
		name: name,
		age: age,
		career: career,
		goals: goals,
	}
}

//Get Goals Method 
func (p *Person) GetGoals() []string {
	return p.goals
}

//Set Goals Method
func (p *Person) SetGoals(goals []string) {
	p.goals = goals
}

//Create a Service Struct
type Service struct {
	people map[string]Person
	goals map[string][]string
}

//Create Service
func CreateService() Service {
	return Service{
		people: make(map[string]Person),
		goals: make(map[string][]string),
	}
}

//Add Person Method
func (s *Service) AddPerson(name string, age int, career string, goals []string) {
	if _, ok := s.people[name]; ok {
		fmt.Println("Person already exists")
		return
	}
	s.people[name] = CreatePerson(name, age, career, goals)
	s.goals[name] = goals
}

//Get Person Goals Method
func (s *Service) GetPersonGoals(name string) {
	if p, ok := s.people[name]; ok {
		fmt.Println(p.GetGoals())
		return
	}
	fmt.Println("Person doesn't exist")
}

//Set Person Goals Method
func (s *Service) SetPersonGoals(name string, goals []string) {
	if p, ok := s.people[name]; ok {
		p.SetGoals(goals)
		s.goals[name] = goals
		return
	}
	fmt.Println("Person doesn't exist")

}

//Provide Coaching Services Method
func (s *Service) ProvideCoachingServices(name string) {
	if p, ok := s.people[name]; ok {
		fmt.Printf("%s's goals: \n", name)
		for _, goal := range p.GetGoals() {
			fmt.Println("-", goal)
		}
		fmt.Println("Personal Growth and Development Coaching Services:")
		fmt.Println("- Understanding of shared vision and goals.")
		fmt.Println("- Exposure to changes in the industry.")
		fmt.Println("- Improved interpersonal communication.")
		fmt.Println("- Improved problem solving and decision making skills.")
		fmt.Println("- Increased self-awareness and confidence.")
		fmt.Println("- Enhanced productivity and work-life balance.")
		fmt.Println("- Improved resilience to bounce back from tough situations.")
		return
	}
	fmt.Println("Person doesn't exist")
}

func main() {
	service := CreateService()
	service.AddPerson("John", 30, "Software Engineer", []string{"Grow to a Senior Level", "Launch Side Project"})
	service.ProvideCoachingServices("John")
}