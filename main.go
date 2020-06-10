package main

import (
	"log"
	"strconv"
)

func main() {

}

//QR object
type QR struct {
	ID    string
	Value string
}

type Merchant struct {
	ID                     string
	QRType                 string
	IsMerchant             bool
	IsDynamic              bool
	MerchantInfo           string // length 99
	MerchantCode           string //0000 default, 9999 for governmental
	TransactionCode        int    // 938 for SDG
	Amount                 float32
	TipIndicator           bool // O
	FixedTipIndicator      bool
	PercentageTipIndicator bool
	FixedTipVal            float32
	PerentageTip           float32
	CountryCode            string
	Name                   string
	City                   string // max 15
	PostalCode             string //OOO 11111
	AdditionalData         string //O ~99 check table
	CRC                    int    // checksum for qr data
	I18nMerchantInfo       string
}

//Encode string to Merchant object
func (m *Merchant) Encode(s string) {
	if len(s) < 5 {
		return
	}

	log.Printf(s[:2])
	switch s[:2] {
	case "00":
		length := toInt(s[2:4])
		if s[4:4+length] == "01" {
			m.IsMerchant = true
		}
		s = s[4+length:] // VERY important

	case "01":
		length := toInt(s[2:4])
		if s[4:4+length] == "12" {
			m.IsDynamic = true
		}
		s = s[4+length:] // VERY important

	case "51":
		length := toInt(s[2:4])
		m.MerchantInfo = s[4:length]
		s = s[length+4:] // FIXME check this one

	case "52":
		m.MerchantCode = s[2:6] // 52 04 0000
		s = s[8:]

	case "53":
		m.TransactionCode = toInt(s[4 : 4+3]) // i was not sure 2 + 3 equals what
		// 5303938
		s = s[7:]

	case "54": // Transaction amount
		length := toInt(s[2:4])
		m.Amount = toFloat(s[4:length])
		s = s[2+2+length:]

	case "55": // there is a bug or inconsitency here
		if s[2:4] == "01" {
			m.TipIndicator = true
		} else if s[2:4] == "02" {
			m.FixedTipIndicator = true
		} else if s[2:4] == "03" {
			m.PercentageTipIndicator = true
		}
		s = s[4:]
	case "56":
		length := toInt(s[2:4])
		m.FixedTipVal = toFloat(s[4 : 4+length])
		s = s[2+2+length:]

	case "57":
		length := toInt(s[2:4])
		m.PerentageTip = toFloat(s[4 : 4+length])
		s = s[2+2+length:]

	case "58":
		m.CountryCode = s[4 : 4+2]
		s = s[6:]

	case "59": // merchant name
		length := toInt(s[2:4]) //59 16 MerchantName
		m.Name = s[4 : 4+length]
		s = s[4+length:]

	case "60": // city code
		length := toInt(s[2:4])

		m.City = s[4 : 4+length]
		s = s[2+2+length:]
	case "61":
		m.PostalCode = s[4 : 4+5] //FIXME
		s = s[9:]
	case "62":
		length := toInt(s[2:4])
		m.AdditionalData = s[4 : 4+length]
		s = s[4+length:]
	case "63":
		m.CRC = toInt(s[2 : 2+4])
		s = s[2+4:]
	}
	m.Encode(s)

}

func toInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func toFloat(s string) float32 {
	i, _ := strconv.ParseFloat(s, 10)
	return float32(i)
}
