package internal

import "runtime"

func LineFeed() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
