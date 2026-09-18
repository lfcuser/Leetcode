/*
1779 Find Nearest Point That Has the Same X or Y Coordinate

You are given two integers, x and y, which represent your current location on a Cartesian grid: (x, y).
You are also given an array points where each points[i] = [ai, bi] represents that a point exists at (ai, bi).
A point is valid if it shares the same x-coordinate or the same y-coordinate as your location.

Return the index (0-indexed) of the valid point with the smallest Manhattan distance from your current location.
If there are multiple, return the valid point with the smallest index. If there are no valid points, return -1.

The Manhattan distance between two points (x1, y1) and (x2, y2) is abs(x1 - x2) + abs(y1 - y2).

Constraints:
    1 <= points.length <= 104
    points[i].length == 2
    1 <= x, y, ai, bi <= 104
*/
/**
 * @param {number} x
 * @param {number} y
 * @param {number[][]} points
 * @return {number}
 */
var nearestValidPoint = function(x, y, points) {
    let minVal = Number.MAX_SAFE_INTEGER;
    let minIdx = null;
    for (let i = 0; i < points.length; i++) {
        let x2 = points[i][0];
        let y2 = points[i][1];
        if (x2 === x || y2 === y) {
            let mDistance = Math.abs(x - x2) + Math.abs(y - y2);
            if (mDistance < minVal) {
                minVal = mDistance;
                minIdx = i;
            }
        }
    }
    return minIdx !== null ? minIdx : -1;
};

console.log(nearestValidPoint(3, 4, [[1,2],[3,1],[2,4],[2,3],[4,4]]), 2);
console.log(nearestValidPoint(3, 4, [[3,4]]), 0);
console.log(nearestValidPoint(3, 4, [[2,3]]), -1);

// node ./js/FindNearestPoint.js
