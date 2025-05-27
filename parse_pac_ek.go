package main

import (
	"fmt"
	"os"
	"github.com/google/go-attestation/attributecert"
)

func main() {
	// Load PAC file
	data, err := os.ReadFile("attributecert/testdata/PlatformCert_ekhash.der")
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

	fmt.Println("============= Parsed Attribute Certificate ===================")
	fmt.Println("Issuer:", ac.Issuer.String())
/*
	fmt.Println("Subject:", ac.Subject.String())
	fmt.Println("Serial Number:", ac.SerialNumber)
	fmt.Println("Not Before:", ac.NotBefore)
	fmt.Println("Not After:", ac.NotAfter)
	fmt.Println("Signature Algorithm:", ac.SignatureAlgorithm)

	// Print platform-specific attributes
	fmt.Println("\n==== Platform Attributes ====")
	fmt.Println("Platform Manufacturer:", ac.PlatformManufacturer)
	fmt.Println("Platform Model:", ac.PlatformModel)
	fmt.Println("Platform Version:", ac.PlatformVersion)
	fmt.Println("Platform Serial:", ac.PlatformSerial)

	// Print holder information
	fmt.Println("\n==== Holder Information ====")
	fmt.Println("Holder Issuer:", ac.Holder.Issuer.String())
	fmt.Println("Holder Serial Number:", ac.Holder.Serial)
	*/
}
