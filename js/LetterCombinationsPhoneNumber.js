/*
    17. Letter Combinations of a Phone Number

    Given a string containing digits from 2-9 inclusive, return all possible letter combinations
    that the number could represent. Return the answer in any order.
    
    A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1
    does not map to any letters.
 */
var letterCombinations = function(digits) {
    const map = getMap();
    let res = [];

    let selected = [];
    for (let i = 0; i < digits.length; i++) {
        if (i === 0) {
            for (let j = 0; j < map[digits[i]].length; j++) {
                selected.push(map[digits[i]][j]);
            }
        } else {
            const len = selected.length;
            for (let j = 0; j < len; j++) {
                if (selected[j] === undefined) {
                    continue;
                }
                for (let k = 0; k < map[digits[i]].length; k++) {
                    selected.push(selected[j] + map[digits[i]][k]);
                }
                delete(selected[j]);
            }
        }
    }
    
    for (let i = 0; i < selected.length; i++) {
        if (selected[i] !== undefined) {
            res.push(selected[i]);
        }
    }
    return res;
};

function getMap() {
    return {
        1: [''],
        2: ['a','b','c'],
        3: ['d','e','f'],
        4: ['g','h','i'],
        5: ['j','k','l'],
        6: ['m','n','o'],
        7: ['p','q','r','s'],
        8: ['t','u','v'],
        9: ['w','x','y','z'],
    };
}

console.log(letterCombinations("23"));
console.log(letterCombinations("2"));
console.log(letterCombinations("234"));

// node ./js/LetterCombinationsPhoneNumber.js
