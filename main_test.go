package main

import (
	"testing"
)

func TestMerchant_Decode(t *testing.T) {
	type fields struct {
		ID                     string
		QRType                 string
		IsMerchant             bool
		IsDynamic              bool
		MerchantInfo           string
		MerchantCode           string
		TransactionCode        int
		Amount                 float32
		TipIndicator           bool
		FixedTipIndicator      bool
		PercentageTipIndicator bool
		FixedTipVal            float32
		PerentageTip           float32
		CountryCode            string
		Name                   string
		City                   string
		PostalCode             string
		AdditionalData         string
		CRC                    int
		I18nMerchantInfo       string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
	}{
		{"smaller test", "000201"},
		// {"smaller test", "00020112"},
		{"EBS sample", "0002010102115138000310901020002090000000650408f8bfe0f85204000053039385802SD5912MerchantName6004City63045DD9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Merchant{}
			m.Decode(tt.args)
			t.Logf("The value of merchant is: %#v", m)
		})
	}
}

func Test_toInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"to int test", args{"8"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt(tt.args.s); got != tt.want {
				t.Errorf("toInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchant_Encode(t *testing.T) {
	type fields struct {
		ID                     string
		QRType                 string
		IsMerchant             bool
		IsDynamic              bool
		MerchantInfo           string
		MerchantCode           string
		TransactionCode        int
		Amount                 float32
		TipIndicator           bool
		FixedTipIndicator      bool
		PercentageTipIndicator bool
		FixedTipVal            float32
		PerentageTip           float32
		CountryCode            string
		Name                   string
		City                   string
		PostalCode             string
		AdditionalData         string
		CRC                    int
		I18nMerchantInfo       string
	}
	tests := []struct {
		name string
		want string
	}{
		{"testing encoding", "something"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Merchant{}
			if got := m.Encode(); got != tt.want {
				t.Errorf("Merchant.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getValue(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"getting formatted length", args{"ahmed"}, "05"},
		{"getting formatted length", args{""}, "00"},
		{"getting formatted length", args{10000001000000}, "14"},
		{"getting formatted length", args{1000.12}, "07"},
		{"getting formatted length", args{101200.12}, "09"},
		{"getting formatted length", args{101200.1200}, "11"}, // notice the truncation after 02f
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValue(tt.args.i); got != tt.want {
				t.Errorf("getValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
