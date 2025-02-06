struct Solution {}

impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        let l1: usize = s1.len();
        let l2: usize = s2.len();
        if l1 > l2 {
            return false;
        }

        let mut s1_vec: Vec<i32> = vec![];
        let mut s2_vec: Vec<i32> = vec![];

        for _i in 0..26 {
            s1_vec.push(0);
            s2_vec.push(0);
            // println!("{:?}, {:?}, {:?}", _i, s1_vec.get(_i), s2_vec.get(_i));
        }

        let mut n: usize = 0;
        while n < l1 {
            s1_vec[s1.as_bytes()[n] as usize - 'a' as usize] += 1;
            s2_vec[s2.as_bytes()[n] as usize - 'a' as usize] += 1;
            n += 1;
        }

        if s1_vec == s2_vec {
            return true;
        }
        // println!("s1_vec: {:?}, s2_vec: {:?}", s1_vec, s2_vec);

        let mut pos: usize = l1;
        let mut start: usize = 0;
        while pos < l2 {
            s2_vec[s2.as_bytes()[pos] as usize - 'a' as usize] += 1;
            s2_vec[s2.as_bytes()[start] as usize - 'a' as usize] -= 1;
            if s1_vec == s2_vec {
                return true;
            }

            pos += 1;
            start += 1;
        }
        return false;
    }
}

fn main() {
    let inputs1: Vec<&str> = vec!["ab", "ab", "abc", "a"];
    let inputs2: Vec<&str> = vec!["eidbaooo", "eidboaoo", "ccccbbbbaaaa", "ab"];
    let outputs: Vec<bool> = vec![true, false, false, true];

    for (pos, _) in outputs.iter().enumerate() {
        let ans = Solution::check_inclusion(inputs1[pos].to_string(), inputs2[pos].to_string());

        if ans != outputs[pos] {
            println!(
                "{:?}, {:?}, {:?}, {:?}",
                inputs1[pos], inputs2[pos], outputs[pos], ans
            );
        }
    }
}
