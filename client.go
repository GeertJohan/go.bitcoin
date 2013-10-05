package bitcoin

import (
	"github.com/GeertJohan/go.httpjsonrpc"
	"crypto/x509"
	"crypto/tls"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
)

// The BitcoindClient allows you to easily retrieve information from your bitcoind instance/server.
type BitcoindClient struct {
	client *httpjsonrpc.Client
}

func slurpFile(filename string) []byte {
        f, err := os.Open(filename)
        if (err != nil) { panic (err) }
        defer f.Close()
        contents, err := ioutil.ReadAll(f)
        if (err != nil) { panic (err) }
        return contents
}

func makeClient(certFile string) *http.Client {
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(slurpFile(certFile))
	if (ok == false) { panic (fmt.Sprintf("No certificates found in %s", certFile)) }
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
	}
	client := &http.Client{Transport: tr}
	return client
}

// Create a new BitcoindClient by http URL (e.g. http://127.0.0.1:8332), username and password.
func NewBitcoindClient(url, username, password, certFile string) *BitcoindClient {
	client := makeClient(certFile)
	bc := &BitcoindClient{
		client: httpjsonrpc.NewClient(url, client),
	}
	bc.client.SetBasicAuth(username, password)
	return bc
}

// Result from a BitcoindClient.GetInfo() call.
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

// GetInfo gives info about the bitcoind status and main account balance.
func (bc *BitcoindClient) GetInfo() (Info, error) {
	rv := Info{}
	_, err := bc.client.Call("getinfo", nil, &rv)
	return rv, err
}

// GetBalance gives balance for given accounts. If a is not specified, returns the server's total available balance.
func (bc *BitcoindClient) GetBalance(a ...string) (Amount, error) {
	var am Amount
	_, err := bc.client.Call("getbalance", a, &am)
	return am, err
}

// ListAccounts returns a map with all accounts. Account name as key. Balance as value.
func (bc *BitcoindClient) ListAccounts() (map[string]Amount, error) {
	m := map[string]Amount{}
	_, err := bc.client.Call("listaccounts", nil, &m)
	return m, err
}

// GetAccount returns the account associated with the given address. 
func (bc *BitcoindClient) GetAccount(name string) (string, error) {
	var rv string
	_, err := bc.client.Call("getaccount", []string{name}, &rv)
	return rv, err
}

// GetAccountAddress returns the current bitcoin address for receiving payments to this account. 
func (bc *BitcoindClient) GetAccountAddress(name string) (string, error) {
	var rv string
	_, err := bc.client.Call("getaccountaddress", []string{name}, &rv)
	return rv, err
}

//++ TODO: good description
func (bc *BitcoindClient) SendToAddress(addr string, amt Amount,
	comment, commentto string) (string, error) {
	var rv string
	_, err := bc.client.Call("sendtoaddress",
		[]interface{}{addr, amt, comment, commentto}, &rv)
	return rv, err
}

//++ TODO: good description
func (bc *BitcoindClient) SendFrom(myact, addr string, amt Amount,
	minconf int, comment, commentto string) (string, error) {
	var rv string
	_, err := bc.client.Call("sendfrom",
		[]interface{}{myact, addr, amt, minconf, comment, commentto}, &rv)
	return rv, err
}

// GetTransaction returns Transaction for given transaction id.
func (bc *BitcoindClient) GetTransaction(txid string) (Transaction, error) {
	rv := Transaction{}
	_, err := bc.client.Call("gettransaction", []string{txid}, &rv)
	return rv, err
}

// AddressInfo contains information about a bitcoin address. See BitcoindClient.ValidateAddress().
type AddressInfo struct {
	Isvalid      bool
	Isscript     bool
	Address      string
	Iscompressed bool
	Account      string
	Ismine       bool
	Pubkey       string
}

// ValidateAddress returns information (AddressInfo) on given address string.
func (bc *BitcoindClient) ValidateAddress(addr string) (AddressInfo, error) {
	rv := AddressInfo{}
	_, err := bc.client.Call("validateaddress", []string{addr}, &rv)
	return rv, err
}

// GetRawTransaction returns a RawTransaction instance for given transaction string.
func (bc *BitcoindClient) GetRawTransaction(txn string) (RawTransaction, error) {
	var rv RawTransaction
	_, err := bc.client.Call("getrawtransaction", []interface{}{txn, 1}, &rv)
	return rv, err
}

// DecodeRawTransaction returns a RawTransaction instance for given transaction string.
func (bc *BitcoindClient) DecodeRawTransaction(txn string) (RawTransaction, error) {
	rv := RawTransaction{}
	_, err := bc.client.Call("decoderawtransaction", []string{txn}, &rv)
	return rv, err
}

// ListTransactions retrieves transactions for given account. It allows for pagination using the `count` and `from` int arguments.
func (bc *BitcoindClient) ListTransactions(acct string,
	count int, from int) ([]Transaction, error) {

	rv := []Transaction{}
	_, err := bc.client.Call("listtransactions", []interface{}{acct, count, from}, &rv)
	return rv, err
}

// GetAddressesByAccount returns the addresses for given account string.
func (bc *BitcoindClient) GetAddressesByAccount(acct string) ([]string, error) {
	rv := []string{}
	_, err := bc.client.Call("getaddressesbyaccount", []interface{}{acct}, &rv)
	return rv, err
}

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
// stop
// listunspent
// submitblock
// getbalance
// move
// getblock
// verifymessage
