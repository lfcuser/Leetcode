/*
794 Valid Tic-Tac-Toe State

Given a Tic-Tac-Toe board as a string array board, return true if and only if it is possible to reach this board position during the course of a valid tic-tac-toe game.

The board is a 3 x 3 array that consists of characters ' ', 'X', and 'O'. The ' ' character represents an empty square.

Here are the rules of Tic-Tac-Toe:

    Players take turns placing characters into empty squares ' '.
    The first player always places 'X' characters, while the second player always places 'O' characters.
    'X' and 'O' characters are always placed into empty squares, never filled ones.
    The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
    The game also ends if all squares are non-empty.
    No more moves can be played if the game is over.

Constraints:

board.length == 3
board[i].length == 3
board[i][j] is either 'X', 'O', or ' '.
*/

/**
 * @param {string[]} board
 * @return {boolean}
 */
var validTicTacToe = function(board) {
    let countX = 0;
    let countO = 0;
    let countRowsX = 0;
    let countRowsO = 0;
    let countColumnsX = 0;
    let countColumnsO = 0;
    let countDiagonalX = 0;
    let countDiagonalO = 0;

    for (let i = 0; i < board.length; i++) {
        countX += board[i].split("X").length - 1;
        countO += board[i].split("O").length - 1;
        countRowsX += board[i].split("XXX").length - 1;
        countRowsO += board[i].split("OOO").length - 1;
    }

    for (let i = 0; i < board.length; i++) {
        if (board[0][i] === board[1][i] && board[0][i] === board[2][i]) {
            if (board[0][i] === "X") {
                countColumnsX++;
            } else if (board[0][i] === "O") {
                countColumnsO++;
            }
        }
    }

    if ((board[0][0] === board[1][1] && board[1][1] === board[2][2])
        || (board[0][2] === board[1][1] && board[1][1] === board[2][0])
    ) {
        if (board[1][1] === "X") {
            countDiagonalX++;
        } else if (board[1][1] === "O") {
            countDiagonalO++;
        }
    }

    const winX = countRowsX + countColumnsX + countDiagonalX;
    const winO = countRowsO + countColumnsO + countDiagonalO;

    return !(countO > countX
        || countX - countO > 1
        || countRowsX + countRowsO > 1
        || countColumnsO + countColumnsX > 1
        || (countDiagonalX === 1 && countRowsX + countColumnsX > 0)
        || (countDiagonalO === 1 && countRowsO + countColumnsO > 0)
        || (winX > 0 && countX === countO)
        || (winO > 0 && countO !== countX)
    );
};


console.log(validTicTacToe(["O  ","   ","   "]), false);
console.log(validTicTacToe(["XOX"," X ","   "]), false);
console.log(validTicTacToe(["XOX","O O","XOX"]), true);
console.log(validTicTacToe(["OXX","XOX","OXO"]), false);
console.log(validTicTacToe(["XOX","OX ","OXO"]), true);
console.log(validTicTacToe(["X  ","X  ","O O"]), true);

// node ./js/ValidTicTacToeState.js