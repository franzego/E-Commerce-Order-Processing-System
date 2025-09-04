package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().Unix())
	//r := rand.New(src)
}

var currencies = []string{"USD", "NGR", "GBP", "EUR"}
var productNames = []string{"Iphones", "Laptop", "TV", "Mouse"}

// Generate random currencies
func RandomCurrency() string {
	return currencies[rand.Intn(len(currencies))]

}

// Generate random product names
func RandomProduct() string {
	return productNames[rand.Intn(len(productNames))]
}

// Generate random float between min and max
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Generate random int between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
