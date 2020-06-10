package main

import (
	"testing"
)

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
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
	}{
		// {"smaller test", "000201"},
		{"EBS sample", "0002010102115138000310901020002090000000650408f8bfe0f85204000053039385802SD5912MerchantName6004City63045DD9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Merchant{}
			m.Encode(tt.args)
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
