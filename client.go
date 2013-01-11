package bitcoin

import (
	"fmt"
	"github.com/GeertJohan/go.httpjsonrpc"
	"net/http"
)

type BitcoindClient struct {
	client *httpjsonrpc.Client
}

func NewBitcoindClient(url, username, password string) *BitcoindClient {
	bc := &BitcoindClient{
		client: httpjsonrpc.NewClient(url, &http.Client{}),
	}
	bc.client.SetBasicAuth(username, password)
	return bc
}

func (bc *BitcoindClient) GetBalance() (*Amount, error) {
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
