struct Solution {}

impl Solution {
    pub fn can_permute_palindrome(s: String) -> bool {
        let mut char_map = std::collections::HashMap::new();

        for c in s.as_bytes() {
            *char_map.entry(c).or_insert(0) += 1;
        }

        let mut odd_count = 0;

        for (_k, v) in char_map {
            if v % 2 == 1 {
                odd_count += 1;
                if odd_count > 1 {
                    return false;
                }
            }
        }
        return true;
    }
}

fn main() {
    let inputs: Vec<&str> = vec!["code", "aab", "carerac"];
    let outputs: Vec<bool> = vec![false, true, true];

    for (pos, _) in outputs.iter().enumerate() {
        let ans = Solution::can_permute_palindrome(inputs[pos].to_string());

        if ans != outputs[pos] {
            println!("{:?}, {:?}, {:?}", inputs[pos], outputs[pos], ans);
        }
    }
}
