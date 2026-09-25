/*
205 Isomorphic Strings
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
Constraints:

    1 <= s.length <= 5 * 104
    t.length == s.length
    s and t consist of any valid ascii character.
*/
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function(s, t) {
    const mapS = Object.fromEntries(s.split('').map((char, i) => [char, t[i]]));
    const mapT = Object.fromEntries(t.split('').map((char, i) => [char, s[i]]));
    const reconstructedT = s.split('').map(char => mapS[char]).join('');
    const reconstructedS = t.split('').map(char => mapT[char]).join('');
    return reconstructedT === t && reconstructedS === s;
};

console.log(isIsomorphic("egg", "add"));
console.log(isIsomorphic("foo", "bar"));
console.log(isIsomorphic("paper", "title"));
console.log(isIsomorphic("bbbaaaba", "aaabbbba"));

// node ./js/IsomorphicStrings.js
