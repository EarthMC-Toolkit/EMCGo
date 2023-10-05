package utils

import (
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	lo "github.com/samber/lo"
	"github.com/sanity-io/litter"
	"golang.org/x/exp/slices"
)

var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var alphabetLen = len(alphabet)
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func AsBool(v string) bool {
	if strings.ToLower(v) == "true" {
		return true
	}

	return false
}

func FloatsToInts(arr []float64) []int {
	return lo.Map(arr, func (num float64, _ int) int {
		return int(num)
	})
}

func Range(args []float64) int {
	if len(args) == 0 {
		return 0 // Handle the case when the input slice is empty
	}

	return int((slices.Max(args) + slices.Min(args)) / 2)
}

func CalcArea(X, Z []float64, numPoints int, divisor float64) int {
	var area float64 = 0 
	j := numPoints - 1

    for i := 0; i < numPoints; i++ { 
        area += (X[j] + X[i]) * (Z[j] - Z[i]) 
        j = i
    }

    return int(math.Abs(float64(area / 2)) / divisor)
}

func CleanString(str string) string {
	re := regexp.MustCompile(`((&#34)|(&\w[a-z0-9].|&[0-9kmnola-z]));`)

	cleaned := re.ReplaceAllString(str, "")

	cleaned = strings.ReplaceAll(str, "&quot;", "\"")
	cleaned = strings.ReplaceAll(str, "&#039;", "\"")
	
	return cleaned
}

func Prettify(i interface{}) string {
	litter.Config.StripPackageNames = true
	return litter.Sdump(i)
}

func HexToInt(hex string) int {
	str := strings.Replace(hex, "0x", "", -1)
	output, _ := strconv.ParseUint(str, 16, 32)

	return int(output)
}

func FormatTimestamp(unixTs float64) string {
	return strconv.FormatFloat(unixTs / 1000, 'f', 0, 64)
}

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
	GetName()	string
}

type Entity struct {
	Name		string
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