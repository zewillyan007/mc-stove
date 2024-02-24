package util

import (
	qrcode "github.com/skip2/go-qrcode"
)

type QRCode struct{}

func NewQRCode() *QRCode {
	return &QRCode{}
}

func (*QRCode) GenerateQRcode(input string) []byte {
	var png []byte
	png, _ = qrcode.Encode(input, qrcode.Medium, 256)
	return png
}
