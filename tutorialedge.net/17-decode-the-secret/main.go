/*
👋 Welcome Gophers! In this challenge, you are tasked with decoding the secret message.

You will need to implement the DecodeSecret function which will take in a string that has been encoded using base64 encoding and decode this string.

This decoded string will be the result of a caesar cipher which has shifted all of the characters of the string up by 1 place. So you will have to ensure that when you return the result, it also decodes this cipher.
*/
package main

import (
	"encoding/base64"
	"fmt"
)

func DecodeSecret(message string) string {
	var newMessage []rune
	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		panic(err)
	}
	for i := range data {
		newMessage = append(newMessage, rune(data[i]-1))
	}

	return string(newMessage)
}

func main() {
	fmt.Println("Decode the Secret")

	message := "VEZEU0ZVVFVTSk9I"
	result := DecodeSecret(message)
	fmt.Println(result)

}
