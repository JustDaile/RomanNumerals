package gen

import (
	"fmt"
)

// RomanNumberalGenerator golang version of the interface given in Java
type RomanNumberalGenerator interface {
	generate(int) string
}

// RomanNumberalGeneratorImpl implementation of the RomanNumberalGenerator interface
type RomanNumberalGeneratorImpl struct{}

// SymbolsByPowerOfTen returns a roman numberal for a unit given is power of ten.
// Its minumum range is 1 and its maximum range is 3999
// returns a string containing the roman numberal symbol or an error if out of range.
func SymbolsByPowerOfTen(pow10 int, unit int) (string, error) {
	if unit < 1 || unit > 9 {
		return "", fmt.Errorf("unit %v out of range 1-9", unit)
	}
	switch pow10 {
	case 0: // units
		return []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}[unit-1], nil
	case 1: // tens
		return []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}[unit-1], nil
	case 2: // hundreds
		return []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}[unit-1], nil
	case 3: // thousands
		if unit > 3 {
			return "", fmt.Errorf("unsupported thousand %v - max is 3", unit)
		}
		return []string{"M", "MM", "MMM"}[unit-1], nil
	}

	return "", fmt.Errorf("unsupported power of 10 %v", pow10)
}

// This implementation is going to use the standard form table described here: https://en.wikipedia.org/wiki/Roman_numerals
// First I will get the unit of each power of ten (Thousands, hundreds, Tens, Units) from the input number.
// I will use the hardcoded table to lookup the correct symbol for the unit in respect of its original power.
// Finally I will concatinate the symbols together and return the result.
//
// Since nothing was specified about how we should handle invalid inputs an empty string will be returned if any error occurs.
// I'm sticking with the interface defined in the example, so I will not be returning an error.
// Alteratively we could panic but this would require the caller attempts a recovery (try recover)
// The simpliest approach is to return an empty string. However this will not inform the caller of the exact problem.
func (RomanNumberalGeneratorImpl) generate(input int) string {
	// I thought of a few ways to extract each number unit from the input number
	// - Convert to string/slice of characters and go from their
	// - Use bitwise operations to exact each one
	//
	// I decided to use simple maths instead to determine the unit the thousands, hundreds, etc
	thousands := input / 1000
	hundreds := (input - 1000*thousands) / 100
	tens := ((input - 1000*thousands) - 100*hundreds) / 10
	units := (((input - 1000*thousands) - 100*hundreds) - 10*tens)

	// Concatinate symbols to this string for the return
	strBuilder := ""

	// We could do each power individually as follows, but this is a lot of repetition
	//
	// if thousands > 0 {
	// 	thouSymbol, e := SymbolsByPowerOfTen(3, thousands)
	// 	if e != nil {
	// 		return ""
	// 	}
	// 	strBuilder += thouSymbol
	// }
	//
	// if hundreds > 0 {
	// 	hundSymbol, e := SymbolsByPowerOfTen(2, hundreds)
	// 	if e != nil {
	// 		return ""
	// 	}
	// 	strBuilder += hundSymbol
	// }
	//
	// if tens > 0 {
	// 	tensSymbol, e := SymbolsByPowerOfTen(1, tens)
	// 	if e != nil {
	// 		return ""
	// 	}
	// 	strBuilder += tensSymbol
	// }
	//
	// if units > 0 {
	// 	unitsSymbol, e := SymbolsByPowerOfTen(0, units)
	// 	if e != nil {
	// 		return ""
	// 	}
	// 	strBuilder += unitsSymbol
	// }

	// We could use a struct or in this case an anonomous struct array so we can use a loop instead of the above example
	// This cuts down on repetition, but looses some readibility.
	loopableUnits := []struct {
		Power int
		Unit  int
	}{
		{
			3, thousands,
		},
		{
			2, hundreds,
		},
		{
			1, tens,
		},
		{
			0, units,
		},
	}
	for _, powerUnit := range loopableUnits {
		if powerUnit.Unit > 0 {
			unitsSymbol, e := SymbolsByPowerOfTen(powerUnit.Power, powerUnit.Unit)
			if e != nil {
				return ""
			}
			strBuilder += unitsSymbol
		}
	}

	return strBuilder
}
