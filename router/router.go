package router

import (
	"fmt"
	"github.com/siddontang/mixer/sqltypes"
	"hash/crc32"
)

const (
	NumberKey int = 1 << iota
	StringKey
)

const (
	MinKey = ""
	MaxKey = ""
)

// [start, end)
type KeyRange struct {
	Start string
	End   string
}

func (kr KeyRange) Contains(i string) bool {
	return kr.Start <= i && (kr.End == MaxKey || i < kr.End)
}

type Rule interface {
	FindNodes(keys ...sqltypes.Value) (map[int][]sqltypes.Value, error)
}

type HashRule struct {
	Rule
	HashMod int
}

func (r *HashRule) FindNodes(keys ...sqltypes.Value) (map[int][]sqltypes.Value, error) {
	m := make(map[int][]sqltypes.Value, len(keys))
	var n int = 0
	for _, key := range keys {
		if key.IsNumeric() {
			i, err := key.ParseUint64()
			if err != nil {
				return nil, err
			}

			n = int(i) % r.HashMod
		} else if key.IsString() {
			i := crc32(key.Raw())

			n = int(i) % r.HashMod
		} else {
			return nil, fmt.Errorf("invalid value type, must number and string")
		}

		m[n] = append(m[n], key)
	}

	return m, nil
}

type RangeRule struct {
	Rule
	Ranges []KeyRange
}

func (r *RangeRule) FindNodes(keys ...sqltypes.Value) (map[int][]sqltypes.Value, error) {
	m := make(map[int][]sqltypes.Value, len(keys))

	return m, nil
}
