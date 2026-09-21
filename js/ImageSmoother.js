/*
661 Image Smoother
An image smoother is a filter of the size 3 x 3 that can be applied to each cell of an image by rounding down the average
of the cell and the eight surrounding cells. If one or more of the surrounding cells of a cell is not present, we do not
consider it in the average. Given an m x n integer matrix img representing the grayscale of an image, return the image
after applying the smoother on each cell of it.
Constraints:
    m == img.length
    n == img[i].length
    1 <= m, n <= 200
    0 <= img[i][j] <= 255
*/
/**
 * @param {number[][]} img
 * @return {number[][]}
 */
var imageSmoother = function(img) {
    let result = [];
    for (let i = 0; i < img.length; i++) {
        result[i] = [];
        let prevRow = img[i - 1] ?? [];
        let nextRow = img[i + 1] ?? [];

        for (let j = 0; j < img[i].length; j++) {
            let a1 = img[i][j-1] ?? null;
            let a2 = img[i][j+1] ?? null;
            let a3 = prevRow[j-1] ?? null;
            let a4 = prevRow[j+1] ?? null;
            let a5 = nextRow[j-1] ?? null;
            let a6 = nextRow[j+1] ?? null;
            let a7 = prevRow[j] ?? null;
            let a8 = nextRow[j] ?? null;
            const arr = [a1, a2, a3, a4, a5, a6, a7, a8];
            let count = 1;
            let sum = img[i][j];
            for (let item of arr) {
                if (item !== null) {
                    count++;
                }
                sum += item ?? 0;
            }
            result[i][j] = Math.floor(sum / count);
        }
    }
    return result;
};

console.log(imageSmoother([[1,1,1],[1,0,1],[1,1,1]]));
console.log(imageSmoother([[100,200,100],[200,50,200],[100,200,100]]));
console.log(imageSmoother([[2,3,4],[5,6,7],[8,9,10],[11,12,13],[14,15,16]]));

// node ./js/ImageSmoother.js
