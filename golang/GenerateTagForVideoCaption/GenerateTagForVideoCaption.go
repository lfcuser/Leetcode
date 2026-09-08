/*
3582 Generate Tag for Video Caption

You are given a string caption representing the caption for a video.

The following actions must be performed in order to generate a valid tag for the video:

	Combine all words in the string into a single camelCase string prefixed with '#'. A camelCase string is one where the first letter of all words except the first one is capitalized. All characters after the first character in each word must be lowercase.

	Remove all characters that are not an English letter, except the first '#'.

	Truncate the result to a maximum of 100 characters.

Return the tag after performing the actions on caption.

Constraints:

	1 <= caption.length <= 150
	caption consists only of English letters and ' '.
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func generateTag(caption string) string {
	reg := regexp.MustCompile(`[^a-zA-Z ]`)
	caption = reg.ReplaceAllString(caption, "")
	caption = strings.ToLower(caption)

	buf := strings.Fields(caption)
	length := len(buf)

	res := "#"

	if length > 0 {
		res += buf[0]
	}

	for i := 1; i < length; i++ {
		res = res + strings.ToUpper(buf[i][:1]) + buf[i][1:]
	}

	length = len(res)
	if length > 100 {
		res = res[0:100]
	}

	return res
}

func main() {
	fmt.Println(generateTag("Leetcode daily streak achieved"))
	fmt.Println(generateTag("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
	fmt.Println(generateTag("   "))
}

// go run ./golang/GenerateTagForVideoCaption/GenerateTagForVideoCaption.go
