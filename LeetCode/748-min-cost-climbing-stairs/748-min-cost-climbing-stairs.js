/**
 * @param {number[]} cost
 * @return {number}
 */
// var minCostClimbingStairs = function (cost) {
//     let n = cost.length;
//     for (let i = 2; i < n; i++) {
//         cost[i] = (Math.min(cost[i - 1], cost[i - 2]) + cost[i]);
//     }
//     return Math.min(cost[n - 1], cost[n - 2]);
// };


var minCostClimbingStairs = function (cost) {
    let n = cost.length;
    for (let i = 2; i < n; i++) {
        m = cost[i - 1] > cost[i - 2] ? cost[i - 2] : cost[i - 1];
        cost[i] = m + cost[i];
    }
    return cost[n - 1] > cost[n - 2] ? cost[n - 2] : cost[n - 1];
};

let input;
let answer;

input = [10, 15, 20];
answer = 15;
console.assert(minCostClimbingStairs(input) === answer, `${input} != ${answer} Fail`);

input = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1];
answer = 6;
console.assert(minCostClimbingStairs(input) === answer, `${input} != ${answer} Fail`);
