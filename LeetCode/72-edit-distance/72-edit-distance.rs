struct Solution {}

impl Solution {
    pub fn min_distance(word1: String, word2: String) -> i32 {
        let m = word1.len();
        let n = word2.len();

        let mut dp: Vec<Vec<i32>> = vec![vec![0; n + 1]; m + 1];

        for i in 1..(m + 1) {
            dp[i][0] = i as i32;
        }

        for j in 1..(n + 1) {
            dp[0][j] = j as i32;
        }

        for i in 1..(m + 1) {
            for j in 1..(n + 1) {
                if word1.as_bytes()[i - 1] == word2.as_bytes()[j - 1] {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] =
                        std::cmp::min(std::cmp::min(dp[i - 1][j - 1], dp[i - 1][j]), dp[i][j - 1])
                            + 1;
                }
            }
        }

        return dp[m][n];
    }
}

fn main() {
    let inputs1: Vec<&str> = vec!["horse", "intention"];
    let inputs2: Vec<&str> = vec!["ros", "execution"];
    let outputs: Vec<i32> = vec![3, 5];

    for pos in 0..outputs.len() {
        let ans = Solution::min_distance(inputs1[pos].to_string(), inputs2[pos].to_string());

        if ans != outputs[pos] {
            println!(
                "{:?}, {:?}, {:?}, {:?}",
                inputs1[pos], inputs2[pos], outputs[pos], ans
            );
        }
    }
}
