package persons

import (
	"encoding/json"
	"fmt"
	exchangerates "realtime-exchange-rates/pkg/exchangeRates"
	"realtime-exchange-rates/utils"
	"sort"
	"strconv"
	"strings"
	"sync"
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

func (persons *Persons) FilterBySalary(amount float64) (Persons, error) {
	var wg sync.WaitGroup
	var qualifiedPersons []Person

	for _, person := range persons.Data {
		wg.Add(1)

		personsCurrency := strings.ToUpper(person.Salary.Currency)
		personsSalary := person.Salary.Value

		// capture people whose salaries are in 'usd' and
		// is greater than or equal to amount constrant
		if personsCurrency == "USD" && personsSalary >= amount {
			qualifiedPersons = append(qualifiedPersons, person)
			continue
		}

		// '%s-USD' becuase we are to find the value of the person salary in dollar.
		currencyPair := fmt.Sprintf("%s-USD", personsCurrency)

		result, err := exchangerates.GetExchangeRatesForCurrencyPair(currencyPair)
		if err != nil {
			return Persons{}, err
		}

		rate, err := strconv.ParseFloat(result.Rate, 64)
		if err != nil {
			return Persons{}, err
		}

		personSalaryInUSD := person.Salary.Value * rate

		if personSalaryInUSD >= float64(amount) {
			qualifiedPersons = append(qualifiedPersons, person)
		}
	}

	wg.Done()

	return Persons{
		Data: qualifiedPersons,
	}, nil
}

func GetPersons(filePath string) (*Persons, error) {
	personsStrings, err := utils.ReadFromJSONFile(filePath)
	if err != nil {
		return nil, err
	}

	var JSONPerson map[string]Person
	err = json.Unmarshal([]byte(personsStrings), &JSONPerson)
	if err != nil {
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
