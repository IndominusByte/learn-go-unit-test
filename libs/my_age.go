package libs

import "strconv"

func MyAge(age int64) string {
	return "my age is " + strconv.FormatInt(age, 10)
}
