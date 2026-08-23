/*
2129 Capitalize the Title

You are given a string title consisting of one or more words separated by a single space, where each word consists of English letters. Capitalize the string by changing the capitalization of each word such that:

    If the length of the word is 1 or 2 letters, change all letters to lowercase.
    Otherwise, change the first letter to uppercase and the remaining letters to lowercase.

Return the capitalized title.

Constraints:

    1 <= title.length <= 100
    title consists of words separated by a single space without any leading or trailing spaces.
    Each word consists of uppercase and lowercase English letters and is non-empty.
*/

/**
 * @param {string} title
 * @return {string}
 */
var capitalizeTitle = function(title) {
    title = title.toLowerCase().split(" ");
    let result = "";

    for (let i = 0; i < title.length; i++) {
        if (title[i].length <= 2) {
            result += title[i] + " ";
            continue;
        }
        result += title[i][0].toUpperCase() + title[i].slice(1) + " ";
    }

    return result.trim();
};

console.log(capitalizeTitle("capiTalIze tHe titLe"));
console.log(capitalizeTitle("First leTTeR of EACH Word"));
console.log(capitalizeTitle("i lOve leetcode"));

// node ./js/CapitalizetheTitle.js