package bitcoin

import (
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

// Result from an Info call.
type Info struct {
	Balance         Amount
	Blocks          int
	Connections     int
	Difficulty      float64
	Errors          string
	KeyPoolOldest   uint32
	KeyPoolSize     int
	PayTxFee        Amount
	ProtocolVersion int
	Proxy           string
	Testnet         bool
	Version         int
	WalletVersion   int
}

func (bc *BitcoindClient) GetInfo() (Info, error) {
	rv := Info{}
	_, err := bc.client.Call("getinfo", nil, &rv)
	return rv, err
}

func (bc *BitcoindClient) GetBalance(a ...string) (Amount, error) {
	var am Amount
	_, err := bc.client.Call("getbalance", a, &am)
	return am, err
}

func (bc *BitcoindClient) ListAccounts() (map[string]Amount, error) {
	m := map[string]Amount{}
	_, err := bc.client.Call("listaccounts", nil, &m)
	return m, err
}

func (bc *BitcoindClient) GetAccount(name string) (string, error) {
	var rv string
	_, err := bc.client.Call("getaccount", []string{name}, &rv)
	return rv, err
}

func (bc *BitcoindClient) GetAccountAddress(name string) (string, error) {
	var rv string
	_, err := bc.client.Call("getaccountaddress", []string{name}, &rv)
	return rv, err
}

func (bc *BitcoindClient) SendToAddress(addr string, amt Amount,
	comment, commentto string) (string, error) {
	var rv string
	_, err := bc.client.Call("sendtoaddress",
		[]interface{}{addr, amt, comment, commentto}, &rv)
	return rv, err
}

func (bc *BitcoindClient) SendFrom(myact, addr string, amt Amount,
	minconf int, comment, commentto string) (string, error) {
	var rv string
	_, err := bc.client.Call("sendfrom",
		[]interface{}{myact, addr, amt, minconf, comment, commentto}, &rv)
	return rv, err
}

func (bc *BitcoindClient) GetTransaction(txid string) (Transaction, error) {
	rv := Transaction{}
	_, err := bc.client.Call("gettransaction", []string{txid}, &rv)
	return rv, err
}

type AddressInfo struct {
	Isvalid      bool
	Isscript     bool
	Address      string
	Iscompressed bool
	Account      string
	Ismine       bool
	Pubkey       string
}

func (bc *BitcoindClient) ValidateAddress(addr string) (AddressInfo, error) {
	rv := AddressInfo{}
	_, err := bc.client.Call("validateaddress", []string{addr}, &rv)
	return rv, err
}

func (bc *BitcoindClient) GetRawTransaction(txn string) (RawTransaction, error) {
	var rv RawTransaction
	_, err := bc.client.Call("getrawtransaction", []interface{}{txn, 1}, &rv)
	return rv, err
}

func (bc *BitcoindClient) DecodeRawTransaction(txn string) (RawTransaction, error) {
	rv := RawTransaction{}
	_, err := bc.client.Call("decoderawtransaction", []string{txn}, &rv)
	return rv, err
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
// getmininginfo
// getnewaddress
// getpeerinfo
// getrawmempool
// getreceivedbyaccount
// getreceivedbyaddress
// createrawtransaction
// getwork
// help
// importprivkey
// sendmany
// sendrawtransaction
// keypoolrefill
// setaccount
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
// getaddressesbyaccount
// getbalance
// move
// getblock
// verifymessage
