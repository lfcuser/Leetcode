/*
1472 Design Browser History

You have a browser of one tab where you start on the homepage and you can visit another url, get back in the history number of steps or move forward in the history number of steps.

Implement the BrowserHistory class:

    BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
    void visit(string url) Visits url from the current page. It clears up all the forward history.
    string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps. Return the current url after moving back in history at most steps.
    string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x, you will forward only x steps. Return the current url after forwarding in history at most steps.

Constraints:

    1 <= homepage.length <= 20
    1 <= url.length <= 20
    1 <= steps <= 100
    homepage and url consist of  '.' or lower case English letters.
    At most 5000 calls will be made to visit, back, and forward.
*/
class BrowserHistory {
    constructor(homepage) {
        this.homepage = homepage;

        this.history = [this.homepage];
        this.current = 0;
        this.last = 0;
    };

    visit(url) {
        if (this.current !== this.last) {
            this.history.splice(this.current + 1, this.history.length);
        }

        this.history.push(url);
        this.last = this.history.length - 1;
        this.current = this.last;
    }

    back(steps) {
        let index = this.current - steps;
        if (index < 0) {
            index = 0;
        }
        this.current = index;
        return this.history[index];
    }

    forward(steps) {
        let index = this.current + steps;
        if (index > this.last) {
            index = this.last;
        }
        this.current = index;
        return this.history[index];
    }
}

var obj = new BrowserHistory("leetcode.com");
obj.visit("google.com");
obj.visit("facebook.com");
obj.visit("youtube.com");

console.log(obj.back(1));
console.log(obj.back(1));
console.log(obj.forward(1));
obj.visit("linkedin.com");
console.log(obj.forward(2));
console.log(obj.back(2));
console.log(obj.back(7));

// node ./js/DesignBrowserHistory.js
