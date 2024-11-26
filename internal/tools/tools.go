package tools

import "strconv"

func GetPort(port int) string {
	return ":" + strconv.Itoa(port)
}
