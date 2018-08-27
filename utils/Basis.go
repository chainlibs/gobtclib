package utils

/*
Description:
basisType is the util object with basis method, you can invoke method like utils.Basis.xxx()
 * Author: architect.bian
 * Date: 2018/08/27 14:04
 */
type basisType struct {}

var Basis basisType
/*
Description:
Bool is a helper routine that allocates a new bool value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:28
 */
func (_ basisType) Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

/*
Description:
Int is a helper routine that allocates a new int value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:28
 */
func (_ basisType) Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

/*
Description:
Uint is a helper routine that allocates a new uint value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:28
 */
func (_ basisType) Uint(v uint) *uint {
	p := new(uint)
	*p = v
	return p
}

/*
Description:
Int32 is a helper routine that allocates a new int32 value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:28
 */
func (_ basisType) Int32(v int32) *int32 {
	p := new(int32)
	*p = v
	return p
}

/*
Description:
Uint32 is a helper routine that allocates a new uint32 value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:28
 */
func Uint32(v uint32) *uint32 {
	p := new(uint32)
	*p = v
	return p
}

/*
Description:
Int64 is a helper routine that allocates a new int64 value to store v and
returns a pointer to it.  This is useful when assigning optional parameters
 * Author: architect.bian
 * Date: 2018/08/27 13:28
 */
func (_ basisType) Int64(v int64) *int64 {
	p := new(int64)
	*p = v
	return p
}

/*
Description:
Uint64 is a helper routine that allocates a new uint64 value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:29
 */
func (_ basisType) Uint64(v uint64) *uint64 {
	p := new(uint64)
	*p = v
	return p
}

/*
Description:
Float64 is a helper routine that allocates a new float64 value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:29
 */
func (_ basisType) Float64(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

/*
Description:
String is a helper routine that allocates a new string value to store v and
returns a pointer to it.  This is useful when assigning optional parameters.
 * Author: architect.bian
 * Date: 2018/08/27 13:29
 */
func (_ basisType) String(v string) *string {
	p := new(string)
	*p = v
	return p
}