// Code generated by "stringer -type EHTTPMethod -trimprefix EHTTPMethod_"; DO NOT EDIT.

package internal

import "strconv"

const _EHTTPMethod_name = "InvalidGETHEADPOSTPUTDELETEOPTIONSPATCH"

var _EHTTPMethod_index = [...]uint8{0, 7, 10, 14, 18, 21, 27, 34, 39}

func (i EHTTPMethod) String() string {
	if i < 0 || i >= EHTTPMethod(len(_EHTTPMethod_index)-1) {
		return "EHTTPMethod(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EHTTPMethod_name[_EHTTPMethod_index[i]:_EHTTPMethod_index[i+1]]
}