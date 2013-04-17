package bitcoin

import (
	"encoding/json"
	"testing"
)

const example = `
{
    "amount": 0.0,
    "blockhash": "00000000022590baa421f6d39bc676020847c575481b7e602cec50d863d751c3",
    "blockindex": 2,
    "blocktime": 1366164482,
    "comment": "test one, eh",
    "confirmations": 2,
    "details": [
        {
            "account": "test2",
            "address": "n1fuNTv1QoPijZ3EoNpfPqBostBSWATstM",
            "amount": -1.08,
            "category": "send",
            "fee": -0.0005
        },
        {
            "account": "",
            "address": "n1fuNTv1QoPijZ3EoNpfPqBostBSWATstM",
            "amount": 1.08,
            "category": "receive"
        }
    ],
    "fee": -0.0005,
    "time": 1366164288,
    "timereceived": 1366164288,
    "to": "this is a test for you",
    "txid": "0154da5ab95cd850bc90f6a8e7d89a19b771386285d892daa6a4df47df999266"
}
`

func TestTransactionParsing(t *testing.T) {
	tx := Transaction{}
	err := json.Unmarshal([]byte(example), &tx)
	if err != nil {
		t.Fatalf("Error parsing transaction: %v", err)
	}

	// TODO:  Add assertions here.

	// Should probably confirm these do something sensible.
	tx.BlockTime()
	tx.TransactionTime()
	tx.TransactionReceivedTime()
}
