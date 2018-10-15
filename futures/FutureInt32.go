package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureInt32 is a future promise to deliver the result of int32 type.
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
type FutureInt32 chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
result of int32 type.
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
func (f FutureInt32) Receive() (int32, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as a int32.
	var result int32
	err = json.Unmarshal(res, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
