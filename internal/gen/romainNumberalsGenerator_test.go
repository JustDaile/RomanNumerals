package gen

import (
	"testing"
)

// RomanNumberalTest defines a tests input and expected output.
type RomanNumberalTest struct {
	Input  int
	Expect string
}

// RomanNumberalTests return a immutable array of RomanNumberalTests.
func RomanNumberalTests() []RomanNumberalTest {
	return []RomanNumberalTest{
		{
			1,
			"I",
		},
		{
			5,
			"V",
		},
		{
			10,
			"X",
		},
		{
			20,
			"XX",
		},
		{
			3999,
			"MMMCMXCIX",
		},
		{
			39,
			"XXXIX",
		},
		{
			246,
			"CCXLVI",
		},
		{
			789,
			"DCCLXXXIX",
		},
		{
			2421,
			"MMCDXXI",
		},
		{
			160,
			"CLX",
		},
		{
			207,
			"CCVII",
		},
		{
			1009,
			"MIX",
		},
		{
			1066,
			"MLXVI",
		},
		// Caveat: Only support numbers between 1 and 3999
		// Out of range, errors will return empty string as described in the romanNumberalGenerator.go
		{
			0,
			"",
		},
		{
			4000,
			"",
		},
	}
}

// TestSymbolsByPowerOfTen tests the SymbolsByPowerOfTen function
// First it tests the range of thousands is between 1 and 3 anything else should return unit out of range error
// Then it checks the remaining powers of ten (Hundreds, Tens, Units) in a nested for loop
func TestSymbolsByPowerOfTen(t *testing.T) {
	// Valid units for the thousands are 1,2,3.
	// The following tests that we can only pass in these units to get a valid result.
	// Otherwise we get an empty string and an error.
	for thousands := 0; thousands <= 4; thousands++ {
		th, e := SymbolsByPowerOfTen(3, thousands)
		if e != nil {
			t.Log(e)
			// We are expecting errors since we are also testing out of range units.
			// So if we get an error and our unit is in range then we fail.
			// Otherwise the error is expected.
			if thousands > 0 && thousands < 3 {
				t.Logf("failed as thousand %v should be in range 0-3", thousands)
				t.Fail()
			}
		} else {
			t.Logf("%v returned %s", thousands, th)
		}
	}

	// Do the other powers of tens (Hundreds, Tens, Units) by
	// iterating the power, then each unit from 0 to 10 inclusively.
	// Include 0 and 10 to ensure that we can only pass in units in range of 1 to 9.
	// The units 0 and 10 should return an error, we test the input was in range and if it was we fail since this would
	// indicate our function is not working as expected.
	//
	// order doesn't matter but makes logging easier to read, since the order of powers will be highest to lowest.
	for power := 2; power >= 0; power-- {
		for unit := 0; unit <= 10; unit++ {
			h, e := SymbolsByPowerOfTen(power, unit)

			if e != nil {
				t.Log(e)
				if unit > 0 && unit < 10 {
					t.Fail()
				}
			} else {
				t.Logf("%v returned %s", unit, h)
			}
		}
	}
}

// TestGenerateRomanNumberals will test a bunch of examples to ensure the correct roman numberals are generated.
func TestGenerateRomanNumberals(t *testing.T) {
	genImpl := RomanNumberalGeneratorImpl{}

	for _, test := range RomanNumberalTests() {
		t.Logf("Expect '%v' when calling generate(%v)", test.Expect, test.Input)
		actual := genImpl.generate(test.Input)
		if test.Expect != actual {
			t.Logf("generate(%v) expected '%s' but was '%s'", test.Input, test.Expect, actual)
			t.Fail()
		}
	}
}
