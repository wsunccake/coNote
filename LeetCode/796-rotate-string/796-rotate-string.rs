struct Solution {}

impl Solution {
    pub fn rotate_string(s: String, goal: String) -> bool {
        let m = s.len();
        let n = goal.len();

        if m > n {
            return false;
        }
        let double_s = format!("{s}{s}");
        return double_s.contains(&goal);
    }
}

fn main() {
    let s10 = &String::from("abcde");
    let s11 = &"abcde".to_string();
    let s12 = &"aa".to_owned();
    let inputs1 = vec![s10, s11, s12];

    let s20 = &"cdeab".to_string();
    let s21 = &"abced".to_string();
    let s22 = &"a".to_string();
    let inputs2 = vec![s20, s21, s22];

    let outputs = vec![true, false, false];

    for pos in 0..outputs.len() {
        let ans = Solution::rotate_string(inputs1[pos].clone(), inputs2[pos].clone());

        if ans != outputs[pos] {
            println!(
                "{:?}, {:?}, {:?}, {:?}",
                inputs1[pos], inputs2[pos], ans, outputs[pos]
            );
        }
    }
}
