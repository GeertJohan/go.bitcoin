package bitcoin

import (
	"fmt"
	"math"
)

const (
	SatoshiPerBitcoin = float64(100000000)
)

// WARNING!!!! This should not be a float!! float is probably not safe enough!
// TODO: create safe Amount struct with json marshalling and pretty printing.
type Amount struct {
	// Storing any amount as satoshi's (1/100,000,000th BTC) is safe. 
	// uint64 can hold ALL satoshi's in the world and thus it wont overflow (unless Amount is used wrong)
	satoshi uint64
}

func (a Amount) String() string {
	mayor := math.Floor(float64(a.satoshi) / SatoshiPerBitcoin)
	minor := math.Mod(float64(a.satoshi), SatoshiPerBitcoin)
	return fmt.Sprintf("%d.%d", mayor, minor)
}

func AmountFromBitcoinString(btc string) Amount {

}
