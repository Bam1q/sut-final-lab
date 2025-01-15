package entity

import (
	"gorm.io/gorm"
)

type Devices struct {
	gorm.Model
	DeviceID   string `valid:"DeviceID must start with 2 uppercase letters followed by 10 digits, matchs(^[A-Z][A-Z]\\d{10}$)"`
	IpAddress  string `valid:"ip~IpAddress must be a valid IP address"`
	DeviceName string
}
