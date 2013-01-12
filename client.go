package bitcoin

import (
	"fmt"
	"github.com/GeertJohan/go.httpjsonrpc"
)

// The BitcoindClient allows you to easily retrieve information from your bitcoind instance/server. 
type BitcoindClient struct {
	client *httpjsonrpc.Client
}

// Create a new BitcoindClient by http URL (e.g. http://127.0.0.1:8332), username and password.
func NewBitcoindClient(url, username, password string) *BitcoindClient {
	bc := &BitcoindClient{
		client: httpjsonrpc.NewClient(url, nil),
	}
	bc.client.SetBasicAuth(username, password)
	return bc
}

func (bc *BitcoindClient) GetBalance() (*Amount, error) {
	var am = &Amount{}

	_, err := bc.client.Call("getbalance", nil, am)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return am, nil
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
