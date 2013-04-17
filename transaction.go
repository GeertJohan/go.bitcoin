package bitcoin

import (
	"time"
)

// A transaction
type Transaction struct {
	Amount        Amount
	Blockhash     string
	Blockindex    int
	BlockTS       uint32 `json:"blocktime"`
	Comment       string
	Confirmations int
	Fee           Amount
	TXTS          uint32 `json:"time"`
	TXRCVTS       uint32 `json:"timereceived"`
	To            string
	TXID          string
	Details       []struct {
		Account string
		Address string
		Amount  Amount
		Catgory string
		Fee     Amount
	}
}

func (t Transaction) BlockTime() time.Time {
	return time.Unix(int64(t.BlockTS), 0)
}

func (t Transaction) TransactionTime() time.Time {
	return time.Unix(int64(t.TXTS), 0)
}

func (t Transaction) TransactionReceivedTime() time.Time {
	return time.Unix(int64(t.TXRCVTS), 0)
}
