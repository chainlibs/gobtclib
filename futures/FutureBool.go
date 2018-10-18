package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureBool is a future promise to deliver the result of bool type.
 * Author: architect.bian
 * Date: 2018/10/15 10:57
 */
type FutureBool chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
result of float64 type.
 * Author: architect.bian
 * Date: 2018/10/15 10:57
 */
func (f FutureBool) Receive() (bool, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return false, err
	}

	// Unmarshal the result as a float64.
	var result bool
	err = json.Unmarshal(res, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}
