package pkg

import "strconv"

type ID = int64

func ParseIdToInt64(val int) ID {
	return int64(val)
}

func ParseIdToString(val int64) string {
	return strconv.Itoa(int(val))
}
