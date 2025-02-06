use std::collections::HashMap;
struct Solution {}

impl Solution {
    // // slow
    // pub fn is_anagram(s: String, t: String) -> bool {
    //     let l1: usize = s.len();
    //     let l2: usize = t.len();
    //     if l1 != l2 {
    //         return false;
    //     }

    //     let mut char_map: HashMap<u8, i32> = HashMap::new();
    //     let mut n: usize;
    //     let mut c: u8;

    //     n = 0;
    //     while n < l1 {
    //         c = s.as_bytes()[n];
    //         if char_map.contains_key(&c) {
    //             char_map.insert(c, *char_map.get(&c).unwrap() + 1);
    //         } else {
    //             char_map.insert(c, 1);
    //         }

    //         n += 1;
    //     }
    //     // println!("{:?}", char_map);

    //     n = 0;
    //     while n < l2 {
    //         c = t.as_bytes()[n];
    //         if char_map.contains_key(&c) {
    //             char_map.insert(c, *char_map.get(&c).unwrap() - 1);
    //         } else {
    //             return false;
    //         }

    //         n += 1;
    //     }

    //     for (_k, v) in char_map.into_iter() {
    //         if v != 0 {
    //             return false;
    //         }
    //     }

    //     return true;
    // }

    // // normal
    // pub fn is_anagram(s: String, t: String) -> bool {
    //     if s.len() != t.len() {
    //         return false;
    //     }

    //     let s = s.as_bytes();
    //     let t = t.as_bytes();

    //     let mut counter = std::collections::HashMap::new();

    //     for c in s {
    //         *counter.entry(c).or_insert(0) += 1;
    //     }

    //     for c in t {
    //         if !counter.contains_key(c) {
    //             return false;
    //         }

    //         *counter.get_mut(c).unwrap() -= 1;
    //     }

    //     for (_, v) in counter {
    //         if v != 0 {
    //             return false;
    //         }
    //     }

    //     true
    // }

    pub fn is_anagram(s: String, t: String) -> bool {
        let mut s_map = HashMap::new();
        let mut t_map = HashMap::new();

        s.chars().for_each(|c| {
            s_map
                .entry(c)
                .and_modify(|counter| *counter += 1)
                .or_insert(1);
        });
        t.chars().for_each(|c| {
            t_map
                .entry(c)
                .and_modify(|counter| *counter += 1)
                .or_insert(1);
        });
        s_map == t_map
    }
}

fn main() {
    let inputs1: Vec<&str> = vec!["anagram", "rat"];
    let inputs2: Vec<&str> = vec!["nagaram", "car"];
    let outputs: Vec<bool> = vec![true, false];

    for (pos, _) in outputs.iter().enumerate() {
        let ans = Solution::is_anagram(inputs1[pos].to_string(), inputs2[pos].to_string());

        if ans != outputs[pos] {
            println!(
                "{:?}, {:?}, {:?}, {:?}",
                inputs1[pos], inputs2[pos], outputs[pos], ans
            );
        }
    }
}
