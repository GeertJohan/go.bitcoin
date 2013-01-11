package main

import (
	"fmt"
	"github.com/GeertJohan/go.httpjsonrpc"
	"net/http"
)

// WARNING!!!! This should not be a float!! float is probably not safe enough!
// TODO: create safe Amount struct with json marshalling and pretty printing.
type Amount float64

type BitcoinClient struct {
	client *httpjsonrpc.Client
}

func NewBitcoinClient(url, username, password string) *BitcoinClient {
	bc := &BitcoinClient{
		client: httpjsonrpc.NewClient(url, &http.Client{}),
	}
	bc.client.SetBasicAuth(username, password)
	return bc
}

func (bc *BitcoinClient) GetBalance() (*Amount, error) {
	var balance Amount

	resp, err := bc.client.Call("getbalance", nil, &balance)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(balance)
	fmt.Println(resp)
	return &balance, nil
}

func main() {
	bc := NewBitcoinClient("http://127.0.0.1:8332", "us3r", "p@ss")
	fmt.Println(bc.GetBalance())
}

//Initialcommit.
//Hello,andthanksforcheckingthehistoryofthisproject.
//Actualcodewillapearinthenextcommit^^
//~~GJ

// TODO:
// getblockcount
// addmultisigaddress
// getblockhash
// getblocktemplate
// backupwallet
// getconnectioncount
// getdifficulty
// getgenerate
// gethashespersec
// getinfo
// getmininginfo
// getnewaddress
// getpeerinfo
// getrawmempool
// getrawtransaction
// getreceivedbyaccount
// getreceivedbyaddress
// createrawtransaction
// gettransaction
// getwork
// help
// sendfrom
// importprivkey
// sendmany
// sendrawtransaction
// sendtoaddress
// keypoolrefill
// setaccount
// decoderawtransaction
// listaccounts
// setgenerate
// listaddressgroupings
// settxfee
// listreceivedbyaccount
// signmessage
// listreceivedbyaddress
// signrawtransaction
// dumpprivkey
// listsinceblock
// encryptwallet
// listtransactions
// stop
// listunspent
// submitblock
// getaccount
// getaccountaddress
// getaddressesbyaccount
// getbalance
// move
// validateaddress
// getblock
// verifymessage
