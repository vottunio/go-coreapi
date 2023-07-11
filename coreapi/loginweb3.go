package coreapi

import (
	"fmt"
	"log"
	"net/http"
)

// Prepares a web3 message for sign-in with ethereum meeting "ERC-4361: Sign-In with Ethereum"
func (c *CoreApi) Web3PrepareMessage(requestDto *MessageWeb3DTO) (*MessageWeb3ResponseDTO, error) {

	var responseDto *MessageWeb3ResponseDTO

	err := c.sendCoreTransaction(PrepareWeb3MessageUrl, http.MethodPost, &requestDto, &responseDto)

	if err != nil {
		log.Printf("An error has raised calling core api to prepare web3 message. %+v", err)
		return nil, err
	}

	return responseDto, nil
}

// Validates a web3 sign-in signature meeting "ERC-4361: Sign-In with Ethereum"
func (c *CoreApi) Web3ValidateSignature(nonce, signature string) error {

	url := fmt.Sprintf(ValidateWeb3SignatureUrl, nonce, signature)

	err := c.sendCoreTransaction(url, http.MethodPost, nil, nil)

	if err != nil {
		log.Printf("An error has raised calling core api to validate web3 message signature. %+v", err)
		return err
	}

	return nil
}
