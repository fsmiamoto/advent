package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var requiredFields = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
}

func main() {
	file := readInputToString()

	passports := strings.Split(file, "\n\n")

	validPassportCount := 0
	for i := range passports {
		requiredFieldCount := 0
		fields := getFields(passports[i])
		for j := range fields {
			if _, ok := requiredFields[fields[j]]; ok {
				requiredFieldCount++
			}
		}

		if requiredFieldCount == len(requiredFields) {
			validPassportCount++
		}
	}

	fmt.Println(validPassportCount)
}

func getFields(s string) []string {
	var result []string
	t := strings.Split(s, "\n")
	for i := range t {
		ts := strings.Split(t[i], " ")
		result = append(result, ts...)
	}
	var finalResult []string
	for i := range result {
		a := strings.Split(result[i], ":")
		finalResult = append(finalResult, a[0])
	}
	return finalResult
}

func readInputToString() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}
