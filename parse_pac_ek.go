package main

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/google/go-attestation/attributecert"
)

func main() {
	// Load PAC file
	data, err := os.ReadFile("attributecert/testdata/TCG_acert.der")
	if err != nil {
		panic(err)
	}

	// Parse the attribute certificate
	ac, err := attributecert.ParseAttributeCertificate(data)
	if err != nil {
		panic(err)
	}

	// Print certificate metadata
	fmt.Println("============= Print in Main ===================")
	//fmt.Printf("Result:\n +v\n", ac)

	b, _ := json.MarshalIndent(ac, "", "  ")
	fmt.Println("----- Attribute Certificate (JSON) -----")
	fmt.Println(string(b))
}
