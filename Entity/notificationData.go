package entity

import "time"

//Notification struct
type Notification struct {
	NotfDescription string
	NotfTitle       string
	NotfLevel       string
	NotfDate        time.Time
}
