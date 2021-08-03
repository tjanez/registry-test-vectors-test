package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	memorySigner "github.com/oasisprotocol/oasis-core/go/common/crypto/signature/signers/memory"
	"github.com/oasisprotocol/oasis-core/go/common/entity"
)

const ENTITY_DESCRIPTOR_SIGNATURE_CONTEXT = "oasis-core/registry: register entity"

func main() {
	// Read JSON as generic data.
	data, _ := ioutil.ReadFile("register-entity.json")
	var vector map[string]interface{}
	err := json.Unmarshal(data, &vector)
	if err != nil {
		fmt.Printf("Error decoding test vector: %s", err)
		os.Exit(1)
	}

	// Parse some JSON fields.
	tx := vector["tx"].(map[string]interface{})
	body := tx["body"].(map[string]interface{})
	signtr := body["signature"].(map[string]interface{})

	// Get Entity.
	entData, err := json.Marshal(body["untrusted_raw_value"])
	if err != nil {
		fmt.Printf("Failed to marshal untrusted_raw_value: %s", err)
		os.Exit(1)
	}
	var ent entity.Entity
	err = json.Unmarshal(entData, &ent)
	if err != nil {
		fmt.Printf("Failed to parse entity: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Entity: %+v\n", ent)

	// Initialize the same signer as used in test vectors.
	// NOTE for Non-Go implementations: Initialize a new ed25519 signer from the private key
	// in test vector's "signer_private_key" field.
	testSigner := memorySigner.NewTestSigner("oasis-core registry test vectors: RegisterEntity signer")
	// Check if signer matches.
	if testSigner.Public().String() != vector["signer_public_key"].(string) {
		fmt.Printf("Test signer doesn't match test vector's signer")
		os.Exit(1)
	}

	// Sign entity descriptor.
	sigEnt, err := entity.SignEntity(testSigner, signature.NewContext(ENTITY_DESCRIPTOR_SIGNATURE_CONTEXT), &ent)
	if err != nil {
		fmt.Printf("Error signing entity: %s", err)
		os.Exit(1)
	}
	// Check if signature matches.
	if sigEnt.Signature.Signature.String() != signtr["signature"] {
		fmt.Printf("Signature doesn't match test vector's signature")
		os.Exit(1)
	}
	fmt.Println("SUCCESS: Entity descriptor signature matches!")
}
