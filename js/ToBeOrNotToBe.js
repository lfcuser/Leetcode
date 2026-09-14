/*
2704 To Be Or Not To Be

Write a function expect that helps developers test their code.
It should take in any value val and return an object with the following two functions.
toBe(val) accepts another value and returns true if the two values === each other.
If they are not equal, it should throw an error "Not Equal".
notToBe(val) accepts another value and returns true if the two values !== each other.
If they are equal, it should throw an error "Equal".
*/
class Hamlet {
    constructor(val) {
        this.val1 = val;
    }
    toBe(val) {
        if (this.val1 === val) {
            return true;
        }
        throw new Error("Not Equal");
    }
    notToBe(val) {
        if (this.val1 !== val) {
            return true;
        }
        throw new Error("Equal");
    }
}
/**
 * @param {string} val
 * @return {Object}
 */
var expect = function(val) {
    return new Hamlet(val);
};

console.log(expect(5).toBe(5), true);
try {
    console.log(expect(5).notToBe(5), "throws Equal");
} catch (e) {
    console.log(e);
}

// node ./js/ToBeOrNotToBe.js
