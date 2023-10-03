package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rn "github.com/random-names/go"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomName(isFemale bool) string {
	var firstName string
	var err error

	lastName, err := rn.GetRandomName("./census-90/all.last", &rn.Options{})
	if err != nil {
		fmt.Println(err)
	}

	if isFemale {
		firstName, err = rn.GetRandomName("./census-90/female.first", &rn.Options{})
		if err != nil {
			fmt.Println(err)
		}

	} else {
		firstName, err = rn.GetRandomName("./census-90/male.first", &rn.Options{})
		if err != nil {
			fmt.Println(err)
		}
	}

	return firstName + " " + lastName
}

func RandomGender() string {
	genders := []string{"male", "female", "other"}
	return genders[rand.Intn(2)]
}

func RandomEmail() string {
	return RandomString(10) + "@example.com"
}

func RandomPhone() string {
	return strconv.Itoa(int(RandomInt(1000000000, 9999999999)))
}

func RandomAge() int {
	return int(RandomInt(18, 100))
}

func RandomTopic() string {
	topics := []string{"Math", "English", "Physics", "Chemistry", "Biology", "History", "Geography", "Literature", "Music", "Art", "Sport"}
	return topics[rand.Intn(10)]
}

func RandomCourse() string {
	courses := []string{"CS", "FA"}
	return courses[rand.Intn(2)]
}

func RandomCity() string {
	cities := []string{"Seoul", "Berlin", "Tokyo", "Paris", "Praha", "Venice", "Boston", "Madrid", "Saigon", "Rome", "New York"}
	return cities[rand.Intn(10)]
}

func RandomCountry() string {
	countries := []string{"Vietnam", "US", "UK", "Japan", "China", "Korea", "Russia", "France", "Germany", "Italy", "Spain"}
	return countries[rand.Intn(10)]
}
