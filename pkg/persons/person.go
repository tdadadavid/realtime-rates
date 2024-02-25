package persons

import (
	"encoding/json"
	"fmt"
	"realtime-exchange-rates/utils"
)

type Salary struct {
	Value float64 `json:"value"`
	Currency string `json:"currency"`
}

type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Salary Salary `json:"salary"`
}

type GroupByCurrencyResult map[string][]Person

type Persons struct {
	Data []Person `json:"data"`
}


func (persons *Persons) Sort(direction string) Persons {

	var sortedPersons []Person
	if direction == "asc" {
		sortedPersons = persons.sortAsc()
	}else{
	  sortedPersons =	persons.sortDesc()
	}
	
	return Persons {
		Data: sortedPersons,
	}
}

func (persons* Persons) sortAsc() []Person {
	return []Person {
		
	}
}

func (persons *Persons) sortDesc() []Person {
	return []Person {
		
	}
}

func (persons *Persons) GroupBySalary() GroupByCurrencyResult {
	
	return GroupByCurrencyResult{
		"USD": []Person{},
		"NGN": []Person{},
		"EUR": []Person{},
	}
}

func (persons *Persons) FilterBySalary(amount int64) []Person {
	return []Person {
		
	}
}

func GetPersons() (*Persons, error) {
	personsStrings, err := utils.ReadFromJSONFile("person.json")
	if err != nil {
		return nil, err
	}

	var JSONPerson map[string]Person
	err = json.Unmarshal([]byte(personsStrings), &JSONPerson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	personsList := castJSONStringToArrayOfPerson(JSONPerson)
	persons := Persons{Data: personsList}

	return &persons, nil
}

func castJSONStringToArrayOfPerson(jsonPerson map[string]Person) []Person {
	var personsList []Person
	for _, person := range jsonPerson {
		personsList = append(personsList, person)
	}
	return personsList
}