/*
2325 Decode the Message

You are given the strings key and message, which represent a cipher key and a secret message, respectively. The steps to decode message are as follows:

	Use the first appearance of all 26 lowercase English letters in key as the order of the substitution table.
	Align the substitution table with the regular English alphabet.
	Each letter in message is then substituted using the table.
	Spaces ' ' are transformed to themselves.

	For example, given key = "happy boy" (actual key would have at least one instance of each letter in the alphabet), we have the partial substitution table of ('h' -> 'a', 'a' -> 'b', 'p' -> 'c', 'y' -> 'd', 'b' -> 'e', 'o' -> 'f').

Return the decoded message.

Constraints:

	26 <= key.length <= 2000
	key consists of lowercase English letters and ' '.
	key contains every letter in the English alphabet ('a' to 'z') at least once.
	1 <= message.length <= 2000
	message consists of lowercase English letters and ' '.
*/
package main

import (
	"fmt"
)

func decodeMessage(key string, message string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	keyAlphabet := make(map[string]string)
	number := 0
	for _, val := range key {
		letter := string(val)
		if letter == " " {
			continue
		}
		_, exists := keyAlphabet[letter]
		if !exists {
			keyAlphabet[letter] = alphabet[number]
			number++
		}
	}
	result := ""
	for _, val := range message {
		letter := string(val)
		if letter == " " {
			result = result + " "
			continue
		}
		result = result + keyAlphabet[letter]
	}
	return result
}

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}

// go run ./golang/DecodeTheMessage/DecodeTheMessage.go
