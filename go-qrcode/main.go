package main

import "fmt"
import qrcode "github.com/skip2/go-qrcode"
import "flag"

func main() {

	//var png []byte
	//png, err := qrcode.Encode("https://github.com/jeffotoni", qrcode.Medium, 256)
	//err := qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.Black, color.White, "qr.png")

	qrCode := flag.String("code", "https://github.com/jeffotoni", "a string")
	flag.Parse()

	if len(*qrCode) <= 0 {
		flag.PrintDefaults()
		return
	}

	err := qrcode.WriteFile(*qrCode, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Gerado imagem qr.png com sucesso")
}
