package bitcoin

import (
	"encoding/json"
	"math/rand"
	"testing"
)

var testValues []uint64

func init() {
	r := rand.New(rand.NewSource(87592894858421))
	testValues = []uint64{0, 1,
		MaximumSatoshis - 1, MaximumSatoshis, MaximumSatoshis + 1}
	for len(testValues) < 10000 {
		testValues = append(testValues, abs(r.Int63n(MaximumSatoshis)))
	}
}

func abs(i int64) uint64 {
	if i < 0 {
		return uint64(-i)
	}
	return uint64(i)
}

func TestStringConversion(t *testing.T) {
	for _, i := range testValues {
		a := Amount(i)
		s := a.String()

		as, err := AmountFromBitcoinsString(s)
		if err != nil && !(a > MaximumSatoshis && err == ErrTooBig) {
			t.Errorf("Error parsing %v from %v: %v",
				s, i, err)
		}

		if err == nil && as != a {
			t.Errorf("Expected %v == %v for %v", as, a, i)
		}
	}
}

func TestJSONEncoding(t *testing.T) {
	for _, i := range testValues {
		thing := struct{ A Amount }{}
		thing.A = Amount(i)

		data, err := json.Marshal(&thing)
		if err != nil {
			t.Errorf("Error on %v: %v", i, err)
		}

		thing.A = 0

		err = json.Unmarshal(data, &thing)
		if err != nil && !(thing.A > MaximumSatoshis && err == ErrTooBig) {
			t.Errorf("Error parsing %s from %v: %v",
				data, i, err)
		}

		if err == nil && thing.A != Amount(i) {
			t.Errorf("Expected %v == %v for %v", thing.A, Amount(i), i)
		}
	}

	// Also check one that can't be parsed
	data := []byte(`{"A": "1.x"}`)
	thing := struct{ A Amount }{}
	err := json.Unmarshal(data, &thing)
	if err == nil {
		t.Errorf("Expected to fail parsing %s, got %v", data, thing)
	}
}
