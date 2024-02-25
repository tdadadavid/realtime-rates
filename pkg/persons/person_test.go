package persons_test

import (
	"realtime-exchange-rates/pkg/persons"
	"testing"

	"github.com/stretchr/testify/assert"
)

var asc = "asc"
var desc = "desc"
var filePath = "../../person.json"


func TestPersonIsSortedInAsc(t *testing.T) {
	test_persons, err := persons.GetPersons(filePath)
	assert.Nil(t, err)

	t.Run("Persons are returned in ascending order", func(t *testing.T) {
		sortedPersonsResult := test_persons.Sort(asc)
		assert.NotNil(t, sortedPersonsResult.Data)

		firstPerson := sortedPersonsResult.Data[0]
		lastPerson := sortedPersonsResult.Data[len(sortedPersonsResult.Data)-1]
		assert.True(t,lastPerson.Salary.Value > firstPerson.Salary.Value)
	})
}

func TestPersonIsSortedInDesc(t *testing.T) {
	test_persons, err := persons.GetPersons(filePath)
	assert.Nil(t, err)

	t.Run("Persons are returned in ascending order", func(t *testing.T) {
		sortedPersonsResult := test_persons.Sort(desc)
		assert.NotNil(t, sortedPersonsResult.Data)

		firstPerson := sortedPersonsResult.Data[0]
		lastPerson := sortedPersonsResult.Data[len(sortedPersonsResult.Data)-1]
		assert.True(t,lastPerson.Salary.Value < firstPerson.Salary.Value)
	})
}

func TestPersonSortingDefaultsToDescIfNotSpecifiedCorrectly(t *testing.T) {
	test_persons, err := persons.GetPersons(filePath)
	assert.Nil(t, err)

	t.Run("Persons are returned in ascending order", func(t *testing.T) {
		sortedPersonsResult := test_persons.Sort("wrong-order")
		assert.NotNil(t, sortedPersonsResult.Data)

		firstPerson := sortedPersonsResult.Data[0]
		lastPerson := sortedPersonsResult.Data[len(sortedPersonsResult.Data)-1]
		assert.True(t,lastPerson.Salary.Value < firstPerson.Salary.Value)
	})
}