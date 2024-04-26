struct Pointer {
    x: usize,
    y: usize,
}

struct Solution {}

impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let mut zero_point_vec = vec![];
        let m = matrix.len();
        let n = matrix[0].len();

        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == 0 {
                    zero_point_vec.push(Pointer { x: i, y: j });
                }
            }
        }

        for p in zero_point_vec {
            for i in 0..m {
                matrix[i][p.y] = 0;
            }
            for j in 0..n {
                matrix[p.x][j] = 0;
            }
        }
    }
}

fn main() {
    let m1: &mut Vec<Vec<i32>> = &mut vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]];
    let m2: &mut Vec<Vec<i32>> = &mut vec![vec![0, 1, 2, 0], vec![3, 4, 5, 2], vec![1, 3, 1, 5]];
    let mut inputs: Vec<&mut Vec<Vec<i32>>> = vec![m1, m2];
    let outputs: Vec<Vec<Vec<i32>>> = vec![
        vec![vec![1, 0, 1], vec![0, 0, 0], vec![1, 0, 1]],
        vec![vec![0, 0, 0, 0], vec![0, 4, 5, 0], vec![0, 3, 1, 0]],
    ];

    for pos in 0..inputs.len() {
        print!("{:?} ", inputs[pos]);
        Solution::set_zeroes(inputs[pos]);
        if inputs[pos] != &outputs[pos] {
            println!("{:?}, {:?}", inputs[pos], outputs[pos]);
        }
    }
}
