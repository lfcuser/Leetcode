/*
3360 Stone Removal Game

Alice and Bob are playing a game where they take turns removing stones from a pile, with Alice going first.

    Alice starts by removing exactly 10 stones on her first turn.
    For each subsequent turn, each player removes exactly 1 fewer stone than the previous opponent.

The player who cannot make a move loses the game.

Given a positive integer n, return true if Alice wins the game and false otherwise.

Example 1:

Input: n = 12

Output: true

Explanation:

    Alice removes 10 stones on her first turn, leaving 2 stones for Bob.
    Bob cannot remove 9 stones, so Alice wins.

Example 2:

Input: n = 1

Output: false

Explanation:

    Alice cannot remove 10 stones, so Alice loses.



Constraints:

    1 <= n <= 50
*/

package main

import (
	"fmt"
	"math"
)

func canAliceWin(n int) bool {
	x := int((21 - math.Sqrt(441-8*float64(n))) / 2)
	return x%2 != 0
}

func main() {
	fmt.Println(canAliceWin(12), true)
	fmt.Println(canAliceWin(2), false)
}

// go run ./golang/StoneRemovalGame/StoneRemovalGame.go

/*
Solution description

1+2+3+...+10=55 - positive
55-10-9-8-...-1=0 - negative

We want to know at which step of the negative progression we will stop if we can subtract N from the total

For example: we have 12 in total. 55 - 10 = 45, and we have 2 left to subtract. 2 < 9, then it was one step.

Formula:
n = ((a(1) + a(x)) / 2) * x
a(1) = 10
a(x) = 11 - x

=>

x^2 - 21x + 2n = 0

=>

D = 441 - 8n

=>

(21 - D^1/2)/2 === x

If x is odd, then Alice wins. because every odd move belongs to her, and the one who makes the last move wins.
*/
