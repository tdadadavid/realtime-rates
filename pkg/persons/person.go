package persons

import (
	"encoding/json"
	"fmt"
	"realtime-exchange-rates/utils"
	"sort"
	"strings"
)

type Salary struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type Person struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Salary Salary `json:"salary"`
}

type GroupByCurrencyResult map[string][]Person

type Persons struct {
	Data []Person `json:"data"`
}

func (persons *Persons) Sort(direction string) Persons {

	var sortedPersons []Person
	if direction == "asc" {
		sortedPersons = persons.SortAsc()
	} else {
		sortedPersons = persons.SortDesc()
	}

	return Persons{
		Data: sortedPersons,
	}
}

func (persons *Persons) SortAsc() []Person {
	dataSet := persons.Duplicate()

	sort.Slice(dataSet, func(i, j int) bool {
		return dataSet[i].Salary.Value < dataSet[j].Salary.Value
	})
	return dataSet
}

func (persons *Persons) SortDesc() []Person {
	dataSet := persons.Duplicate()

	sort.Slice(dataSet, func(i, j int) bool {
		return dataSet[i].Salary.Value > dataSet[j].Salary.Value
	})
	return dataSet
}

func (persons *Persons) Duplicate() []Person {
	duplicate := make([]Person, len(persons.Data))
	copy(duplicate, persons.Data)
	return duplicate
}

func (persons *Persons) GroupByCurrency() GroupByCurrencyResult {

	var currencyToPersons = map[string][]Person{} //hashmap

	// for each person
	for _, person := range persons.Data {
		// get the current person's currency
		currentCurrency := strings.ToUpper(person.Salary.Currency)

		if currentCurrency == "" {
			currentCurrency = "NO-SALARY" // this edge case might never happen but if it does.
		}

		// check if the currency is already in the hashmap
		if value, ok := currencyToPersons[currentCurrency]; ok {
			value = append(value, person)              // add this person to the category
			currencyToPersons[currentCurrency] = value // update the map
		} else {
			currencyToPersons[currentCurrency] = []Person{person} // add the new category.
		}
	}

	return currencyToPersons
}

func (persons *Persons) FilterBySalary(amount int64) []Person {
	return []Person{}
}

func GetPersons(filePath string) (*Persons, error) {
	personsStrings, err := utils.ReadFromJSONFile(filePath)
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
