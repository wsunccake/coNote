struct Solution {}

impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let mut raw_map = std::collections::HashMap::new();
        let n = matrix.len();
        let mut i = 0;
        let mut j = 0;

        while i < n {
            j = 0;
            while j < n {
                raw_map.insert(format!("{}_{}", i, j), matrix[i][j]);
                j += 1;
            }
            i += 1;
        }

        i = 0;
        while i < n {
            j = 0;
            while j < n {
                matrix[i][j] = *raw_map.get(&format!("{}_{}", n - j - 1, i)).unwrap();
                j += 1;
            }
            i += 1;
        }
    }
}

fn main() {
    let ref1: &mut Vec<Vec<i32>> = &mut vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]];
    let ref2: &mut Vec<Vec<i32>> = &mut vec![
        vec![5, 1, 9, 11],
        vec![2, 4, 8, 10],
        vec![13, 3, 6, 7],
        vec![15, 14, 12, 16],
    ];
    let mut inputs: Vec<&mut Vec<Vec<i32>>> = vec![ref1, ref2];
    let outputs: Vec<Vec<Vec<i32>>> = vec![
        vec![vec![7, 4, 1], vec![8, 5, 2], vec![9, 6, 3]],
        vec![
            vec![15, 13, 2, 5],
            vec![14, 3, 4, 1],
            vec![12, 6, 8, 9],
            vec![16, 7, 10, 11],
        ],
    ];

    for pos in 0..outputs.len() {
        Solution::rotate(inputs[pos]);
        println!("{:?}, {:?}", inputs[pos], outputs[pos]);
        // if inputs[pos] != outputs[pos] {
        //     println!("{:?}, {:?}, {:?}", &inputs[pos], outputs[pos], ans);
        // }
    }
}
