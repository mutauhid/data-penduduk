package utils

import (
	"data-penduduk/models"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func GenerateNIK(people models.People) string {
	var nik string
	var dobString string
	var randomNumber int

	dobString = people.DOB.Format("02-01-2006")
	fmt.Println("dobstring", dobString)
	var split = strings.Split(dobString, "-")
	var year = split[2]
	if len(year) > 2 {
		year = year[len(year)-2:]
	}
	var month = split[1]
	var date = split[0]
	fmt.Println("datesplit", date)
	if people.Gender == "wanita" {
		dateInt, _ := strconv.Atoi(date)
		fmt.Println(dateInt)
		dateInt += 40
		fmt.Println("date", dateInt)

		date = strconv.Itoa(dateInt)
		fmt.Println(date)
	}
	randomNumber = 1000 + rand.Intn(9000)
	randomString := strconv.Itoa(randomNumber)

	nik = people.ProvinceID + people.RegencyID + people.DistrictID + date + month + year + randomString

	return nik

}
