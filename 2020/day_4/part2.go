package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var eyeColors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

type ValidationFunc func(string) bool

var requiredFields = map[string]ValidationFunc{
	"byr": func(value string) bool {
		year, _ := strconv.Atoi(value)
		return year >= 1920 && year <= 2002
	},
	"iyr": func(value string) bool {
		year, _ := strconv.Atoi(value)
		return year >= 2010 && year <= 2020
	},
	"eyr": func(value string) bool {
		year, _ := strconv.Atoi(value)
		return year >= 2020 && year <= 2030
	},
	"hgt": func(value string) bool {
		if strings.Index(value, "cm") != -1 {
			splitted := strings.Split(value, "cm")
			height, _ := strconv.Atoi(splitted[0])
			return height >= 150 && height <= 193
		}
		splitted := strings.Split(value, "in")
		height, _ := strconv.Atoi(splitted[0])
		return height >= 59 && height <= 76
	},
	"hcl": func(value string) bool {
		splitted := strings.Split(value, "#")
		if len(splitted) != 2 {
			return false
		}
		_, err := strconv.ParseInt(splitted[1], 16, 64)
		return err == nil
	},
	"ecl": func(value string) bool {
		_, ok := eyeColors[value]
		return ok
	},
	"pid": func(value string) bool {
		_, err := strconv.Atoi(value)
		return err == nil && len(value) == 9
	},
}

func main() {
	file := readInputToString()

	passports := strings.Split(file, "\n\n")

	validPassportCount := 0
	for i := range passports {
		validFieldCount := 0

		fields := getFields(passports[i])
		for i := range fields {
			fieldAndValue := strings.Split(fields[i], ":")
			if isValid, ok := requiredFields[fieldAndValue[0]]; ok {
				if isValid(fieldAndValue[1]) {
					validFieldCount++
				}
			}
		}

		if validFieldCount == len(requiredFields) {
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
	return result
}

func readInputToString() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}
