package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureResult is a future promise to deliver the result of a RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/08/26 18:56
 */
type FutureResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
data structure requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/08/26 18:56
 */
func (r FutureResult) Receive(result interface{}) (*interface{}, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

/*
Description:
receiveFuture receives from the passed futureResult channel to extract a
reply or any errors.  The examined errors include an error in the
futureResult and the error in the reply from the server.  This will block
until the result is available on the passed channel.
 * Author: architect.bian
 * Date: 2018/09/14 12:56
 */
func ReceiveFuture(f chan *base.Response) ([]byte, error) { //TODO futerReceive?!
	// Wait for a response on the returned channel.
	r := <-f
	return r.Result, r.Err
}

/*
Description:
newFutureError returns a new future result channel that already has the
passed error waitin on the channel with the reply set to nil.  This is useful
to easily return errors from the various Async functions.
 * Author: architect.bian
 * Date: 2018/09/14 12:57
 */
func NewFutureError(err error) chan *base.Response {
	responseChan := make(chan *base.Response, 1)
	responseChan <- &base.Response{Err: err}
	return responseChan
}