/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
// var twoSum = function(nums, target) {
//     let result = [0, 0];
//     let tmpMap = {};

//     nums.forEach(function (item, index, array) {
//         let goal = target - item;
//         if (goal.toString() in tmpMap) {
//             result = [tmpMap[goal.toString()], index];
//             return;
//         }
//         tmpMap[item.toString()] = index;

//     });

//     return result;
// };


// var twoSum = function(nums, target) {
//     const tmpMap = {};

//     for (let index = 0; index < nums.length; index++) {
//         const goal = target - nums[index];
//         if (goal in tmpMap) {
//             return [tmpMap[goal], index];
//         }
//         tmpMap[nums[index]] = index;
//     }

//     return [0, 0];
// };


var twoSum = function(nums, target) {
    const tmpMap = {};

    for (let index = 0; index < nums.length; index++) {
        const num = nums[index];
        const goal = target - num;
        const goal_index = tmpMap[goal];
        if (goal_index !== undefined) {
            return [goal_index, index];
        }
        tmpMap[num] = index;
    }

    return [0, 0];
};

let input1;
let input2;
let answer1;
let answer2;

input1 = [1, 2, 3];
input2 = 6;
answer1 = 0;
answer2 = 0;
let sol1 = twoSum(input1, input2);
console.assert(sol1[0] === answer1, `${answer1} Fail`);
console.assert(sol1[1] === answer2, `${answer2} Fail`);

input1 = [3, 2, 4];
input2 = 6;
answer1 = 1;
answer2 = 2;
let sol2 = twoSum([3, 2, 4], 6);
console.assert(sol2[0] === answer1, `${answer1} Fail`);
console.assert(sol2[1] === answer2, `${answer2} Fail`);

input1 = [3, 3];
input2 = 6;
answer1 = 0;
answer2 = 1;
let sol3 = twoSum([3, 3], 6)
console.assert(sol3[0] === answer1, `${answer1} Fail`);
console.assert(sol3[1] === answer2, `${answer2} Fail`);
