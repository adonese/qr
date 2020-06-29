package main

import (
	"testing"
)

func Test_computeCRC(t *testing.T) {
	type args struct {
		data []byte
		poly uint32
	}
	v := "FFFF0002010102115138000310901020002090000000650408f8bfe0f85204000053039385802SD5916Merchant Name6008City"
	data := args{data: []byte(v), poly: 1021}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"crc32 for data", data, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeCRC(tt.args.data, tt.args.poly); got != tt.want {
				t.Errorf("computeCRC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeCRC16(t *testing.T) {
	type args struct {
		data []byte
	}
	// v := "0002010102115138000310901020002090000000650408f8bfe0f85204000053039385802SD5916Merchant Name6008City"
	// v2 := "00020101021226410014A000000615000101065016640209123456789520499995303458540510.005802MY5909QRCSDNBHD6005BANGI610543650"
	v3 := "00020101021226410014A000000615000101065016640209123456789520499995303458540510.005802MY5909QRCSDNBHD6005BANGI6105436506304"
	// data := args{data: []byte(v)}
	// data2 := args{data: []byte(v2)}
	data3 := args{data: []byte(v3)}

	tests := []struct {
		name string
		args args
		want uint16
	}{
		// {"compute crc", data, 0x5DD9},
		// {"new case", data2, 0xBFCA}, //0x6304BFCA
		// {"from code", args{data: []byte("123456789")}, 0x906e},
		{"new data", data3, 0xBFCA}, //0x0F96
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeCRC16(tt.args.data); got != tt.want {
				t.Errorf("ComputeCRC16(): 0x%04x want 0x%04x", got, tt.want)
			}
		})
	}
}
