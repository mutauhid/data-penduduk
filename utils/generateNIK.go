package utils

import (
	"data-penduduk/models"
	"math/rand"
	"strconv"
)

func GenerateNIK(people models.People) string {
	var nik string
	var dobString string
	var randomNumber int

	dobString = people.DOB.Format("300698")
	randomNumber = 1000 + rand.Intn(9000)
	randomString := strconv.Itoa(randomNumber)

	nik = people.ProvinceID + people.RegencyID + people.DistrictID + dobString + randomString

	return nik

}
