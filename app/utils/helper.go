package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func InArray(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func StrToHayStack(val string) []string {
	// Split the input string by commas
	result := strings.Split(val, ",")

	// Return the slice of substrings
	return result
}

func ConvertToString(i interface{}) (string, error) {
	str, ok := i.(string)
	if !ok {
		return "", fmt.Errorf("interface{} does not hold a string value")
	}
	return str, nil
}

func HelperRandomNumber(length int) int {
	if length <= 0 {
		return 0
	}

	// Create a new random generator with a seed based on the current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Calculate the minimum and maximum values for the given length
	min := int64(1)
	for i := 1; i < length; i++ {
		min *= 10
	}
	max := min*10 - 1

	// Generate the random number within the desired range
	randomNumber := r.Int63n(max-min+1) + min

	// Ensure the number has the exact number of digits
	result, err := strconv.Atoi(fmt.Sprintf("%0*d", length, randomNumber))
	if err != nil {
		return 0
	}

	return result
}
