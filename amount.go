package bitcoin

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	SatoshiPerBitcoin = 100000000
)

var (
	regexpDecimalBitcoinString = regexp.MustCompile(`[0-9]*\.[0-9]{8}`)

	errorInvalidDecimalBitcoinString = errors.New("Invalid decimal bitcoin")
)

// WARNING!!!! This should not be a float!! float is probably not safe enough!
// TODO: create safe Amount struct with json marshalling and pretty printing.
type Amount struct {
	// Storing any amount as satoshi's (1/100,000,000th BTC) is safe. 
	// uint64 can hold ALL satoshi's in the world and thus it wont overflow (unless Amount is used wrong)
	satoshi uint64
}

func (a *Amount) String() string {
	var mayor, minor string
	satoshiString := strconv.FormatUint(a.satoshi, 10)
	if len(satoshiString) > 8 {
		mayor = satoshiString[:len(satoshiString)-8]
		minor = satoshiString[len(satoshiString)-8:]
	} else {
		mayor = "0"
		minor = strings.Repeat("0", 8-len(satoshiString)) + satoshiString
	}
	return fmt.Sprintf("%s.%s", mayor, minor)
}

func (a *Amount) UnmarshalJSON(data []byte) error {
	s, err := satoshiFromDecimalBitcoinString(string(data))
	if err != nil {
		return err
	}
	a.satoshi = s
	return nil
}

func AmountFromDecimalBitcoinString(decimalBitcoin string) (*Amount, error) {
	s, err := satoshiFromDecimalBitcoinString(decimalBitcoin)
	if err != nil {
		return nil, err
	}
	return &Amount{satoshi: s}, nil
}

func satoshiFromDecimalBitcoinString(decimalBitcoin string) (uint64, error) {

	if !regexpDecimalBitcoinString.MatchString(decimalBitcoin) {
		return 0, errorInvalidDecimalBitcoinString
	}

	// Split fields on the point.
	fields := strings.Split(decimalBitcoin, ".")
	// Get the value in front of the period (mayor).
	mayor, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return 0, err
	}
	// Get the value after the period (minor).
	minor, _ := strconv.ParseUint(fields[1], 10, 64) // discard error, decimalBitcoin has been checked with regexp.

	return (mayor * SatoshiPerBitcoin) + minor, nil
}
