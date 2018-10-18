package client

import (
	"github.com/chainlibs/gobtclib/futures"
)

/*
Description:
AbandonTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See AbandonTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) AbandonTransactionAsync(txid string) futures.FutureResult {
	cmd := NewCommand("abandontransaction", txid)
	return c.sendCmd(cmd)
}

/*
Description:
Mark in-wallet transaction <txid> as abandoned
This will mark this transaction and all its in-wallet descendants as abandoned which will allow
for their inputs to be respent.  It can be used to replace "stuck" or evicted transactions.
It only works on transactions which are not included in a block and are not currently in the mempool.
It has no effect on transactions which are already abandoned.
 * Author: architect.bian
 * Date: 2018/10/15 13:29
 */
func (c *Client) AbandonTransaction(txid string) (error) {
	_, err := c.AbandonTransactionAsync(txid).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
AbortRescanAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See AbortRescan for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) AbortRescanAsync() futures.FutureResult {
	cmd := NewCommand("abortrescan")
	return c.sendCmd(cmd)
}

/*
Description:
Stops current wallet rescan triggered by an RPC call, e.g. by an importprivkey call.
 * Author: architect.bian
 * Date: 2018/10/15 13:34
 */
func (c *Client) AbortRescan() (error) {
	_, err := c.AbortRescanAsync().Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
AddMultiSigAddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See AddMultiSigAddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) AddMultiSigAddressAsync(nrequired int32, keys []string, label *string, addressType *string) futures.FutureResult {
	cmd := NewCommand("addmultisigaddress", nrequired)
	cmd.AddArgs(keys)
	if addressType != nil {
		cmd.AddArgs(*label)
		cmd.AddArgs(*addressType)
	} else if label != nil {
		cmd.AddArgs(*label)
	}
	return c.sendCmd(cmd)
}

/*
Description: 
Add a nrequired-to-sign multisignature address to the wallet. Requires a new wallet backup.
Each key is a Bitcoin address or hex-encoded public key.
This functionality is only intended for use with non-watchonly addresses.
See `importaddress` for watchonly p2sh address support.
 * Author: architect.bian
 * Date: 2018/10/15 13:38
 */
func (c *Client) AddMultiSigAddress(nrequired int32, keys []string) (*interface{}, error) {
	return c.AddMultiSigAddressAsync(nrequired, keys, nil, nil).Receive()
}

/*
Description:
AddMultiSigAddressEntire Add a nrequired-to-sign multisignature address to the wallet. Requires a new wallet backup.
Each key is a Bitcoin address or hex-encoded public key.
This functionality is only intended for use with non-watchonly addresses.
See `importaddress` for watchonly p2sh address support.
'label' is specified, assign address to that label.
 * Author: architect.bian
 * Date: 2018/10/15 13:38
 */
func (c *Client) AddMultiSigAddressEntire(nrequired int32, keys []string, label string, addressType string) (*interface{}, error) {
	return c.AddMultiSigAddressAsync(nrequired, keys, &label, &addressType).Receive()
}

/*
Description:
BackupWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See BackupWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) BackupWalletAsync(destination string) futures.FutureResult {
	cmd := NewCommand("backupwallet", destination)
	return c.sendCmd(cmd)
}

/*
Description:
Safely copies current wallet file to destination, which can be a directory or a path with filename.
 * Author: architect.bian
 * Date: 2018/10/15 14:31
 */
func (c *Client) BackupWallet(destination string) (error) {
	_, err := c.BackupWalletAsync(destination).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
BumpFeeAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See BumpFee for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) BumpFeeAsync(txid string, options map[string]interface{}) futures.FutureResult {
	cmd := NewCommand("bumpfee", txid, options)
	return c.sendCmd(cmd)
}

/*
Description: 

 * Author: architect.bian
 * Date: 2018/10/15 14:33
 */
func (c *Client) BumpFee(txid string, options map[string]interface{}) (*interface{}, error) {
	return c.BumpFeeAsync(txid, options).Receive()
}

/*
Description:
CreateWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CreateWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) CreateWalletAsync(walletName string, disablePrivateKeys bool) futures.FutureResult {
	cmd := NewCommand("createwallet", walletName, disablePrivateKeys)
	return c.sendCmd(cmd)
}

/*
Description:
Creates and loads a new wallet.
 * Author: architect.bian
 * Date: 2018/10/15 14:38
 */
func (c *Client) CreateWallet(walletName string, disablePrivateKeys bool) (*interface{}, error) {
	return c.CreateWalletAsync(walletName, disablePrivateKeys).Receive()
}

/*
Description:
DumpPrivkeyAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DumpPrivkey for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) DumpPrivkeyAsync(address string) futures.FutureString {
	cmd := NewCommand("dumpprivkey", address)
	return c.sendCmd(cmd)
}

/*
Description:
Reveals the private key corresponding to 'address'.
Then the importprivkey can be used with this output
 * Author: architect.bian
 * Date: 2018/10/15 14:41
 */
func (c *Client) DumpPrivkey(address string) (*string, error) {
	return c.DumpPrivkeyAsync(address).Receive()
}

/*
Description:
DumpWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DumpWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) DumpWalletAsync(filename string) futures.FutureResult {
	cmd := NewCommand("dumpwallet", filename)
	return c.sendCmd(cmd)
}

/*
Description:
Dumps all wallet keys in a human-readable format to a server-side file. This does not allow overwriting existing files.
Imported scripts are included in the dumpfile, but corresponding BIP173 addresses, etc. may not be added automatically by importwallet.
Note that if your wallet contains keys which are not derived from your HD seed (e.g. imported keys), these are not covered by
only backing up the seed itself, and must be backed up too (e.g. ensure you back up the whole dumpfile).
 * Author: architect.bian
 * Date: 2018/10/15 14:44
 */
func (c *Client) DumpWallet(filename string) (*interface{}, error) {
	return c.DumpWalletAsync(filename).Receive()
}

/*
Description:
EncryptWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See EncryptWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) EncryptWalletAsync(passphrase string) futures.FutureResult {
	cmd := NewCommand("encryptwallet", passphrase)
	return c.sendCmd(cmd)
}

/*
Description:
Encrypts the wallet with 'passphrase'. This is for first time encryption.
After this, any calls that interact with private keys such as sending or signing
will require the passphrase to be set prior the making these calls.
Use the walletpassphrase call for this, and then walletlock call.
If the wallet is already encrypted, use the walletpassphrasechange call.
Note that this will shutdown the server.
 * Author: architect.bian
 * Date: 2018/10/15 14:46
 */
func (c *Client) EncryptWallet(passphrase string) (error) {
	_, err := c.EncryptWalletAsync(passphrase).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
GetAddressesByLabelAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetAddressesByLabel for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetAddressesByLabelAsync(label string) futures.FutureResult {
	cmd := NewCommand("getaddressesbylabel", label)
	return c.sendCmd(cmd)
}

/*
Description: 
Returns the list of addresses assigned the specified label.
 * Author: architect.bian
 * Date: 2018/10/15 14:53
 */
func (c *Client) GetAddressesByLabel(label string) (*interface{}, error) {
	return c.GetAddressesByLabelAsync(label).Receive()
}

/*
Description:
GetAddressInfoAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetAddressInfo for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetAddressInfoAsync(address string) futures.FutureResult {
	cmd := NewCommand("getaddressinfo", address)
	return c.sendCmd(cmd)
}

/*
Description:
Return information about the given bitcoin address. Some information requires the address
to be in the wallet.
 * Author: architect.bian
 * Date: 2018/10/15 14:55
 */
func (c *Client) GetAddressInfo(address string) (*interface{}, error) { //TODO results.GetAddressInfoResult
	return c.GetAddressInfoAsync(address).Receive()
}

/*
Description:
GetBalanceAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetBalance for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetBalanceAsync(dummy string, minConfirm int32, includeWatchonly bool) futures.FutureFloat64 {
	cmd := NewCommand("getbalance", dummy, minConfirm, includeWatchonly)
	return c.sendCmd(cmd)
}

/*
Description: 
Returns the total available balance.
 * Author: architect.bian
 * Date: 2018/10/15 15:03
 */
func (c *Client) GetBalance() (float64, error) {
	return c.GetBalanceEntire("*", 0, false)
}

/*
Description:
Returns the total available balance.
The available balance is what the wallet considers currently spendable, and is
thus affected by options which limit spendability such as -spendzeroconfchange.
 * Author: architect.bian
 * Date: 2018/10/16 16:03
 */
func (c *Client) GetBalanceEntire(dummy string, minConfirm int32, includeWatchonly bool) (float64, error) {
	return c.GetBalanceAsync(dummy, minConfirm, includeWatchonly).Receive()
}

/*
Description:
GetNewAddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetNewAddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetNewAddressAsync(label *string, addressType *string) futures.FutureString {
	cmd := NewCommand("getnewaddress")
	if addressType != nil {
		cmd.AddArgs(*label)
		cmd.AddArgs(*addressType)
	} else if label != nil {
		cmd.AddArgs(*label)
	}
	return c.sendCmd(cmd)
}

/*
Description: 
Returns a new Bitcoin address for receiving payments.
If 'label' is specified, it is added to the address book
so payments received with the address will be associated with 'label'.
 * Author: architect.bian
 * Date: 2018/10/15 15:06
 */
func (c *Client) GetNewAddress() (*string, error) {
	return c.GetNewAddressAsync(nil, nil).Receive()
}

/*
Description:
GetNewAddressEntire Returns a new Bitcoin address for receiving payments.
If 'label' is specified, it is added to the address book
so payments received with the address will be associated with 'label'.
 * Author: architect.bian
 * Date: 2018/10/15 15:06
 */
func (c *Client) GetNewAddressEntire(label string, addressType string) (*string, error) {
	return c.GetNewAddressAsync(&label, &addressType).Receive()
}

/*
Description:
GetRawChangeAddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetRawChangeAddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetRawChangeAddressAsync(addressType *string) futures.FutureString {
	cmd := NewCommand("getrawchangeaddress")
	if addressType != nil {
		cmd.AddArgs(*addressType)
	}
	return c.sendCmd(cmd)
}

/*
Description: 
Returns a new Bitcoin address, for receiving change.
This is for use with raw transactions, NOT normal use.
 * Author: architect.bian
 * Date: 2018/10/15 15:15
 */
func (c *Client) GetRawChangeAddress() (*string, error) {
	return c.GetRawChangeAddressAsync(nil).Receive()
}

/*
Description:
Returns a new Bitcoin address, for receiving change.
This is for use with raw transactions, NOT normal use.
 * Author: architect.bian
 * Date: 2018/10/15 15:15
 */
func (c *Client) GetRawChangeAddressEntire(addressType string) (*string, error) {
	return c.GetRawChangeAddressAsync(&addressType).Receive()
}

/*
Description:
GetReceivedByAddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetReceivedByAddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetReceivedByAddressAsync(address string, minconfirms int32) futures.FutureFloat64 {
	cmd := NewCommand("getreceivedbyaddress", address, minconfirms)
	return c.sendCmd(cmd)
}

/*
Description:
Returns the total amount received by the given address in transactions with
at least minconfirms confirmations.
 * Author: architect.bian
 * Date: 2018/10/15 15:20
 */
func (c *Client) GetReceivedByAddress(address string, minconfirms int32) (float64, error) {
	return c.GetReceivedByAddressAsync(address, minconfirms).Receive()
}

/*
Description:
GetTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetTransactionAsync(txid string, includeWatchonly bool) futures.FutureResult {
	cmd := NewCommand("gettransaction", txid, includeWatchonly)
	return c.sendCmd(cmd)
}

/*
Description: 
Get detailed information about in-wallet transaction <txid>
 * Author: architect.bian
 * Date: 2018/10/15 15:21
 */
func (c *Client) GetTransaction(txid string, includeWatchonly bool) (*interface{}, error) { //TODO results.GetTransactionResult
	return c.GetTransactionAsync(txid, includeWatchonly).Receive()
}

/*
Description:
GetUnconfirmedBalanceAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetUnconfirmedBalance for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetUnconfirmedBalanceAsync() futures.FutureFloat64 {
	cmd := NewCommand("getunconfirmedbalance")
	return c.sendCmd(cmd)
}

/*
Description: 
Returns the server's total unconfirmed balance
 * Author: architect.bian
 * Date: 2018/10/15 15:40
 */
func (c *Client) GetUnconfirmedBalance() (float64, error) {
	return c.GetUnconfirmedBalanceAsync().Receive()
}

/*
Description:
GetWalletInfoAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetWalletInfo for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) GetWalletInfoAsync() futures.FutureResult {
	cmd := NewCommand("getwalletinfo")
	return c.sendCmd(cmd)
}

/*
Description:
Returns an object containing various wallet state info.
 * Author: architect.bian
 * Date: 2018/10/15 15:41
 */
func (c *Client) GetWalletInfo() (*interface{}, error) { //TODO results.GetWalletInfo
	return c.GetWalletInfoAsync().Receive()
}

/*
Description:
ImportaddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See Importaddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ImportaddressAsync(address string, label string, rescan bool, p2sh bool) futures.FutureString {
	cmd := NewCommand("importaddress", address, label, rescan, p2sh)
	return c.sendCmd(cmd)
}

/*
Description: 
Adds an address or script (in hex) that can be watched as if it were in your wallet but cannot be used to spend.
Requires a new wallet backup.
Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
may report that the imported address exists but related transactions are still missing, leading to temporarily
incorrect/bogus balances and unspent outputs until rescan completes.
If you have the full public key, you should call importpubkey instead of this.

Note: If you import a non-standard raw script in hex form, outputs sending to it will be treated
as change, and not show up in many RPCs.
 * Author: architect.bian
 * Date: 2018/10/15 15:42
 */
func (c *Client) Importaddress(address string) (*string, error) {
	return c.ImportaddressAsync(address, "", true, false).Receive()
}

/*
Description:
Adds an address or script (in hex) that can be watched as if it were in your wallet but cannot be used to spend. Requires a new wallet backup.
 * Author: architect.bian
 * Date: 2018/10/15 15:42
 */
func (c *Client) ImportaddressEntire(address string, label string, rescan bool, p2sh bool) (*string, error) {
	return c.ImportaddressAsync(address, label, rescan, p2sh).Receive()
}

/*
Description:
ImportMultiAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ImportMulti for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ImportMultiAsync(requests interface{}, option interface{}) futures.FutureResult {
	cmd := NewCommand("importmulti", requests, option)
	return c.sendCmd(cmd)
}

/*
Description:
TODO https://bitcoincore.org/en/doc/0.17.0/rpc/wallet/importmulti/
 * Author: architect.bian
 * Date: 2018/10/15 15:49
 */
func (c *Client) ImportMulti(requests interface{}, option interface{}) (*interface{}, error) {
	return c.ImportMultiAsync(requests, option).Receive()
}

/*
Description:
ImportPrivkeyAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ImportPrivkey for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ImportPrivkeyAsync(privkey string, label string, rescan bool) futures.FutureResult {
	cmd := NewCommand("importprivkey", privkey, label, rescan)
	return c.sendCmd(cmd)
}

/*
Description:
Adds a private key (as returned by dumpprivkey) to your wallet. Requires a new wallet backup.
Hint: use importmulti to import more than one private key.
 * Author: architect.bian
 * Date: 2018/10/15 15:54
 */
func (c *Client) ImportPrivkey(privkey string) (error) {
	return c.ImportPrivkeyEntire(privkey, "", true)
}

/*
Description:
Adds a private key (as returned by dumpprivkey) to your wallet. Requires a new wallet backup.
Hint: use importmulti to import more than one private key.
 * Author: architect.bian
 * Date: 2018/10/15 15:54
 */
func (c *Client) ImportPrivkeyEntire(privkey string, label string, rescan bool) (error) {
	_, err := c.ImportPrivkeyAsync(privkey, label, rescan).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
ImportPrunedFundsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ImportPrunedFunds for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ImportPrunedFundsAsync(rawTransaction []byte, txOutProof []byte) futures.FutureResult {
	cmd := NewCommand("importprunedfunds", string(rawTransaction), string(txOutProof))
	return c.sendCmd(cmd)
}

/*
Description:
Imports funds without rescan. Corresponding address or script must previously be included in wallet.
Aimed towards pruned wallets. The end-user is responsible to import additional transactions that subsequently
spend the imported outputs or rescan after the point in the blockchain the transaction is included.
 * Author: architect.bian
 * Date: 2018/10/15 15:58
 */
func (c *Client) ImportPrunedFunds(rawTransaction []byte, txOutProof []byte) (error) {
	_, err := c.ImportPrunedFundsAsync(rawTransaction, txOutProof).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
ImportPubkeyAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ImportPubkey for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ImportPubkeyAsync(pubkey string, label string, rescan bool) futures.FutureResult {
	cmd := NewCommand("importpubkey", pubkey, label, rescan)
	return c.sendCmd(cmd)
}

/*
Description:
Adds a public key (in hex) that can be watched as if it were in your wallet but cannot be used to spend.
Requires a new wallet backup.
 * Author: architect.bian
 * Date: 2018/10/15 16:03
 */
func (c *Client) ImportPubkey(pubkey string) (error) {
	return c.ImportPubkeyEntire(pubkey, "", true)
}

/*
Description:
Adds a public key (in hex) that can be watched as if it were in your wallet but cannot be used to spend.
Requires a new wallet backup.
 * Author: architect.bian
 * Date: 2018/10/15 16:03
 */
func (c *Client) ImportPubkeyEntire(pubkey string, label string, rescan bool) (error) {
	_, err := c.ImportPubkeyAsync(pubkey, label, rescan).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
ImportWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ImportWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ImportWalletAsync(filename string) futures.FutureResult {
	cmd := NewCommand("importwallet")
	return c.sendCmd(cmd)
}

/*
Description:
Imports keys from a wallet dump file (see dumpwallet). Requires a new wallet backup to include imported keys.
 * Author: architect.bian
 * Date: 2018/10/15 16:07
 */
func (c *Client) ImportWallet(filename string) (error) {
	_, err := c.ImportWalletAsync(filename).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
KeypoolRefillAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See KeypoolRefill for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) KeypoolRefillAsync(newsize int32) futures.FutureResult {
	cmd := NewCommand("keypoolrefill", newsize)
	return c.sendCmd(cmd)
}

/*
Description:
Fills the keypool.
newsize     (numeric, optional, default=100) The new keypool size
 * Author: architect.bian
 * Date: 2018/10/15 16:15
 */
func (c *Client) KeypoolRefill(newsize int32) (error) {
	_, err := c.KeypoolRefillAsync(newsize).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
ListAddressGroupingsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListAddressGroupings for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListAddressGroupingsAsync() futures.FutureResult {
	cmd := NewCommand("listaddressgroupings")
	return c.sendCmd(cmd)
}

/*
Description: 
Lists groups of addresses which have had their common ownership
made public by common use as inputs or as the resulting change
in past transactions
 * Author: architect.bian
 * Date: 2018/10/15 16:30
 */
func (c *Client) ListAddressGroupings() (*interface{}, error) {
	return c.ListAddressGroupingsAsync().Receive()
}

/*
Description:
ListLabelsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListLabels for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListLabelsAsync(purpose string) futures.FutureResult {
	cmd := NewCommand("listlabels", purpose)
	return c.sendCmd(cmd)
}

/*
Description:
Returns the list of all labels, or labels that are assigned to addresses with a specific purpose.
purpose An empty string is the same as not providing this argument.
 * Author: architect.bian
 * Date: 2018/10/15 16:31
 */
func (c *Client) ListLabels(purpose string) (*interface{}, error) {
	return c.ListLabelsAsync(purpose).Receive()
}

/*
Description:
ListLockUnspentAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListLockUnspent for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListLockUnspentAsync() futures.FutureResult {
	cmd := NewCommand("listlockunspent")
	return c.sendCmd(cmd)
}

/*
Description:
Returns list of temporarily unspendable outputs.
See the lockunspent call to lock and unlock transactions for spending.
 * Author: architect.bian
 * Date: 2018/10/15 16:32
 */
func (c *Client) ListLockUnspent() (*interface{}, error) {
	return c.ListLockUnspentAsync().Receive()
}

/*
Description:
ListReceivedByAddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListReceivedByAddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListReceivedByAddressAsync(minconfirms int32, includeEmpty bool, includeWatchonly bool,
	addressFilter string)futures.FutureResult {
	cmd := NewCommand("listreceivedbyaddress", minconfirms, includeEmpty, includeWatchonly)
	if addressFilter != "" {
		cmd.AddArgs(addressFilter)
	}
	return c.sendCmd(cmd)
}

/*
Description:
List balances by receiving address.
 * Author: architect.bian
 * Date: 2018/10/15 16:40
 */
func (c *Client) ListReceivedByAddress(minconfirms int32, includeEmpty bool, includeWatchonly bool, addressFilter string) (
	*interface{}, error) { //TODO results.ListReceivedByAddressResult
	return c.ListReceivedByAddressAsync(minconfirms, includeEmpty, includeWatchonly, addressFilter).Receive()
}

/*
Description:
ListSinceBlockAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListSinceBlock for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListSinceBlockAsync(blockhash string, targetConfirmations int32, includeWatchonly bool, includeRemoved bool) futures.FutureResult {
	cmd := NewCommand("listsinceblock", blockhash, targetConfirmations, includeWatchonly, includeRemoved)
	return c.sendCmd(cmd)
}

/*
Description:
Get all transactions in blocks since block [blockhash], or all transactions if omitted.
If "blockhash" is no longer a part of the main chain, transactions from the fork point onward are included.
Additionally, if include_removed is set, transactions affecting the wallet which were removed are returned
in the "removed" array.
 * Author: architect.bian
 * Date: 2018/10/15 16:42
 */
func (c *Client) ListSinceBlock(blockhash string, targetConfirmations int32, includeWatchonly bool, includeRemoved bool) (
	*interface{}, error) { //TODO results.Transaction
	return c.ListSinceBlockAsync(blockhash, targetConfirmations, includeWatchonly, includeRemoved).Receive()
}

/*
Description:
ListTransactionsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListTransactions for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListTransactionsAsync(dummy string, count int32, skip int32, includeWatchonly bool) futures.FutureResult {
	cmd := NewCommand("listtransactions", dummy, count, skip, includeWatchonly)
	return c.sendCmd(cmd)
}

/*
Description:
Returns up to 'count' most recent transactions skipping the first 'from' transactions for account 'account'.
Note that the "account" argument and "otheraccount" return value have been removed in V0.17. To use this RPC with
an "account" argument, restart bitcoind with -deprecatedrpc=accounts
 * Author: architect.bian
 * Date: 2018/10/15 16:46
 */
func (c *Client) ListTransactions() (*interface{}, error) {
	return c.ListTransactionsEntire("*", 10, 0, false)
}

/*
Description:
Returns up to 'count' most recent transactions skipping the first 'from' transactions for account 'account'.
Note that the "account" argument and "otheraccount" return value have been removed in V0.17. To use this RPC with
an "account" argument, restart bitcoind with -deprecatedrpc=accounts
 * Author: architect.bian
 * Date: 2018/10/15 16:46
 */
func (c *Client) ListTransactionsEntire(dummy string, count int32, skip int32, includeWatchonly bool) (*interface{}, error) {
	return c.ListTransactionsAsync(dummy, count, skip, includeWatchonly).Receive()
}

/*
Description:
ListUnspentAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListUnspent for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListUnspentAsync(minconfirm int32, maxconfirm int32, addresses []string, includeUnsafe bool,
	queryOptions map[string]interface{}) futures.FutureResult {
	cmd := NewCommand("listunspent", minconfirm, maxconfirm, addresses, includeUnsafe, queryOptions)
	return c.sendCmd(cmd)
}

/*
Description: 
Returns array of unspent transaction outputs
with between minconf and maxconf (inclusive) confirmations.
Optionally filter to only include txouts paid to specified addresses.
 * Author: architect.bian
 * Date: 2018/10/15 16:49
 */
func (c *Client) ListUnspent() (*interface{}, error) {
	return c.ListUnspentEntire(0, 9999999, nil, true, nil)
}

/*
Description:
Returns array of unspent transaction outputs
with between minconf and maxconf (inclusive) confirmations.
Optionally filter to only include txouts paid to specified addresses.
 * Author: architect.bian
 * Date: 2018/10/15 16:49
 */
func (c *Client) ListUnspentEntire(minconfirm int32, maxconfirm int32, addresses []string, includeUnsafe bool,
	queryOptions map[string]interface{}) (*interface{}, error) {
	return c.ListUnspentAsync(minconfirm, maxconfirm, addresses, includeUnsafe, queryOptions).Receive()
}

/*
Description:
ListWalletsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ListWallets for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) ListWalletsAsync() futures.FutureStringArray {
	cmd := NewCommand("listwallets")
	return c.sendCmd(cmd)
}

/*
Description:
Returns a list of currently loaded wallets.
For full information on the wallet, use "getwalletinfo"
 * Author: architect.bian
 * Date: 2018/10/15 16:54
 */
func (c *Client) ListWallets() (*[]string, error) {
	return c.ListWalletsAsync().Receive()
}

/*
Description:
LoadWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See LoadWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) LoadWalletAsync(filename string) futures.FutureResult {
	cmd := NewCommand("loadwallet", filename)
	return c.sendCmd(cmd)
}

/*
Description: 
Loads a wallet from a wallet file or directory.
Note that all wallet command-line options used when starting bitcoind will be
applied to the new wallet (eg -zapwallettxes, upgradewallet, rescan, etc).
 * Author: architect.bian
 * Date: 2018/10/15 17:07
 */
func (c *Client) LoadWallet(filename string) (*interface{}, error) {
	return c.LoadWalletAsync(filename).Receive()
}

/*
Description:
LockUnspentAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See LockUnspent for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) LockUnspentAsync(unlock bool, transactions []map[string]interface{}) futures.FutureBool {
	cmd := NewCommand("lockunspent", unlock, transactions)
	return c.sendCmd(cmd)
}

/*
Description: 
Updates list of temporarily unspendable outputs.
Temporarily lock (unlock=false) or unlock (unlock=true) specified transaction outputs.
If no transaction outputs are specified when unlocking then all current locked transaction outputs are unlocked.
A locked transaction output will not be chosen by automatic coin selection, when spending bitcoins.
Locks are stored in memory only. Nodes start with zero locked outputs, and the locked output list
is always cleared (by virtue of process exit) when a node stops or fails.
Also see the listunspent call
 * Author: architect.bian
 * Date: 2018/10/15 17:12
 */
func (c *Client) LockUnspent(unlock bool, transactions []map[string]interface{}) (bool, error) {
	return c.LockUnspentAsync(unlock, transactions).Receive()
}

/*
Description:
RemovePrunedFundsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See RemovePrunedFunds for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) RemovePrunedFundsAsync(txid string) futures.FutureResult {
	cmd := NewCommand("removeprunedfunds", txid)
	return c.sendCmd(cmd)
}

/*
Description: 
Deletes the specified transaction from the wallet. Meant for use with pruned wallets and as a companion to
importprunedfunds. This will affect wallet balances.
 * Author: architect.bian
 * Date: 2018/10/15 17:22
 */
func (c *Client) RemovePrunedFunds(txid string) (error) {
	_, err := c.RemovePrunedFundsAsync(txid).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
RescanBlockChainAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See RescanBlockChain for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) RescanBlockChainAsync(startHeight int32, stopHeight int32) futures.FutureResult {
	cmd := NewCommand("rescanblockchain")
	return c.sendCmd(cmd)
}

/*
Description:
Rescan the local blockchain for wallet related transactions.
 * Author: architect.bian
 * Date: 2018/10/15 17:25
 */
func (c *Client) RescanBlockChain() (*interface{}, error) {
	cmd := NewCommand("rescanblockchain")
	return futures.FutureResult(c.sendCmd(cmd)).Receive()
}

/*
Description:
Rescan the local blockchain for wallet related transactions.
 * Author: architect.bian
 * Date: 2018/10/15 17:25
 */
func (c *Client) RescanBlockChainEntire(startHeight int32, stopHeight int32) (*interface{}, error) {
	return c.RescanBlockChainAsync(startHeight, stopHeight).Receive()
}

/*
Description:
SendManyAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SendMany for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) SendManyAsync(dummy string, amounts map[string]interface{}, minConfirm int32, comment string,
	subtractfeefrom []string, replaceable bool, confTarget int32, estimateMode string) futures.FutureString {
	cmd := NewCommand("sendmany")
	return c.sendCmd(cmd)
}

/*
Description:
sendmany "" {"address":amount,...} ( minconf "comment" ["address",...] replaceable conf_target "estimate_mode")

Send multiple times. Amounts are double-precision floating point numbers.
Note that the "fromaccount" argument has been removed in V0.17. To use this RPC with a "fromaccount" argument, restart
bitcoind with -deprecatedrpc=accounts
 * Author: architect.bian
 * Date: 2018/10/15 17:53
 */
func (c *Client) SendMany(amounts map[string]interface{}) (*string, error) {
	cmd := NewCommand("sendmany", "", amounts)
	return futures.FutureString(c.sendCmd(cmd)).Receive()
}

/*
Description: 
SendManyEntire Send multiple times. Amounts are double-precision floating point numbers.
Note that the "fromaccount" argument has been removed in V0.17.
 * Author: architect.bian
 * Date: 2018/10/15 17:54
 */
func (c *Client) SendManyEntire(dummy string, amounts map[string]interface{}, minConfirm int32, comment string,
	subtractfeefrom []string, replaceable bool, confTarget int32, estimateMode string) (*string, error) {
	return c.SendManyAsync(dummy, amounts, minConfirm, comment, subtractfeefrom, replaceable, confTarget, estimateMode).Receive()
}

/*
Description:
SendToAddressAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SendToAddress for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) SendToAddressAsync(address string, amount float64, comment string, commentTo string,
	subtractfeefromamount bool, replaceable bool, confTarget int32, estimateMode string) futures.FutureString {
	cmd := NewCommand("sendtoaddress", address, amount, comment, commentTo, subtractfeefromamount, replaceable,
		confTarget, estimateMode)
	return c.sendCmd(cmd)
}

/*
Description: 
sendtoaddress "address" amount ( "comment" "comment_to" subtractfeefromamount replaceable conf_target "estimate_mode")

Send an amount to a given address.
 * Author: architect.bian
 * Date: 2018/10/15 18:07
 */
func (c *Client) SendToAddress(address string, amount float64) (*string, error) {
	cmd := NewCommand("sendtoaddress", address, amount)
	return futures.FutureString(c.sendCmd(cmd)).Receive()
}

/*
Description: 
Send an amount to a given address.
 * Author: architect.bian
 * Date: 2018/10/15 18:08
 */
func (c *Client) SendToAddressEntire(address string, amount float64, comment string, commentTo string,
	subtractfeefromamount bool, replaceable bool, confTarget int32, estimateMode string) (*string, error) {
	return c.SendToAddressAsync(address, amount, comment, commentTo, subtractfeefromamount, replaceable,
		confTarget, estimateMode).Receive()
}

/*
Description:
SetHDSeedAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SetHDSeed for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) SetHDSeedAsync(newkeypool bool, seed string) futures.FutureResult {
	cmd := NewCommand("sethdseed", newkeypool, seed)
	return c.sendCmd(cmd)
}

/*
Description: 
Set or generate a new HD wallet seed. Non-HD wallets will not be upgraded to being a HD wallet. Wallets that are already
HD will have a new HD seed set so that new keys added to the keypool will be derived from this new seed.

Note that you will need to MAKE A NEW BACKUP of your wallet after setting the HD wallet seed.
 * Author: architect.bian
 * Date: 2018/10/15 18:24
 */
func (c *Client) SetHDSeed() (error) {
	cmd := NewCommand("sethdseed")
	_, err := futures.FutureResult(c.sendCmd(cmd)).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
Set or generate a new HD wallet seed. Non-HD wallets will not be upgraded to being a HD wallet. Wallets that are already
HD will have a new HD seed set so that new keys added to the keypool will be derived from this new seed.
 * Author: architect.bian
 * Date: 2018/10/15 18:28
 */
func (c *Client) SetHDSeedEntire(newkeypool bool, seed string) (error) {
	_, err := c.SetHDSeedAsync(newkeypool, seed).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
SetTXFeeAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SetTXFee for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) SetTXFeeAsync(amount float64) futures.FutureBool {
	cmd := NewCommand("settxfee", amount)
	return c.sendCmd(cmd)
}

/*
Description: 
Set the transaction fee per kB for this wallet. Overrides the global -paytxfee command line parameter.
 * Author: architect.bian
 * Date: 2018/10/15 18:30
 */
func (c *Client) SetTXFee(amount float64) (bool, error) {
	return c.SetTXFeeAsync(amount).Receive()
}

/*
Description:
SignMessageAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SignMessage for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) SignMessageAsync(address string, message string) futures.FutureString {
	cmd := NewCommand("signmessage", address, message)
	return c.sendCmd(cmd)
}

/*
Description: 
Sign a message with the private key of an address
 * Author: architect.bian
 * Date: 2018/10/15 18:31
 */
func (c *Client) SignMessage(address string, message string) (*string, error) {
	return c.SignMessageAsync(address, message).Receive()
}

/*
Description:
SignRawtransactionWithWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SignRawtransactionWithWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) SignRawtransactionWithWalletAsync(tx []byte, prevtxs interface{}, sighashtype string) futures.FutureResult {
	cmd := NewCommand("signrawtransactionwithwallet", tx, prevtxs, sighashtype)
	return c.sendCmd(cmd)
}

/*
Description: 
signrawtransactionwithwallet "hexstring" ( [{"txid":"id","vout":n,"scriptPubKey":"hex","redeemScript":"hex"},...] sighashtype )

Sign inputs for raw transaction (serialized, hex-encoded).
The second optional argument (may be null) is an array of previous transaction outputs that
this transaction depends on but may not yet be in the block chain.
 * Author: architect.bian
 * Date: 2018/10/15 18:31
 */
func (c *Client) SignRawtransactionWithWallet(tx []byte) (*interface{}, error) {
	cmd := NewCommand("signrawtransactionwithwallet")
	return futures.FutureResult(c.sendCmd(cmd)).Receive()
}

/*
Description: 
Sign inputs for raw transaction (serialized, hex-encoded).
The second optional argument (may be null) is an array of previous transaction outputs that
this transaction depends on but may not yet be in the block chain.
 * Author: architect.bian
 * Date: 2018/10/15 18:37
 */
func (c *Client) SignRawtransactionWithWalletEntire(tx []byte, prevtxs interface{}, sighashtype string) (*interface{}, error) {
	return c.SignRawtransactionWithWalletAsync(tx, prevtxs, sighashtype).Receive()
}

/*
Description:
UnloadWalletAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See UnloadWallet for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) UnloadWalletAsync(walletName string) futures.FutureResult {
	cmd := NewCommand("unloadwallet", walletName)
	return c.sendCmd(cmd)
}

/*
Description:
Unloads the wallet referenced by the request endpoint otherwise unloads the wallet specified in the argument.
Specifying the wallet name on a wallet endpoint is invalid.
 * Author: architect.bian
 * Date: 2018/10/15 18:42
 */
func (c *Client) UnloadWallet(walletName string) (error) {
	_, err := c.UnloadWalletAsync(walletName).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
WalletCreateFundedPSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See WalletCreateFundedPSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) WalletCreateFundedPSBTAsync() futures.FutureString {
	cmd := NewCommand("walletcreatefundedpsbt")
	return c.sendCmd(cmd)
}

func (c *Client) WalletCreateFundedPSBT() (*string, error) { //TODO unimplement
	return c.WalletCreateFundedPSBTAsync().Receive()
}

/*
Description:
WalletLockAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See WalletLock for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) WalletLockAsync() futures.FutureResult {
	cmd := NewCommand("walletlock")
	return c.sendCmd(cmd)
}

/*
Description:
Removes the wallet encryption key from memory, locking the wallet.
After calling this method, you will need to call walletpassphrase again
before being able to call any methods which require the wallet to be unlocked.
 * Author: architect.bian
 * Date: 2018/10/15 18:48
 */
func (c *Client) WalletLock() (error) {
	_, err := c.WalletLockAsync().Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
WalletPassphraseAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See WalletPassphrase for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) WalletPassphraseAsync(passphrase string, timeout int32) futures.FutureResult {
	cmd := NewCommand("walletpassphrase", passphrase, timeout)
	return c.sendCmd(cmd)
}

/*
Description: 
Stores the wallet decryption key in memory for 'timeout' seconds.
This is needed prior to performing transactions related to private keys such as sending bitcoins
 * Author: architect.bian
 * Date: 2018/10/15 18:50
 */
func (c *Client) WalletPassphrase(passphrase string, timeout int32) (error) {
	_, err := c.WalletPassphraseAsync(passphrase, timeout).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
WalletPassphraseChangeAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See WalletPassphraseChange for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) WalletPassphraseChangeAsync(oldpassphrase string, newpassphrase string) futures.FutureString {
	cmd := NewCommand("walletpassphrasechange", oldpassphrase, newpassphrase)
	return c.sendCmd(cmd)
}

/*
Description:
Changes the wallet passphrase from 'oldpassphrase' to 'newpassphrase'.
 * Author: architect.bian
 * Date: 2018/10/15 18:51
 */
func (c *Client) WalletPassphraseChange(oldpassphrase string, newpassphrase string) (error) {
	_, err := c.WalletPassphraseChangeAsync(oldpassphrase, newpassphrase).Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
WalletProcessPSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See WalletProcessPSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 12:46
 */
func (c *Client) WalletProcessPSBTAsync(psbt string, sign bool, sighashtype string, bip32derivs bool) futures.FutureResult {
	cmd := NewCommand("walletprocesspsbt", psbt, sign, sighashtype, bip32derivs)
	return c.sendCmd(cmd)
}

/*
Description:
Update a PSBT with input information from our wallet and then sign inputs
that we can sign for.
 * Author: architect.bian
 * Date: 2018/10/15 18:53
 */
func (c *Client) WalletProcessPSBT(psbt string) (*interface{}, error) {
	return c.WalletProcessPSBTEntire(psbt, true, "ALL", false)
}

/*
Description:
Update a PSBT with input information from our wallet and then sign inputs
that we can sign for.
 * Author: architect.bian
 * Date: 2018/10/15 18:53
 */
func (c *Client) WalletProcessPSBTEntire(psbt string, sign bool, sighashtype string, bip32derivs bool) (*interface{}, error) {
	return c.WalletProcessPSBTAsync(psbt, sign, sighashtype, bip32derivs).Receive()
}