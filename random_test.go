package jabufaker

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	randInt := RandomInt(0, 100)

	assert.NotEmpty(t, randInt)                                // Not empty data
	assert.Equal(t, "int64", reflect.TypeOf(randInt).String()) // Type must be `int64`
}

func TestRandomString(t *testing.T) {
	randString := RandomString(10)

	assert.NotEmpty(t, randString)                                 // Not Empty data
	assert.Equal(t, "string", reflect.TypeOf(randString).String()) // Type must be `string`
}

func TestRandomPerson(t *testing.T) {
	randPerson := RandomPerson()

	assert.NotEmpty(t, randPerson)                                 // Not Empty data
	assert.Equal(t, "string", reflect.TypeOf(randPerson).String()) // Type must be `string`
}

func TestRandomMoney(t *testing.T) {
	randMoney := RandomMoney()

	assert.NotEmpty(t, randMoney)                                // Not Empty data
	assert.Equal(t, "int64", reflect.TypeOf(randMoney).String()) // Type must be `int64`
}

func TestRandomCurrency(t *testing.T) {
	randCurrencies := RandomCurrency()

	assert.NotEmpty(t, randCurrencies)                                 // Not Empty data
	assert.Equal(t, "string", reflect.TypeOf(randCurrencies).String()) // Type must be `string`
}

func TestRandomEmail(t *testing.T) {
	randEmail := RandomEmail()
	removeDomain := strings.ReplaceAll(randEmail, "@test.com", "")
	containDomain := strings.Contains(randEmail, "@test.com")
	fmt.Println(randEmail)
	assert.NotEmpty(t, removeDomain)                              // not empty before @test.com
	assert.True(t, containDomain)                                 // must be contain @test.com
	assert.Equal(t, "string", reflect.TypeOf(randEmail).String()) // Type must be `string`

	// must be lower case
	isLower := IsLower(randEmail)
	assert.True(t, isLower)
}

func TestRandomProvince(t *testing.T) {
	randProvince := RandomProvince()

	assert.NotEmpty(t, randProvince)                                 // Not Empty data
	assert.Equal(t, "string", reflect.TypeOf(randProvince).String()) // Type must be `string`
}

func TestRandomRegency(t *testing.T) {
	// randProvince a references random regencies
	randProvince := RandomProvince()

	randRegency := RandomRegency(randProvince)

	assert.NotEmpty(t, randRegency)                                 // Not Empty data
	assert.Equal(t, "string", reflect.TypeOf(randRegency).String()) // Type must be `string`
}
