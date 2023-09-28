package random

import (
	"math/rand"
	"strconv"
)

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomTopic(isUnique bool) string {
	topics := []string{"Math", "English", "Physics",
		"Chemistry", "Biology", "History",
		"Geography", "Literature", "Music",
		"Art", "Sport", "Computer Science",
		"Foreign Language", "Economics", "Politics",
	}

	// unique from topics
	if isUnique {
		return topics[rand.Intn(len(topics))] + " " + strconv.Itoa(int(RandomInt(1, 10)))
	}

	return topics[rand.Intn(len(topics))]
}
