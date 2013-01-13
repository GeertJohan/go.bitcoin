package bitcoin

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// One bitcoin consists of 100,000,000 (100-milion) satoshi's.
const SatoshisPerBitcoin = 100000000

var (
	// Regular expression to test for a string to be a valid full bitcoin string
	regexpFullBitcoinsString = regexp.MustCompile(`[0-9]+\.[0-9]{8}`)
	// Error to be returned when a given bitcoin string isn't valid.
	errorInvalidBitcoinString = errors.New("Invalid bitcoin string")
)

// Amount represents any bitcoin value and presents convenient methods for calculation and formatting (pretty-printing).
// Amount is not linked to a certain address or transaction.
type Amount struct {
	// Storing any amount as satoshi's (1/100,000,000th BTC) is safe. 
	// uint64 can hold ALL satoshi's in the world and thus it wont overflow (unless Amount is used wrong)
	satoshis uint64
}

// Formats the Amount as full bitcoin string
// The returned string always complies to the regex `[0-9]+\.[0-9]{8}`.
// e.g.: 0.12345678 or 10.01020304 or 0.00100000 or 1004.12345678
func (a *Amount) String() string {
	var mayor, minor string
	satoshisString := strconv.FormatUint(a.satoshis, 10)
	if len(satoshisString) > 8 {
		mayor = satoshisString[:len(satoshisString)-8]
		minor = satoshisString[len(satoshisString)-8:]
	} else {
		mayor = "0"
		minor = strings.Repeat("0", 8-len(satoshisString)) + satoshisString
	}
	return fmt.Sprintf("%s.%s", mayor, minor)
}

// Returns value as amount in satoshi's, as uint64. uint64 can theoretically hold all satoshi's in the universe. (1 BTC = 100,000,000 Satoshi).
func (a *Amount) SatoshisUint64() uint64 {
	return a.satoshis
}

// Returns value as amount in satoshi's, formated as base 10 decimal string. (1 BTC = 100,000,000).
func (a *Amount) SatoshisString() string {
	return strconv.FormatUint(a.satoshis, 10)
}

// Implementing the json.Unmashaler inteface. Checks and parses given `data []byte` with satoshisFromBitcoinsString.
func (a *Amount) UnmarshalJSON(data []byte) error {
	s, err := satoshisFromBitcoinsString(string(data))
	if err != nil {
		return err
	}
	a.satoshis = s
	return nil
}

// Implementing the json.Marshaler interface. Returned []byte has same contents as Amount.String().
func (a *Amount) MarshalJSON() ([]byte, error) {
	return []byte(a.String()), nil
}

// Create new Amount object from a bitcoin string.
func AmountFromBitcoinsString(bitcoins string) (*Amount, error) {
	s, err := satoshisFromBitcoinsString(bitcoins)
	if err != nil {
		return nil, err
	}
	return &Amount{satoshis: s}, nil
}

// Create new Amount object from amount Satoshi's uint64.
func AmountFromSatoshisUint64(satoshis uint64) *Amount {
	return &Amount{
		satoshis: satoshis,
	}
}

// Internal helper function. Checks and converts bitcoin string to satoshi's.
func satoshisFromBitcoinsString(bitcoins string) (uint64, error) {

	if !regexpFullBitcoinsString.MatchString(bitcoins) {
		return 0, errorInvalidBitcoinString
	}

	// Wouldn't it be a lot easier to remove the period from the string and parse the whole string to uint64 satoshi's?
	// We know for a fact that the string has 8 point decimal fracions, so removing the dot will give the value as satoshi's.
	// Or do we want to accept decimal bitcoin strings that do not have 8 point decimal fractions? (e.g. "0.25")
	// The last one wouldn't be so bad, as user input is likely to have less points... And this lib should support (and check) user input.
	// TODO: accept [0-9]+\.[0-9]{1,8}
	//				(between 1 and 8 decimal fraction points)
	// TODO: accept [0-9]+
	//				(no decimal fraction)
	// Use switch true ?
	// Use if/elseif/elseif/else ?

	// Split fields on the point.
	fields := strings.Split(bitcoins, ".")
	// Get the value in front of the period (mayor).
	mayor, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return 0, err
	}
	// Get the value after the period (minor).
	minor, _ := strconv.ParseUint(fields[1], 10, 64) // discard error, decimalBitcoin has been checked with regexp.

	return (mayor * SatoshisPerBitcoin) + minor, nil
}
