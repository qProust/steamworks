// Code generated by "stringer -type ERemoteStoragePlatform -trimprefix ERemoteStoragePlatform_"; DO NOT EDIT.

package internal

import "strconv"

const (
	_ERemoteStoragePlatform_name_0 = "AllNoneWindowsOSX"
	_ERemoteStoragePlatform_name_1 = "PS3"
	_ERemoteStoragePlatform_name_2 = "Linux"
	_ERemoteStoragePlatform_name_3 = "Reserved2"
)

var (
	_ERemoteStoragePlatform_index_0 = [...]uint8{0, 3, 7, 14, 17}
)

func (i ERemoteStoragePlatform) String() string {
	switch {
	case -1 <= i && i <= 2:
		i -= -1
		return _ERemoteStoragePlatform_name_0[_ERemoteStoragePlatform_index_0[i]:_ERemoteStoragePlatform_index_0[i+1]]
	case i == 4:
		return _ERemoteStoragePlatform_name_1
	case i == 8:
		return _ERemoteStoragePlatform_name_2
	case i == 16:
		return _ERemoteStoragePlatform_name_3
	default:
		return "ERemoteStoragePlatform(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}