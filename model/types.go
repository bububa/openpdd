package model

import "strconv"

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = Uint64(i)
	return
}

func (u64 Uint64) Value() uint64 {
	return uint64(u64)
}

// Int support string quoted number in json
type Int int

// UnmarshalJSON implement json Unmarshal interface
func (i32 *Int) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.Atoi(string(b))
	*i32 = Int(i)
	return
}

func (i32 Int) Value() int {
	return int(i32)
}
