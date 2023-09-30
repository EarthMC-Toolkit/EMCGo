package utils

import (
	"math"
	"math/rand"
	"strings"
	"time"

	lo "github.com/samber/lo"
)

var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var alphabetLen = len(alphabet)
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		ch := alphabet[rand.Intn(alphabetLen)]
		sb.WriteRune(ch)
	}

	return sb.String()
}

func hypot(num, input, radius int) bool {
    return num <= (input + radius) && num >= (input - radius)
}

func manhattan(x1, z1, x2, z2 float64) float64 {
   return math.Abs(x2 - x1) + math.Abs(z2 - z1)
}

func euclidean(x1, z1, x2, z2 float64) float64 {
    return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(z2 - z1, 2))
}

func GetExisting(a1 []any, a2 []string, key string) any {
	arr := lo.Map(a2, func(x string, index int) any {
		return nil // TODO: Implement find here
	})

	if len(arr) > 1 {
		return arr
	} else {
		return arr[0]
	}
}

