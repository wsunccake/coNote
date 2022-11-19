/**
 * @param {number[]} ranks
 * @param {character[]} suits
 * @return {string}
 */
var bestHand = function (ranks, suits) {
    let suitsObj = {};
    for (const key of suits) {
        if (Object.hasOwnProperty.call(suitsObj, key)) {
            suitsObj[key] += 1;
        } else {
            suitsObj[key] = 1;
        }
    }
    if (Object.keys(suitsObj).length === 1) {
        return "Flush";
    }

    let ranksObj = {};
    let m = 0;
    for (const key of ranks) {
        if (Object.hasOwnProperty.call(ranksObj, key)) {
            ranksObj[key] += 1;
        } else {
            ranksObj[key] = 1;
        }
        if (m < ranksObj[key]) {
            m = ranksObj[key];
        }
    }
    let resultObj = {
        4: "Three of a Kind",
        3: "Three of a Kind",
        2: "Pair",
    };

    let r = "High Card";
    if (Object.hasOwnProperty.call(resultObj, m)) {
        r = resultObj[m];
    }

    return r;
};

let suits;
let ranks;
let answer;
let sol;

ranks = [13, 2, 3, 1, 9];
suits = ['a', 'a', 'a', 'a', 'a'];
answer = "Flush";
sol = bestHand(ranks, suits);
console.assert(sol === answer, `${answer} Fail`);

ranks = [4, 4, 2, 4, 4];
suits = ['d', 'a', 'a', 'b', 'c'];
answer = "Three of a Kind";
sol = bestHand(ranks, suits);
console.assert(sol === answer, `${answer} Fail`);

ranks = [10, 10, 2, 12, 9];
suits = ['a', 'b', 'c', 'a', 'd'];
answer = "Pair";
sol = bestHand(ranks, suits);
console.assert(sol === answer, `${answer} Fail`);

ranks = [2, 10, 7, 10, 7];
suits = ['a', 'b', 'a', 'd', 'b'];
answer = "Pair";
sol = bestHand(ranks, suits);
console.assert(sol === answer, `${answer} Fail`);
