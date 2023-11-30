use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();
        let mut goal: i32;
        let mut results: Vec<i32> = Vec::from([0, 0]);

        let mut i = 0;
        for n in &nums {
            goal = target - *n;
            if map.contains_key(&goal) {
                results[0] = *map.get(&goal).unwrap();
                results[1] = i;
                break;
            } else {
                map.insert(*n, i);
            }
            i += 1;
        }

        results
    }
}

fn main() {
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    let ans1 = 0;
    let ans2 = 1;

    let a = Solution::two_sum(nums, target);
    assert_eq!(a[0], ans1);
    assert_eq!(a[1], ans2);

    let nums = vec![3, 2, 4];
    let target = 6;
    let ans1 = 1;
    let ans2 = 2;

    let a = Solution::two_sum(nums, target);
    assert_eq!(a[0], ans1);
    assert_eq!(a[1], ans2);

    let nums = vec![3, 3];
    let target = 6;
    let ans1 = 0;
    let ans2 = 1;

    let a = Solution::two_sum(nums, target);
    assert_eq!(a[0], ans1);
    assert_eq!(a[1], ans2);
}
