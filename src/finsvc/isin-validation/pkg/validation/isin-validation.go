package validation

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	countryCodeLength        = 2
	securityIdentifierLength = 9
	checkDigitLength         = 1

	validIsinLength = countryCodeLength + securityIdentifierLength + checkDigitLength

	lengthMismatchError                 = "Provided ISIN doesn't have a valid length. Provided: %d, required %d"
	firstTwoSymbolsNotAlphabeticalError = "Provided ISIN doesn't have enough alphabetical symbols at the beginning. " +
		"Provided: %s, required country code symbols: %d"
	lastSymbolIsNotNumberError = "Provided ISIN doesn't have the last symbol number. Provided ISIN: %s"

	checkSumIsNotValidError = "Checksum digit is not valid for the provided ISIN. Expected checksum: %d"
)

//The algorithm for calculating the ISIN check digit from is the following:
//1. Convert any alphabetic letters to their numeric equivalents using the above table.
//   Beginning with the least significant digit (on the right), multiply every other digit by 2.
//   Add up the resulting digits, calling the result SUM.
//2. Find the smallest number ending with a zero that is greater than or equal to SUM, and call it VALUE.
//3. Subtract SUM from VALUE, giving the check digit.

func ValidateIsin(isin string) *Result {
	fmt.Println("Validating ISIN", isin)

	runesRepresenstationISIN := []rune(isin)
	violations := quickValidations(&runesRepresenstationISIN, isin)
	if !(len(violations) == 0) {
		return &Result{
			Message: strings.Join(violations, "; "),
			Status:  NotValid,
		}
	}

	isinLength := len(runesRepresenstationISIN) - 1
	runesWithoutCheckSum := runesRepresenstationISIN[0:isinLength]
	numbers := convertToNumbers(&runesWithoutCheckSum)

	sum := calculateSum(numbers)
	roundValue := roundUpValue(sum)
	checksumDigit := int(runesRepresenstationISIN[isinLength] - '0')
	expectedChecksum := roundValue - sum

	validIsin := expectedChecksum == checksumDigit

	if !validIsin {
		return &Result{
			Message: fmt.Sprintf(checkSumIsNotValidError, expectedChecksum),
			Status:  NotValid,
		}
	}

	return &Result{
		Message: "Validation successful",
		Status:  OK,
	}
}

func quickValidations(runesIsin *[]rune, isin string) []string {
	violations := make([]string, 0)
	providedIsinLength := len(*runesIsin)

	if providedIsinLength != validIsinLength {
		violations = append(violations, fmt.Sprintf(lengthMismatchError, providedIsinLength, validIsinLength))
	} else {
		if !unicode.IsNumber((*runesIsin)[validIsinLength-1]) {
			violations = append(violations, fmt.Sprintf(lastSymbolIsNotNumberError, isin))
		}
	}
	if providedIsinLength > 2 && (!unicode.IsLetter((*runesIsin)[0]) || !unicode.IsLetter((*runesIsin)[1])) {
		violations = append(violations, fmt.Sprintf(firstTwoSymbolsNotAlphabeticalError, isin, countryCodeLength))
	}
	return violations
}

// 6 0 4 8 0 3 14 8 6 3 2 0 0
func calculateSum(numbers []int) int {
	sum := 0
	count := 0
	valueToAdd := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		value := numbers[i]
		if count%2 == 0 {
			valueToAdd = value * 2
		} else {
			valueToAdd = value
		}
		if valueToAdd > 10 {
			firstDigit := valueToAdd % 10
			secondDigit := valueToAdd / 10
			sum = sum + firstDigit + secondDigit
		} else {
			sum = sum + valueToAdd
		}
		count = count + 1
	}
	return sum
}

func roundUpValue(sum int) int {
	tmp := sum + 9 //45 + 9 = 54 - (54 % 10) = 50
	return tmp - tmp%10
}

func convertToNumbers(isin *[]rune) []int {
	result := make([]int, 0, len(*isin)+10)
	for _, symbol := range *isin {
		if unicode.IsNumber(symbol) {
			result = append(result, int(symbol-'0'))
		} else {
			numberEquiualent := calculateNumberEquiualent(symbol)
			result = append(result, int(numberEquiualent/10), int(numberEquiualent%10))
		}
	}
	return result
}

func calculateNumberEquiualent(symbol rune) int32 {
	return symbol - 55
}
