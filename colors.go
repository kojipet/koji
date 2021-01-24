package main

import "runtime"

var (
	nc = "\033[0m"

	brightblack   = "\033[1;30m"
	brightred     = "\033[1;31m"
	brightgreen   = "\033[1;32m"
	brightyellow  = "\033[1;33m"
	brightpurple  = "\033[1;34m"
	brightmagenta = "\033[1;35m"
	brightcyan    = "\033[1;36m"
	brightwhite   = "\033[1;37m"

	black   = "\033[0;30m"
	red     = "\033[0;31m"
	green   = "\033[0;32m"
	yellow  = "\033[0;33m"
	purple  = "\033[0;34m"
	magenta = "\033[0;35m"
	cyan    = "\033[0;36m"
	white   = "\033[0;37m"
)

func osCheck() bool {
	if runtime.GOOS == "windows" {
		Debug("OS Detected: Windows. Colors disabled.")
		nc = ""

		brightblack = ""
		brightred = ""
		brightgreen = ""
		brightyellow = ""
		brightpurple = ""
		brightmagenta = ""
		brightcyan = ""
		brightwhite = ""

		black = ""
		red = ""
		green = ""
		yellow = ""
		purple = ""
		magenta = ""
		cyan = ""
		white = ""
		return true
	}
	return false
}
