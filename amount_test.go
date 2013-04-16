package bitcoin

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

func TestSatoshiStringConversion(t *testing.T) {
	f := func(i uint64) bool {
		a := Amount(i)
		s := a.SatoshisString()

		as, err := AmountFromBitcoinsString(s)
		if err != nil {
			return a > MaximumValue && err == ErrTooBig
		}

		t.Logf("%v -> %v (%v)", i, as, a)

		return as == a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStringConversion(t *testing.T) {
	f := func(i uint64) bool {
		a := Amount(i)
		s := a.String()

		as, err := AmountFromBitcoinsString(s)
		if err != nil {
			return a > MaximumValue && err == ErrTooBig
		}

		return as == a
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestJSONEncoding(t *testing.T) {
	f := func(i uint64) bool {
		thing := struct {
			A Amount
		}{}
		thing.A = Amount(i)

		data, err := json.Marshal(&thing)
		if err != nil {
			return false
		}

		thing.A = 0

		err = json.Unmarshal(data, &thing)
		if err != nil {
			return i > MaximumValue && err == ErrTooBig
		}

		return thing.A == Amount(i)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
