package strings

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmptyStr = errors.New(`strings pks: string is empty`)
)

func SplitStrToInt32s(str string) (is []int32, err error) {
	if len(str) < 1 {
		return nil, ErrEmptyStr
	}
	s := strings.Split(str, `,`)
	is = make([]int32, 0, len(s))
	var i int64
	for _, s2 := range s {
		if i, err = strconv.ParseInt(s2, 10, 32); err != nil {
			return nil, err
		}
		is = append(is, int32(i))
	}
	return
}
