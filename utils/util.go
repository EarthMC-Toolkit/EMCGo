package utils

import (
	"math"
	"math/rand"
	"reflect"
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

func Hypot(num, input, radius int) bool {
    return num <= (input + radius) && num >= (input - radius)
}

func Manhattan(x1, z1, x2, z2 float64) float64 {
   return math.Abs(x2 - x1) + math.Abs(z2 - z1)
}

func Euclidean(x1, z1, x2, z2 float64) float64 {
    return math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(z2 - z1, 2))
}

func GetFieldValue(item interface{}, key string) string {
    field := reflect.ValueOf(item).FieldByName(key)
    return lo.Ternary(field.IsValid(), field.String(), "")
}

type EntityAccess interface {
	GetName() string
}

type Entity struct {
	Name 	string
}

func (e Entity) GetName() string {
	return e.Name
}

func GetExisting[T EntityAccess](a1 []T, a2 []string) any {
	arr := lo.Map(a2, func(v string, index int) any {
		val, found := lo.Find(a1, func(el T) bool {
			return strings.EqualFold(v, el.GetName())
		})

		return lo.Ternary[any](found, val, nil)
	})

	return lo.Ternary[any](len(arr) > 1, arr, arr[0])
}