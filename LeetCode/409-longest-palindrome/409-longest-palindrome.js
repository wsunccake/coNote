/**
 * @param {string} s
 * @return {number}
 */
 var longestPalindrome = function (s) {
    let stringObj = {};
    for (const iterator of s) {
        if (stringObj[iterator] === undefined) {
            stringObj[iterator] = 1;
        } else {
            stringObj[iterator] += 1;
        }
    }

    let r = 0;
    let q = 0;
    for (const key in stringObj) {
        r += parseInt(stringObj[key] / 2);
        if (stringObj[key] % 2 == 1) {
            q = 1
        }
    }

    return r * 2 + q
};


// var longestPalindrome = function (s) {
//     let stringObj = {};
//     let iterator;
//     for (const iterator of s) {
//         if (Object.hasOwnProperty.call(stringObj, iterator)) {
//             stringObj[iterator] += 1;
//         } else {
//             stringObj[iterator] = 1;
//         }
//     }

//     let r = 0;
//     let q = 0;
//     let key;
//     for (key in stringObj) {
//         r += parseInt(stringObj[key] / 2);
//         if (stringObj[key] % 2 == 1) {
//             q = 1
//         }
//     }

//     return r * 2 + q
// };

let input;
let answer;
let sol;

input = "abccccdd";
answer = 7;
sol = longestPalindrome(input);
console.assert(sol === answer, `${answer} Fail`);

input = "a";
answer = 1;
sol = longestPalindrome(input);
console.assert(sol === answer, `${answer} Fail`);
