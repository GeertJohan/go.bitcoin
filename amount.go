package bitcoin

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// One bitcoin consists of 100,000,000 (100-milion) satoshi's.
const (
	SatoshisPerBitcoin = 100000000
	MaximumBitcoins    = 21e6
	MaximumSatoshis    = MaximumBitcoins * SatoshisPerBitcoin
)

var (
	// Regular expression to test for a string to be a valid strict bitcoin string
	regexpRoundBitcoinsString  = regexp.MustCompile(`^[0-9]+$`)
	regexpLooseBitcoinsString  = regexp.MustCompile(`^[0-9]+\.[0-9]{1,8}$`)
	regexpStrictBitcoinsString = regexp.MustCompile(`^[0-9]+\.[0-9]{8}$`)

	// Error to be returned when a given bitcoin string isn't valid.
	errorInvalidBitcoinsString       = errors.New(`Invalid bitcoins string. Expecting a string that conforms to either '[0-9]+' or '[0-9]+\\.[0-9]{1,8}' or '[0-9]+\\.[0-9]{8}'.`)
	errorInvalidRoundBitcoinsString  = errors.New(`Invalid round bitcoins string. Expecting a string that conforms to '[0-9]+'.`)
	errorInvalidLooseBitcoinsString  = errors.New(`Invalid strict bitcoins string. Expecting a string that conforms to '[0-9]+\.[0-9]{1,8}'.`)
	errorInvalidStrictBitcoinsString = errors.New(`Invalid strict bitcoins string. Expecting a string that conforms to '[0-9]+\\.[0-9]{8}'.`)

	ErrTooBig = errors.New("Amount exceeds maximum possible bitcoin value.")
)

// Amount represents any bitcoin value and presents convenient methods for calculation and formatting (pretty-printing).
// Amount is not linked to a certain address or transaction.
type Amount uint64

// Formats the Amount as full bitcoin string
// The returned string always complies to the regex `[0-9]+\.[0-9]{8}`.
// e.g.: 0.12345678 or 10.01020304 or 0.00100000 or 1004.12345678
func (a Amount) String() string {
	var mayor, minor string
	satoshisString := strconv.FormatUint(uint64(a), 10)
	if len(satoshisString) > 8 {
		mayor = satoshisString[:len(satoshisString)-8]
		minor = satoshisString[len(satoshisString)-8:]
	} else {
		mayor = "0"
		minor = strings.Repeat("0", 8-len(satoshisString)) + satoshisString
	}
	return fmt.Sprintf("%s.%s", mayor, minor)
}

// Returns value as amount in satoshi's, formated as base 10 decimal string. (1 BTC = 100,000,000).
func (a Amount) SatoshisString() string {
	return strconv.FormatUint(uint64(a), 10)
}

// Implementing the json.Unmashaler inteface. Checks and parses given `data []byte` with satoshisFromBitcoinsString.
func (a *Amount) UnmarshalJSON(data []byte) error {
	s, err := satoshisFromBitcoinsString(string(data))
	if err != nil {
		return err
	}
	*a = Amount(s)
	return nil
}

// Implementing the json.Marshaler interface. Returned []byte has same contents as Amount.String().
func (a Amount) MarshalJSON() ([]byte, error) {
	return []byte(a.String()), nil
}

// Create new Amount object from a bitcoin string.
func AmountFromBitcoinsString(bitcoins string) (Amount, error) {
	s, err := satoshisFromBitcoinsString(bitcoins)
	if s > MaximumSatoshis {
		return 0, ErrTooBig
	}
	return Amount(s), err
}

// Internal helper function. Checks and converts bitcoin string to satoshi's.
func satoshisFromBitcoinsString(bitcoins string) (uint64, error) {

	if regexpStrictBitcoinsString.MatchString(bitcoins) {
		return satoshisFromStrictBitcoinsString(bitcoins)
	}

	if regexpLooseBitcoinsString.MatchString(bitcoins) {
		return satoshisFromLooseBitcoinsString(bitcoins)
	}

	if regexpRoundBitcoinsString.MatchString(bitcoins) {
		return satoshisFromRoundBitcoinsString(bitcoins)
	}

	// couldn't try any valid format.
	return 0, errorInvalidBitcoinsString
}

func satoshisFromRoundBitcoinsString(bitcoins string) (uint64, error) {
	// Check that given string is valid.
	if !regexpRoundBitcoinsString.MatchString(bitcoins) {
		return 0, errorInvalidRoundBitcoinsString
	}
	// convert bitcoins string to bitcoinsUint64
	bitcoinsUint64, err := strconv.ParseUint(bitcoins, 10, 64)
	// multiply bitcoinsUint64 with amount of satoshis in a bitcoin.. satoshis is what we want.
	satoshis := bitcoinsUint64 * SatoshisPerBitcoin
	// done
	return satoshis, err
}

func satoshisFromLooseBitcoinsString(bitcoins string) (uint64, error) {
	// Check that given string is valid.
	if !regexpLooseBitcoinsString.MatchString(bitcoins) {
		return 0, errorInvalidLooseBitcoinsString
	}

	// Split fields on the point.
	fields := strings.Split(bitcoins, ".")

	// Glue the fields together again (without a dot) and append with zero's to have the string represent the amount of satoshis
	satoshiString := fields[0] + fields[1] + strings.Repeat("0", 8-len(fields[1]))
	satoshis, _ := strconv.ParseUint(satoshiString, 10, 64) // discard error, we're pretty sure that the given string represents a valid uint64.

	return satoshis, nil
}

// get amount of satoshis from a strict bitcoin string. A strict bitcoin string complies to [0-9]+\.[0-9]{8}
func satoshisFromStrictBitcoinsString(bitcoins string) (uint64, error) {
	// Check if given string is valid.
	if !regexpStrictBitcoinsString.MatchString(bitcoins) {
		return 0, errorInvalidStrictBitcoinsString
	}

	// Remove the dot from the bitcoins string
	// We have a strict/complete bitcoin string so removing the dot will result in the amount of satoshi's as string
	bitcoins = strings.Replace(bitcoins, ".", "", 1)

	// format string
	satoshis, err := strconv.ParseUint(bitcoins, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Could not convert bitcoin string to satoshis: %s", err)
	}

	// All done
	return satoshis, nil
}
