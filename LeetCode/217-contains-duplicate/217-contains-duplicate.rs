use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn show(nums: Vec<i32>) {
        for (pos, e) in nums.iter().enumerate() {
            println!("{:?} -> {:?}", pos, e)
        }
    }

    // // slow
    // pub fn contains_duplicate(nums: Vec<i32>) -> bool {
    //     let mut is_duplicate = false;
    //     let mut num_map: HashMap<i32, i32> = HashMap::new();
    //     for n in nums {
    //         if num_map.contains_key(&n) {
    //             is_duplicate = true;
    //             break;
    //         } else {
    //             num_map.insert(n, 1);
    //         }
    //     }
    //     return is_duplicate;
    // }

    // fast
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut is_duplicate = false;
        let mut num_map: HashMap<&i32, i32> = HashMap::new();

        for n in &nums {
            if num_map.contains_key(n) {
                is_duplicate = true;
                break;
            } else {
                num_map.insert(n, 1);
            }
        }

        return is_duplicate;
    }
}
fn main() {
    let inputs = vec![
        vec![1, 2, 3, 1],
        vec![1, 2, 3, 4],
        vec![1, 1, 1, 3, 3, 4, 3, 2, 4, 2],
    ];
    let outputs = vec![true, false, true];

    for (pos, _) in inputs.iter().enumerate() {
        Solution::show(inputs[pos].to_vec());
        let ans = Solution::contains_duplicate(inputs[pos].to_vec());

        if ans != outputs[pos] {
            println!("{:?}, {:?}, {:?}", inputs[pos], outputs[pos], ans);
        }
    }
}
