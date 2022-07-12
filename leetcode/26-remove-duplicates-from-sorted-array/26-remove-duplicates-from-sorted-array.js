/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  let count = 0;
  let current = null;

  for (let i = 0; i < nums.length; i++) {
    if (current !== nums[i]) {
      nums[count] = nums[i];
      count++;
      current = nums[i];
    }
  }
  return count;
};


// test function
const test = function(q, a) {
  const sol = removeDuplicates(q);
  console.assert(sol === a.length, `Fail: sol ${sol}, length ${a.length}`);
  for (let i = 0; i < sol; i++) {
    console.assert(a[i] === q[i], `Fail: ${a[i]} !== ${q[i]}`);
  }
};


const q1 = [1, 1, 2];
const a1 = [1, 2];
test(q1, a1);

const q2 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
const a2 = [0, 1, 2, 3, 4];
test(q2, a2);
