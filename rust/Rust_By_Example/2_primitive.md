# 2. primitive

`scalar type`

- signed integer: i8, i16, i32, i64, i128 and isize (pointer size)
- unsigned integer: u8, u16, u32, u64, u128 and usize (pointer size)
- floating point: f32, f64
- char unicode scalar value like 'a', 'α' and '∞' (4 bytes each)
- bool either true or false
- unit type (), whose only possible value is an empty tuple: ()

`compound type`

- array like [1, 2, 3]
- tuple like (1, true)

```rust
fn main() {
    // annotated type
    let logical: bool = true;
    let a_float: f64 = 1.0; // regular annotation
    let an_integer = 5i32; // suffix annotation

    // default type
    let default_float = 3.0; // `f64`
    let default_integer = 7; // `i32`

    // type can be inferred from context
    let mut inferred_type = 12; // type `i64` is inferred from another line
    inferred_type = 4294967296i64;

    // mutable variable
    let mut mutable = 12; // mutable `i32`
    mutable = 21; // change value

    // mutable = true; // type cannot change

    // variable can be overwritten with shadowing.
    let mutable = true;
}
```

## 2.1 literal and operator

```rust
fn main() {
    println!("1 + 2 = {}", 1u32 + 2); // integer addition
    println!("1 - 2 = {}", 1i32 - 2); // integer subtraction

    // println!("1 - 2 = {}", 1u32 - 2); // overflow

    // scientific notation
    println!("1e4 is {}, -2.5e-3 is {}", 1e4, -2.5e-3);

    // short-circuiting boolean logic
    println!("true AND false is {}", true && false);
    println!("true OR false is {}", true || false);
    println!("NOT true is {}", !true);

    // bitwise operation
    println!("0011 AND 0101 is {:04b}", 0b0011u32 & 0b0101);
    println!("0011 OR 0101 is {:04b}", 0b0011u32 | 0b0101);
    println!("0011 XOR 0101 is {:04b}", 0b0011u32 ^ 0b0101);
    println!("1 << 5 is {}", 1u32 << 5);
    println!("0x80 >> 2 is 0x{:x}", 0x80u32 >> 2);

    // use underscores to improve readability
    println!("One million is written as {}", 1_000_000u32);
}
```

## 2.2 tuple

```rust
use std::fmt;

fn reverse(pair: (i32, bool)) -> (bool, i32) {
    let (int_param, bool_param) = pair;

    (bool_param, int_param)
}

// activity
#[derive(Debug)]
struct Matrix(f32, f32, f32, f32);

impl fmt::Display for Matrix {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({} {})\n({} {})", self.0, self.1, self.2, self.3)
    }
}

fn transpose(m: Matrix) -> String {
    format!("({} {})\n({} {})", m.0, m.2, m.1, m.3)
}

fn main() {
    // tuple with a bunch of different types
    let long_tuple = (
        1u8, 2u16, 3u32, 4u64, -1i8, -2i16, -3i32, -4i64, 0.1f32, 0.2f64, 'a', true,
    );

    // value can be extracted from the tuple using index
    println!("Long tuple first value: {}", long_tuple.0);
    println!("Long tuple second value: {}", long_tuple.1);

    // tuple can be tuple member
    let tuple_of_tuples = ((1u8, 2u16, 2u32), (4u64, -1i8), -2i16);

    // tuple are printable
    println!("tuple of tuples: {:?}", tuple_of_tuples);

    // tuple (more than 12 elements) cannot be printed
    // let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
    // println!("Too long tuple: {:?}", too_long_tuple);

    let pair = (1, true);
    println!("Pair is {:?}", pair);

    println!("The reversed pair is {:?}", reverse(pair));

    println!("One element tuple: {:?}", (5u32,)); // tuple
    println!("Just an integer: {:?}", (5u32)); // integer

    // tuple can be destructured to create bindings
    let tuple = (1, "hello", 4.5, true);
    let (a, b, c, d) = tuple;
    println!("{:?}, {:?}, {:?}, {:?}", a, b, c, d);

    let matrix = Matrix(1.1, 1.2, 2.1, 2.2);
    println!("{:?}", matrix);
    println!("{}", matrix);
    println!("Matrix:\n{}", matrix);
    println!("Transpose:\n{}", transpose(matrix));
}
```

## 2.3 array and slice

```rust
use std::mem;

fn analyze_slice(slice: &[i32]) {
    println!("First element of the slice: {}", slice[0]);
    println!("The slice has {} elements", slice.len());
}

fn main() {
    // fixed-size array (type signature is superfluous)
    let xs: [i32; 5] = [1, 2, 3, 4, 5];

    // all elements can be initialized to the same value
    let ys: [i32; 500] = [0; 500];

    // indexing starts at 0
    println!("First element of the array: {}", xs[0]);
    println!("Second element of the array: {}", xs[1]);

    // `len` returns the count of elements in the array
    println!("Number of elements in array: {}", xs.len());

    // array are stack allocated
    println!("Array occupies {} bytes", mem::size_of_val(&xs));

    // array can be automatically borrowed as slices
    println!("Borrow the whole array as a slice.");
    analyze_slice(&xs);

    // slice can point to a section of an array
    // [starting_index..ending_index]
    // `starting_index` is the first position in the slice.
    // `ending_index` is one more than the last position in the slice.
    println!("Borrow a section of the array as a slice.");
    analyze_slice(&ys[1..4]);

    // empty slice `&[]`:
    let empty_array: [u32; 0] = [];
    assert_eq!(&empty_array, &[]);
    assert_eq!(&empty_array, &[][..]); // same but more verbose

    // array can be safely accessed using `.get`, which returns `Option`
    // this can be matched as shown below, or used with `.expect()`
    for i in 0..xs.len() + 1 {
        // Oops, one element too far!
        match xs.get(i) {
            Some(xval) => println!("{}: {}", i, xval),
            None => println!("Slow down! {} is too far!", i),
        }
    }

    // out of bound indexing on array causes compile time error
    //println!("{}", xs[5]);
    // out of bound indexing on slice causes runtime error
    //println!("{}", xs[..][5]);
}
```
