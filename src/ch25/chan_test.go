package ch25

import "time"

//csp abbr. communicating sequential processes

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}
