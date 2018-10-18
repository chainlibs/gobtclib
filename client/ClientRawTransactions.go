package client

import (
	"github.com/chainlibs/gobtclib/futures"
)

/*
Description:
CombinePSBT returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CombinePSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CombinePSBTAsync(txs []string) futures.FutureResult {
	cmd := NewCommand("combinepsbt", txs)
	return c.sendCmd(cmd)
}

/*
Description: 
Combine multiple partially signed Bitcoin transactions into one transaction.
Implements the Combiner role.
 * Author: architect.bian
 * Date: 2018/10/15 20:18
 */
func (c *Client) CombinePSBT(txs []string) (*interface{}, error) {
	return c.CombinePSBTAsync(txs).Receive()
}

/*
Description:
CombineRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CombineRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CombineRawTransactionAsync(txs []string) futures.FutureResult {
	cmd := NewCommand("combinerawtransaction", txs)
	return c.sendCmd(cmd)
}

/*
Description:
Combine multiple partially signed transactions into one transaction.
The combined transaction may be another partially signed transaction or a
fully signed transaction.
 * Author: architect.bian
 * Date: 2018/10/15 20:18
 */
func (c *Client) CombineRawTransaction(txs []string) (*interface{}, error) {
	return c.CombineRawTransactionAsync(txs).Receive()
}

/*
Description:
ConvertTopSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See ConvertTopSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) ConvertTopSBTAsync(tx []byte, permitsigdata bool, iswitness bool) futures.FutureString {
	cmd := NewCommand("converttopsbt", tx, permitsigdata, iswitness)
	return c.sendCmd(cmd)
}

/*
Description:
Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction
createpsbt and walletcreatefundedpsbt should be used for new applications.
 * Author: architect.bian
 * Date: 2018/10/15 20:22
 */
func (c *Client) ConvertTopSBT(tx []byte) (*string, error) {
	cmd := NewCommand("converttopsbt")
	return futures.FutureString(c.sendCmd(cmd)).Receive()
}

/*
Description: 
Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction
createpsbt and walletcreatefundedpsbt should be used for new applications.
 * Author: architect.bian
 * Date: 2018/10/15 20:23
 */
func (c *Client) ConvertTopSBTEntire(tx []byte, permitsigdata bool, iswitness bool) (*string, error) {
	return c.ConvertTopSBTAsync(tx, permitsigdata, iswitness).Receive()
}

/*
Description:
CreatepSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CreatepSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CreatepSBTAsync(inputs interface{}, outputs interface{}, locktime int32, replaceable bool) futures.FutureString {
	cmd := NewCommand("createpsbt", inputs, outputs, locktime, replaceable)
	return c.sendCmd(cmd)
}

/*
Description:
Creates a transaction in the Partially Signed Transaction format.
Implements the Creator role.
 * Author: architect.bian
 * Date: 2018/10/15 20:27
 */
func (c *Client) CreatepSBT(inputs interface{}, outputs interface{}, locktime int32, replaceable bool) (*string, error) {
	return c.CreatepSBTAsync(inputs, outputs, locktime, replaceable).Receive()
}

/*
Description:
CreateRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See CreateRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) CreateRawTransactionAsync(inputs interface{}, outputs interface{}, locktime int32, replaceable bool) futures.FutureString {
	cmd := NewCommand("createrawtransaction", inputs, outputs, locktime, replaceable)
	return c.sendCmd(cmd)
}

/*
Description:
Create a transaction spending the given inputs and creating new outputs.
Outputs can be addresses or data.
Returns hex-encoded raw transaction.
Note that the transaction's inputs are not signed, and
it is not stored in the wallet or transmitted to the network.
 * Author: architect.bian
 * Date: 2018/10/15 20:31
 */
func (c *Client) CreateRawTransaction(inputs interface{}, outputs interface{}, locktime int32, replaceable bool) (*string, error) {
	return c.CreateRawTransactionAsync(inputs, outputs, locktime, replaceable).Receive()
}

/*
Description:
DecodepSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DecodepSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) DecodepSBTAsync(psbt string) futures.FutureResult {
	cmd := NewCommand("decodepsbt", psbt)
	return c.sendCmd(cmd)
}

/*
Description: 
Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.
 * Author: architect.bian
 * Date: 2018/10/15 20:32
 */
func (c *Client) DecodepSBT(psbt string) (*interface{}, error) {
	return c.DecodepSBTAsync(psbt).Receive()
}

/*
Description:
DecodeRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DecodeRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) DecodeRawTransactionAsync(tx []byte, isWitness bool) futures.FutureResult {
	cmd := NewCommand("decoderawtransaction", string(tx), isWitness)
	return c.sendCmd(cmd)
}

/*
Description: 
Return a JSON object representing the serialized, hex-encoded transaction.
 * Author: architect.bian
 * Date: 2018/10/15 20:34
 */
func (c *Client) DecodeRawTransaction(tx []byte, isWitness bool) (*interface{}, error) {
	return c.DecodeRawTransactionAsync(tx, isWitness).Receive()
}

/*
Description:
DecodeScriptAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See DecodeScript for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) DecodeScriptAsync(scriptHex string) futures.FutureResult {
	cmd := NewCommand("decodescript", scriptHex)
	return c.sendCmd(cmd)
}

/*
Description: 
Decode a hex-encoded script.
 * Author: architect.bian
 * Date: 2018/10/15 20:36
 */
func (c *Client) DecodeScript(scriptHex string) (*interface{}, error) {
	return c.DecodeScriptAsync(scriptHex).Receive()
}

/*
Description:
FinalizepSBTAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See FinalizepSBT for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) FinalizepSBTAsync(psbt string, extract bool) futures.FutureResult {
	cmd := NewCommand("finalizepsbt", psbt, extract)
	return c.sendCmd(cmd)
}

/*
Description:
Finalize the inputs of a PSBT. If the transaction is fully signed, it will produce a
network serialized transaction which can be broadcast with sendrawtransaction. Otherwise a PSBT will be
created which has the final_scriptSig and final_scriptWitness fields filled for inputs that are complete.
Implements the Finalizer and Extractor roles.
 * Author: architect.bian
 * Date: 2018/10/15 20:38
 */
func (c *Client) FinalizepSBT(psbt string, extract bool) (*interface{}, error) {
	return c.FinalizepSBTAsync(psbt, extract).Receive()
}

/*
Description:
FundRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See FundRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) FundRawTransactionAsync(tx []byte, options map[string]interface{}, iswitness bool) futures.FutureResult {
	cmd := NewCommand("fundrawtransaction", tx, options, iswitness)
	return c.sendCmd(cmd)
}

/*
Description:
Add inputs to a transaction until it has enough in value to meet its out value.
This will not modify existing inputs, and will add at most one change output to the outputs.
No existing outputs will be modified unless "subtractFeeFromOutputs" is specified.
Note that inputs which were signed may need to be resigned after completion since in/outputs have been added.
The inputs added will not be signed, use signrawtransaction for that.
Note that all existing inputs must have their previous output transaction be in the wallet.
Note that all inputs selected must be of standard form and P2SH scripts must be
in the wallet using importaddress or addmultisigaddress (to calculate fees).
You can see whether this is the case by checking the "solvable" field in the listunspent output.
Only pay-to-pubkey, multisig, and P2SH versions thereof are currently supported for watch-only
 * Author: architect.bian
 * Date: 2018/10/15 20:36
 */
func (c *Client) FundRawTransaction(tx []byte, options map[string]interface{}, iswitness bool) (*interface{}, error) {
	return c.FundRawTransactionAsync(tx, options, iswitness).Receive()
}

/*
Description:
GetRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) GetRawTransactionAsync(txid string, verbose bool, blockhash string) futures.FutureResult {
	cmd := NewCommand("getrawtransaction", txid, verbose, blockhash)
	return c.sendCmd(cmd)
}

/*
Description:
NOTE: By default this function only works for mempool transactions. If the -txindex option is
enabled, it also works for blockchain transactions. If the block which contains the transaction
is known, its hash can be provided even for nodes without -txindex. Note that if a blockhash is
provided, only that block will be searched and if the transaction is in the mempool or other
blocks, or if this node does not have the given block available, the transaction will not be found.
DEPRECATED: for now, it also works for transactions with unspent outputs.

Return the raw transaction data.
 * Author: architect.bian
 * Date: 2018/10/15 20:41
 */
func (c *Client) GetRawTransaction(txid string, blockhash string) ([]byte, error) {
	cmd := NewCommand("getrawtransaction", txid, false, blockhash)
	return futures.FutureByteArray(c.sendCmd(cmd)).Receive()
}

/*
Description:
Return the verbose raw transaction data.
 * Author: architect.bian
 * Date: 2018/10/15 20:42
 */
func (c *Client) GetRawTransactionVerbose(txid string, blockhash string) (*interface{}, error) {
	return c.GetRawTransactionAsync(txid, true, blockhash).Receive()
}

/*
Description:
SendRawTransactionAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SendRawTransaction for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) SendRawTransactionAsync(tx []byte, allowhighfees bool) futures.FutureResult {
	cmd := NewCommand("sendrawtransaction", tx, allowhighfees)
	return c.sendCmd(cmd)
}

/*
Description: 
Submits raw transaction (serialized, hex-encoded) to local node and network.

Also see createrawtransaction and signrawtransaction calls.
 * Author: architect.bian
 * Date: 2018/10/15 20:54
 */
func (c *Client) SendRawTransaction(tx []byte, allowhighfees bool) (*interface{}, error) {
	return c.SendRawTransactionAsync(tx, allowhighfees).Receive()
}

/*
Description:
SignRawTransactionWithKeyAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See SignRawTransactionWithKey for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) SignRawTransactionWithKeyAsync(tx []byte, privkeys []string, prevtxs interface{}, sighashtype string) futures.FutureResult {
	cmd := NewCommand("signrawtransactionwithkey", tx, privkeys, prevtxs, sighashtype)
	return c.sendCmd(cmd)
}

/*
Description:
Sign inputs for raw transaction (serialized, hex-encoded).
The second argument is an array of base58-encoded private
keys that will be the only keys used to sign the transaction.
The third optional argument (may be null) is an array of previous transaction outputs that
this transaction depends on but may not yet be in the block chain.
 * Author: architect.bian
 * Date: 2018/10/15 21:11
 */
func (c *Client) SignRawTransactionWithKey(tx []byte, privkeys []string, prevtxs interface{}, sighashtype string) (*interface{}, error) {
	return c.SignRawTransactionWithKeyAsync(tx, privkeys, prevtxs, sighashtype).Receive()
}

/*
Description:
TestMempoolAcceptAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See TestMempoolAccept for more details.
 * Author: architect.bian
 * Date: 2018/10/15 20:09
 */
func (c *Client) TestMempoolAcceptAsync(rawTXs []string, allowHighFees bool) futures.FutureResult {
	cmd := NewCommand("testmempoolaccept", rawTXs, allowHighFees)
	return c.sendCmd(cmd)
}

/*
Description:
Returns if raw transaction (serialized, hex-encoded) would be accepted by mempool.
 * Author: architect.bian
 * Date: 2018/10/15 23:13
 */
func (c *Client) TestMempoolAccept(rawTXs []string, allowHighFees bool) (*interface{}, error) {
	return c.TestMempoolAcceptAsync(rawTXs, allowHighFees).Receive()
}
