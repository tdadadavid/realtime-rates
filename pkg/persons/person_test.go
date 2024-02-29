package persons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var asc = "asc"
var desc = "desc"
var filePath = "../../person.json"

func TestPersonIsSortedInAsc(t *testing.T) {
	test_persons, err := GetPersons(filePath)
	assert.Nil(t, err)

	t.Run("Persons are returned in ascending order", func(t *testing.T) {
		sortedPersonsResult := test_persons.Sort(asc)
		assert.NotNil(t, sortedPersonsResult.Data)

		firstPerson := sortedPersonsResult.Data[0]
		lastPerson := sortedPersonsResult.Data[len(sortedPersonsResult.Data)-1]
		assert.True(t, lastPerson.Salary.Value > firstPerson.Salary.Value)
	})
}

func TestPersonIsSortedInDesc(t *testing.T) {
	test_persons, err := GetPersons(filePath)
	assert.Nil(t, err)

	t.Run("Persons are returned in descending order when stated", func(t *testing.T) {
		sortedPersonsResult := test_persons.Sort(desc)
		assert.NotNil(t, sortedPersonsResult.Data)

		firstPerson := sortedPersonsResult.Data[0]
		lastPerson := sortedPersonsResult.Data[len(sortedPersonsResult.Data)-1]
		assert.True(t, lastPerson.Salary.Value < firstPerson.Salary.Value)
	})
}

func TestPersonSortingDefaultsToDescIfNotSpecifiedCorrectly(t *testing.T) {
	test_persons, err := GetPersons(filePath)
	assert.Nil(t, err)

	t.Run("Persons are returned in descending order if wrong order was provided", func(t *testing.T) {
		sortedPersonsResult := test_persons.Sort("wrong-order")
		assert.NotNil(t, sortedPersonsResult.Data)

		firstPerson := sortedPersonsResult.Data[0]
		lastPerson := sortedPersonsResult.Data[len(sortedPersonsResult.Data)-1]
		assert.True(t, lastPerson.Salary.Value < firstPerson.Salary.Value)
	})
}

func TestPersonsGrouping(t *testing.T) {
	test_persons := Persons{
		Data: []Person{
			{
				Name: "test1",
				Id:   "1",
				Salary: Salary{
					Value:    30,
					Currency: "USD",
				},
			},
			{
				Name: "test2",
				Id:   "2",
				Salary: Salary{
					Value:    405,
					Currency: "NGN",
				},
			},
			{
				Name: "test2",
				Id:   "2",
				Salary: Salary{
					Value:    30,
					Currency: "usd",
				},
			},
			{
				Name:   "test2",
				Id:     "2",
				Salary: Salary{},
			},
		},
	}

	groupedPersonsResult := test_persons.GroupByCurrency()

	t.Run("It groups persons based on same currency despite letter-case difference", func(t *testing.T) {
		assert.NotNil(t, groupedPersonsResult)
		assert.True(t, len(groupedPersonsResult["NGN"]) == 1)
		assert.True(t, len(groupedPersonsResult["USD"]) == 2)
	})

	t.Run("People without salary(currency) should be put under 'NO-SALARY' group", func(t *testing.T) {
		assert.True(t, len(groupedPersonsResult["NO-SALARY"]) == 1)
	})
}

func TestFilterByUSDConstraint(t *testing.T) {
	test_persons := Persons{
		Data: []Person{
			{
				Name: "test1",
				Id:   "1",
				Salary: Salary{
					Value:    30,
					Currency: "USD",
				},
			},
			{
				Name: "David Dada",
				Id:   "1",
				Salary: Salary{
					Value:    3000,
					Currency: "USD",
				},
			},
			{
				Name: "test2",
				Id:   "2",
				Salary: Salary{
					Value:    405,
					Currency: "NGN",
				},
			},
			{
				Name: "test2",
				Id:   "2",
				Salary: Salary{
					Value:    30,
					Currency: "GBP", // greater than $100
				},
			},
			{
				Name: "test2",
				Id:   "2",
				Salary: Salary{
					Value:    300,
					Currency: "GBP", // greater than $100
				},
			},
			{
				Name: "test2",
				Id:   "2",
				Salary: Salary{
					Value:    0.12,
					Currency: "CAD",
				},
			},
		},
	}

	t.Run("People with less than $100 equivalent-salay are filtered out", func(t *testing.T) {
		people_with_salary_equal_or_greater_than_100, err := test_persons.FilterBySalary(100)
		assert.Nil(t, err)

		not_up_to_2_people_earn_equal_or_above_100_dollars := len(people_with_salary_equal_or_greater_than_100.Data) > 2

		assert.False(t, not_up_to_2_people_earn_equal_or_above_100_dollars)
	})
}
