# `Jabufaker` for generating a fake sample data
Jabufaker is a package for generating a fake sample data. Data generated based on Indonesia country.

And also this package is build based on developer `(Rizky Darmawan)` anxiety in thinking about what sample data to include in an application that is being developed.

The beginning of code in this package, obtained when followed a course udemy with title `Backend Master Class [Golang + Postgres + Kubernetes + gRPC]`.

You will find beginning the code this package in there. And me added some code to generate some data I needed. You also can use and contribute by adding some code you need to this package.

Criticism and suggestions about this package are very useful for me. My contacts can be found in the [overview](https://github.com/rizkydarmawan-letenk) section of this repository.

# Contents
* [Jabufaker](#jabufaker-for-generating-a-fake-sample-data)
    * [Contents](#contents) 
    * [Installation](#installation)
    * [Function to use](#function-to-use)
        * [Generate random number](#generate-random-number)
        * [Generate random string](#generate-random-string)
        * [Generate random person](#generate-random-person)
        * [Generate random money](#generate-random-money)
        * [Generate random currency](#generate-random-currency)
        * [Generate random email](#generate-random-email)
        * [Generate random provincies](#generate-random-province)
        * [Generate random regencies](#generate-random-regencies)

# Installation
To install jabufaker package, follow these steps:
1. You can use the below Go command:
```
go get -u github.com/rizkydarmawan-letenk/jabufaker
```

2. Import it in your code:
```
import "github.com/rizkydarmawan-letenk/jabufaker"
```

# Function to use
## Generate random number
This function is used to generate a number integer.
```
RandomInt(min, max)
```
- `min` is minimal number to be generated.
- `max` is maximal number to be generated.

**Example:**
```
randInt := jabufaker.RandomInt(0, 100)
fmt.Println(randInt)
```

## Generate random string
This function is used to generate a string.
```
RandomString(n)
```
- `n` is much string to be generated.

**Example:**
```
randString := jabufaker.RandomString(10)
fmt.Println(randString)
```

## Generate random person.
This function is used to generate a name person.
```
RandomPerson()
```
**Example:**
```
randPerson := jabufaker.RandomPerson()
fmt.Println(randPerson)
```

## Generate random money.
This function is used to generate a random number integer start from 0 to 100.
```
RandomMoney()
```

**Example:**
```
randMoney := jabufaker.RandomMoney()
fmt.Println(randMoney)
```

## Generate random currency.
This function is used to generate a random currency. Currently only available IDR, USD, EUR.
```
 RandomCurrency()
```

**Example:**
```
randCurrencies := jabufaker.RandomCurrency()
fmt.Println(randCurrencies)
```


## Generate random email.
This function is used to generate random email address.
```
RandomEmail()
```

**Example:**
```
randEmail := jabufaker.RandomEmail()
fmt.Println(randEmail)
```

## Generate random province.
This function is used to generate random province.
**Note:** Currently only available some provinces are based on Sumatera Island.
```
RandomProvince()
```

**Example:**
```
randProvince := jabufaker.RandomProvince()
fmt.Println(randProvince)
```
## Generate random regencies
This function is used to generate random regencies, which is the regencies generated based on provinces selected from function RandomProvince(). 

**Note:** Make sure before generating random regencies, first generate provinces and make as argument on function RandomRegencies().
```
RandomRegency(random_province_selected)
```

**Example:**
```
randProvince := jabufaker.RandomProvince()

randRegency := jabufaker.RandomRegency(randProvince)
fmt.Println(randRegency)
```
