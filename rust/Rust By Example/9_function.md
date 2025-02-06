# 9. function

```rust
// unlike C/C++, there's no restriction on the order of function
fn main() {
    fizzbuzz_to(100);
}

// function return boolean value
fn is_divisible_by(lhs: u32, rhs: u32) -> bool {
    // corner case, early return
    if rhs == 0 {
        return false;
    }

    // expression, `return` keyword is not necessary
    lhs % rhs == 0
}

// function "don't" return value, actually return the unit type `()`
fn fizzbuzz(n: u32) -> () {
    if is_divisible_by(n, 15) {
        println!("fizzbuzz");
    } else if is_divisible_by(n, 3) {
        println!("fizz");
    } else if is_divisible_by(n, 5) {
        println!("buzz");
    } else {
        println!("{}", n);
    }
}

// function return `()`, return type can be omitted from signature
fn fizzbuzz_to(n: u32) {
    for n in 1..=n {
        fizzbuzz(n);
    }
}
```

## 9.1 method

```rust
struct Point {
    x: f64,
    y: f64,
}

// implementation block, `Point` associated function & method
impl Point {
    fn origin() -> Point {
        Point { x: 0.0, y: 0.0 }
    }

    fn new(x: f64, y: f64) -> Point {
        Point { x: x, y: y }
    }
}

struct Rectangle {
    p1: Point,
    p2: Point,
}

impl Rectangle {
    // `&self` is sugar for `self: &Self`
    // `Self` is type of caller object
    // in the case `Self` = `Rectangle`
    fn area(&self) -> f64 {
        // `self` access struct fields via the dot operator
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

        ((x1 - x2) * (y1 - y2)).abs()
    }

    fn perimeter(&self) -> f64 {
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

        2.0 * ((x1 - x2).abs() + (y1 - y2).abs())
    }

    // `&mut self` desugars to `self: &mut Self`
    fn translate(&mut self, x: f64, y: f64) {
        self.p1.x += x;
        self.p2.x += x;

        self.p1.y += y;
        self.p2.y += y;
    }
}

struct Pair(Box<i32>, Box<i32>);

impl Pair {
    fn destroy(self) {
        let Pair(first, second) = self;
        println!("Destroying Pair({}, {})", first, second);
    }
}

fn main() {
    let rectangle = Rectangle {
        p1: Point::origin(),
        p2: Point::new(3.0, 4.0),
    };

    println!("Rectangle perimeter: {}", rectangle.perimeter());
    println!("Rectangle area: {}", rectangle.area());

    let mut square = Rectangle {
        p1: Point::origin(),
        p2: Point::new(1.0, 1.0),
    };

    // error! `rectangle` is immutable
    // rectangle.translate(1.0, 0.0);

    square.translate(1.0, 1.0);

    let pair = Pair(Box::new(1), Box::new(2));

    pair.destroy();

    // error! `destroy` call "consumed" `pair`
    // pair.destroy();
}
```

## 9.2 closure

```rust
fn main() {
    let outer_var = 42;

    // error! regular function can't refer to variable in the enclosing environment
    // fn function(i: i32) -> i32 { i + outer_var }

    // closure is anonymous, here binding them to references
    let closure_annotated = |i: i32| -> i32 { i + outer_var };
    let closure_inferred = |i| i + outer_var;

    // call the closure
    println!("closure_annotated: {}", closure_annotated(1));
    println!("closure_inferred: {}", closure_inferred(1));
    // error! closure's type has been inferred
    // println!(
    //     "cannot reuse closure_inferred with another type: {}",
    //     closure_inferred(42i64)
    // );

    // closure take no arguments and return `i32`
    let one = || 1;
    println!("closure returning one: {}", one());
}
```

### 9.2.1 capturing

Closures capture variable

- by reference: `&T`

- by mutable reference: `&mut T`

- by value: `T`

```rust
fn closure_fn() {
    let color = String::from("green");
    let print = || println!("`color`: {}", color);
    print();

    let _reborrow = &color;
    print();

    let _color_moved = color;
    // error! borrow `color`
    // print();
}

fn closure_fn_mut() {
    let mut count = 0;
    let mut inc = || {
        count += 1;
        println!("`count`: {}", count);
    };

    inc();

    // error! closure mutably borrows `count`
    // let _reborrow = &count;
    inc();

    let _count_reborrowed = &mut count;
}

fn closure_fn_once() {
    use std::mem;
    let movable = Box::new(3);

    let consume = || {
        println!("`movable`: {:?}", movable);
        mem::drop(movable);
    };

    consume();

    // error! only one caller
    // consume();
}

fn closure_fn_move() {
    let haystack = vec![1, 2, 3];

    let contains = move |needle| haystack.contains(needle);

    println!("{}", contains(&1));
    println!("{}", contains(&4));

    // println!("There're {} elements in vec", haystack.len());
    // compile-time error: because borrow checker doesn't allow re-using variable after it
    // has been moved.
}

fn main() {
    closure_fn();
    closure_fn_mut();
    closure_fn_once();
    closure_fn_move();
}
```

### 9.2.2 as input parameter

```rust
// <F> denote that F is a "Generic type parameter"
fn apply<F>(f: F)
where
    // closure take no input and returns nothing
    F: FnOnce(),
{
    // ^ TODO: try changing this to `Fn` or `FnMut`.

    f();
}

fn apply_to_3<F>(f: F) -> i32
where
    F: Fn(i32) -> i32,
{
    f(3)
}

fn main() {
    use std::mem;

    let greeting = "hello";
    let mut farewell = "goodbye".to_owned();

    let diary = || {
        println!("I said {}.", greeting);

        farewell.push_str("!!!");
        println!("Then I screamed {}.", farewell);
        println!("Now I can sleep. zzzzz");

        mem::drop(farewell);
    };

    apply(diary);

    // `double` satisfies `apply_to_3`'s trait bound
    let double = |x| 2 * x;

    println!("3 doubled: {}", apply_to_3(double));
}
```

### 9.2.3 type anonymity

```rust
// `F` must be generic.
fn apply<F>(f: F)
where
    F: FnOnce(),
{
    f();
}
```

```rust
fn apply<F>(f: F)
where
    F: Fn(),
{
    f();
}

fn main() {
    let x = 7;
    let print = || println!("{}", x);

    apply(print);
}
```

### 9.2.4 input function

```rust
fn call_me<F: Fn()>(f: F) {
    f();
}

fn function() {
    println!("I'm a function!");
}

fn main() {
    let closure = || println!("I'm a closure!");

    call_me(closure);
    call_me(function);
}
```

### 9.2.5 as output parameter

```rust
fn create_fn() -> impl Fn() {
    let text = "Fn".to_owned();

    move || println!("This is a: {}", text)
}

fn create_fnmut() -> impl FnMut() {
    let text = "FnMut".to_owned();

    move || println!("This is a: {}", text)
}

fn create_fnonce() -> impl FnOnce() {
    let text = "FnOnce".to_owned();

    move || println!("This is a: {}", text)
}

fn main() {
    let fn_plain = create_fn();
    let mut fn_mut = create_fnmut();
    let fn_once = create_fnonce();

    fn_plain();
    fn_mut();
    fn_once();
}
```

### 9.2.6 example in std

#### 9.2.6.1 Iterator::any

```rust
pub trait Iterator {
    type Item;

    fn any<F>(&mut self, f: F) -> bool where
        F: FnMut(Self::Item) -> bool;
}
```

```rust
fn main() {
    let vec1 = vec![1, 2, 3];
    let vec2 = vec![4, 5, 6];

    println!("2 in vec1: {}", vec1.iter().any(|&x| x == 2));
    println!("2 in vec2: {}", vec2.into_iter().any(|x| x == 2));

    println!("vec1 len: {}", vec1.len());
    println!("First element of vec1 is: {}", vec1[0]);

    let array1 = [1, 2, 3];
    let array2 = [4, 5, 6];

    println!("2 in array1: {}", array1.iter().any(|&x| x == 2));
    println!("2 in array2: {}", array2.into_iter().any(|x| x == 2));
}
```

#### 9.2.6.2 Searching through iterators

```rust
pub trait Iterator {
    type Item;

    fn find<P>(&mut self, predicate: P) -> Option<Self::Item> where
        P: FnMut(&Self::Item) -> bool;
}
```

```rust
fn main() {
    let vec1 = vec![1, 2, 3];
    let vec2 = vec![4, 5, 6];

    let mut iter = vec1.iter();
    let mut into_iter = vec2.into_iter();

    println!("Find 2 in vec1: {:?}", iter.find(|&&x| x == 2));
    println!("Find 2 in vec2: {:?}", into_iter.find(|&x| x == 2));

    let array1 = [1, 2, 3];
    let array2 = [4, 5, 6];

    println!("Find 2 in array1: {:?}", array1.iter().find(|&&x| x == 2));
    println!(
        "Find 2 in array2: {:?}",
        array2.into_iter().find(|&x| x == 2)
    );

    let vec = vec![1, 9, 3, 3, 13, 2];

    let index_of_first_even_number = vec.iter().position(|&x| x % 2 == 0);
    assert_eq!(index_of_first_even_number, Some(5));

    let index_of_first_negative_number = vec.into_iter().position(|x| x < 0);
    assert_eq!(index_of_first_negative_number, None);
}
```

---

## 9.3 higher order function

```rust
fn is_odd(n: u32) -> bool {
    n % 2 == 1
}

fn main() {
    println!("Find the sum of all the numbers with odd squares under 1000");
    let upper = 1000;

    // imperative approach
    let mut acc = 0;
    // iterate: 0, 1, 2, ... to infinity
    for n in 0.. {
        let n_squared = n * n;

        if n_squared >= upper {
            // break loop if exceeded the upper limit
            break;
        } else if is_odd(n_squared) {
            // accumulate value, if it's odd
            acc += n_squared;
        }
    }
    println!("imperative style: {}", acc);

    // functional approach
    let sum_of_squared_odd_numbers: u32 = (0..)
        .map(|n| n * n)
        .take_while(|&n_squared| n_squared < upper)
        .filter(|&n_squared| is_odd(n_squared))
        .sum();
    println!("functional style: {}", sum_of_squared_odd_numbers);
}
```

## 9.4 diverging function

```rust
fn main() {
    fn sum_odd_numbers(up_to: u32) -> u32 {
        let mut acc = 0;
        for i in 0..up_to {
            let addition: u32 = match i % 2 == 1 {
                true => i,
                false => continue,
            };
            acc += addition;
        }
        acc
    }
    println!(
        "Sum of odd numbers up to 9 (excluding): {}",
        sum_odd_numbers(9)
    );
}
```
