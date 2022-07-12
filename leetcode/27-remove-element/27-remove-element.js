/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
  for (let i = nums.length - 1; i >= 0 ; i--) {
    if (nums[i] == val) {
       nums.splice(i, 1);
    }
  }
  return nums.length;
};


let test = function(q, t, a) {
  const sol = removeElement(q, t);
  console.assert(sol === a.length, `Fail: length, ${sol} !== ${a.length}`);
  for (let i = 0; i < sol; i++) {
    console.assert(q[i] === a[i], `Fail: index ${i}: ${q[i]} !== ${a[i]}`);
  }
};

const q1 = [3, 2, 2, 3];
const t1 = 3;
const a1 = [2, 2];
test(q1, t1, a1);

const q2 = [0, 1, 2, 2, 3, 0, 4, 2];
const t2 = 2;
const a2 = [0, 1, 3, 0, 4]
test(q2, t2, a2);
