package main

import (

	"fmt"
	"os";

)

// ReadTLSClientKey reads the content of a TLS client key file
// and returns its content as a byte array.
func ReadTLSCertificate(tlsClientKey string) ([]byte, error) {
	// Check if the file exists
	if _, err := os.Stat(tlsClientKey); os.IsNotExist(err) {
		return nil, fmt.Errorf("the file %s does not exist", tlsClientKey)
	}

	// Read the file into a byte array using os.ReadFile
	keyBytes, err := os.ReadFile(tlsClientKey)
	if err != nil {
		return nil, fmt.Errorf("could not read the file: %v", err)
	}

	return keyBytes, nil
}