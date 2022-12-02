package internal

import "strconv"

func GetFile(day int) string {
	day_string := strconv.Itoa(day)
	if len(day_string) < 2 {
		day_string = "0" + day_string
	}
	return "day" + day_string + ".in"
}

func GetFunc(day int) string {
	day_string := strconv.Itoa(day)
	if len(day_string) < 2 {
		day_string = "0" + day_string
	}
	return "Day" + day_string
}
