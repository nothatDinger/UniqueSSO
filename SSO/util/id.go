package util

import "github.com/labstack/gommon/random"

func NewSMSCode() string {
	return random.String(6, random.Numeric)
}
