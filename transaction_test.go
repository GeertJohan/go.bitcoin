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

const exampleRaw = `
{
    "blockhash": "0000000000000cb851f3d140a26ce403d2bacf7e491c05cb83c229a590424bbf",
    "blocktime": 1308126369,
    "confirmations": 98861,
    "hex": "0100000001e832fc68d36c70bb004024bedd7e4c03773ae271aef8512ce5b10619cac47657000000008c493046022100f31a55708e58e90734d4c9bcbad6bf5efbf7c3f4947b3425f7bfb77071007086022100e292f8500ccbd09d44e9e5974e3fbeb4cee4abf34703b0dc92dde4ed22e96ddd0141043699d871786b743470137a9d4d9eeb5e0d6a3c82ee700fbe934446c29af861c0c72e751c3f31b8fcb2ed5074a1a3758883d40e8c53e036b02da92ef0b6c5d0a8ffffffff0200c08946010000001976a914a1a621ba261e3d9b3986af158b1274a21c2b7a2788ac80969800000000001976a9147abe8906e9293e41e0db37df516ad7088eb1c6a988ac00000000",
    "locktime": 0,
    "time": 1308126369,
    "txid": "619d6871027c7cdafa2c4ce76d9a5b351fa6e53b1dad9e9c2b63c19b40d07698",
    "version": 1,
    "vin": [
        {
            "scriptSig": {
                "asm": "3046022100f31a55708e58e90734d4c9bcbad6bf5efbf7c3f4947b3425f7bfb77071007086022100e292f8500ccbd09d44e9e5974e3fbeb4cee4abf34703b0dc92dde4ed22e96ddd01 043699d871786b743470137a9d4d9eeb5e0d6a3c82ee700fbe934446c29af861c0c72e751c3f31b8fcb2ed5074a1a3758883d40e8c53e036b02da92ef0b6c5d0a8",
                "hex": "493046022100f31a55708e58e90734d4c9bcbad6bf5efbf7c3f4947b3425f7bfb77071007086022100e292f8500ccbd09d44e9e5974e3fbeb4cee4abf34703b0dc92dde4ed22e96ddd0141043699d871786b743470137a9d4d9eeb5e0d6a3c82ee700fbe934446c29af861c0c72e751c3f31b8fcb2ed5074a1a3758883d40e8c53e036b02da92ef0b6c5d0a8"
            },
            "sequence": 4294967295,
            "txid": "5776c4ca1906b1e52c51f8ae71e23a77034c7eddbe244000bb706cd368fc32e8",
            "vout": 0
        }
    ],
    "vout": [
        {
            "n": 0,
            "scriptPubKey": {
                "addresses": [
                    "1FjiovQdhWcpbK5QxkryAgWDZY8MpMzP62"
                ],
                "asm": "OP_DUP OP_HASH160 a1a621ba261e3d9b3986af158b1274a21c2b7a27 OP_EQUALVERIFY OP_CHECKSIG",
                "hex": "76a914a1a621ba261e3d9b3986af158b1274a21c2b7a2788ac",
                "reqSigs": 1,
                "type": "pubkeyhash"
            },
            "value": 54.784
        },
        {
            "n": 1,
            "scriptPubKey": {
                "addresses": [
                    "1CC1h2UPUvdHsXs1kVvPsoRXGFomtri78X"
                ],
                "asm": "OP_DUP OP_HASH160 7abe8906e9293e41e0db37df516ad7088eb1c6a9 OP_EQUALVERIFY OP_CHECKSIG",
                "hex": "76a9147abe8906e9293e41e0db37df516ad7088eb1c6a988ac",
                "reqSigs": 1,
                "type": "pubkeyhash"
            },
            "value": 0.1
        }
    ]
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

func TestRawTransactionParsing(t *testing.T) {
	tx := RawTransaction{}
	err := json.Unmarshal([]byte(exampleRaw), &tx)
	if err != nil {
		t.Fatalf("Error parsing transaction: %v", err)
	}

	// TODO:  Add assertions here.

	// Should probably confirm these do something sensible.
	tx.BlockTime()
	tx.TransactionTime()
}
