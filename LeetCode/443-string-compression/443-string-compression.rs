struct Solution {}

impl Solution {
    pub fn compress(chars: &mut Vec<char>) -> i32 {
        let mut char_vec: Vec<char> = vec![chars[0]];
        let mut num_vec: Vec<i32> = vec![1];

        for i in 0..(chars.len() - 1) {
            if chars[i] == chars[i + 1] {
                num_vec[char_vec.len() - 1] += 1;
            } else {
                char_vec.push(chars[i + 1]);
                num_vec.push(1);
            }
        }

        let mut tmp_str: String = "".to_string();
        for i in 0..(char_vec.len()) {
            tmp_str += &char_vec[i].to_string();
            if num_vec[i] != 1 {
                tmp_str = tmp_str.clone() + &num_vec[i].to_string();
            }
        }

        let mut pos = 0;
        let tmp_vec: Vec<char> = tmp_str.chars().collect();
        for c in tmp_vec {
            chars[pos] = c;
            pos += 1;
        }

        return pos as i32;
    }
}

fn main() {
    let ref1 = &mut vec!['a', 'a', 'b', 'b', 'c', 'c', 'c'];
    let ref2 = &mut vec!['a'];
    let ref3 = &mut vec![
        'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b',
    ];
    let ref4 = &mut vec![
        'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'a', 'b',
    ];
    let mut inputs = vec![ref1, ref2, ref3, ref4];
    let outputs: Vec<Vec<char>> = vec![
        vec!['a', '2', 'b', '2', 'c', '3'],
        vec!['a'],
        vec!['a', 'b', '1', '2'],
        vec!['b', '1', '1', 'a', 'b'],
    ];
    let nums: Vec<i32> = vec![6, 1, 4, 5];

    for pos in 0..nums.len() {
        let ans = Solution::compress(inputs[pos]) as usize;
        if &inputs[pos][0..ans] != outputs[pos] {
            println!("{:?}, {:?}, {:?}", &inputs[pos], outputs[pos], ans);
        }
    }
}
