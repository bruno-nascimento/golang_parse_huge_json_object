package models

import (
	"gorm.io/gorm"
)

const (
	KEY         = "key"
	NAME        = "name"
	CITY        = "city"
	COUNTRY     = "country"
	ALIAS       = "alias"
	REGIONS     = "regions"
	COORDINATES = "coordinates"
	PROVINCE    = "province"
	TIMEZONE    = "timezone"
	CODE        = "code"
)

var AllModels = []interface{}{&Port{}, &Alias{}, &Regions{}, &Unlocs{}}

type Port struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	Name      *string   `json:"name"`
	City      *string   `json:"city"`
	Country   *string   `json:"country"`
	Latitude  *float64  `json:"latitude"`
	Longitude *float64  `json:"longitude"`
	Province  *string   `json:"province"`
	Timezone  *string   `json:"timezone"`
	Code      *string   `json:"code"`
	Alias     []Alias   `json:"alias"`
	Regions   []Regions `json:"regions"`
	Unlocs    []Unlocs  `json:"unlocs"`
}

func NewPortFromMap(m map[string]interface{}) *Port {
	p := &Port{
		Name:    ptrStr(m[NAME]),
		City:    ptrStr(m[CITY]),
		Country: ptrStr(m[COUNTRY]),
		Latitude: func() *float64 {
			if c, ok := m[COORDINATES].([]interface{}); ok {
				if len(c) == 2 {
					l := c[0].(float64)
					return &l
				}
			}
			return nil
		}(),
		Longitude: func() *float64 {
			if c, ok := m[COORDINATES].([]interface{}); ok {
				if len(c) == 2 {
					l := c[1].(float64)
					return &l
				}
			}
			return nil
		}(),
		Province: ptrStr(m[PROVINCE]),
		Timezone: ptrStr(m[TIMEZONE]),
		Code:     ptrStr(m[CODE]),
	}
	return p
}

type Alias struct {
	Alias  string
	PortID string
	gorm.Model
}

type Regions struct {
	Region string
	PortID string
	gorm.Model
}

type Unlocs struct {
	PortID string
	Unloc  string
	gorm.Model
}

func ptrStr(i interface{}) *string {
	if str, ok := i.(string); ok && str != "" {
		return &str
	}
	return nil
}
