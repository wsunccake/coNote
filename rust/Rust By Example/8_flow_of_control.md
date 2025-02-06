# 8. flow of control

## 8.1 if/else

```rust
fn main() {
    let n = 5;

    if n < 0 {
        print!("{} is negative", n);
    } else if n > 0 {
        print!("{} is positive", n);
    } else {
        print!("{} is zero", n);
    }

    let big_n = if n < 10 && n > -10 {
        println!(", and is a small number, increase ten-fold");
        10 * n // expression return `i32`
    } else {
        println!(", and is a big number, halve the number");
        n / 2 // expression return `i32`
    };
    // `let` bindings need it, semicolon here

    println!("{} -> {}", n, big_n);
}
```

## 8.2 loop

```rust
fn main() {
    let mut count = 0u32;

    println!("Let's count until infinity!");

    // infinite loop
    loop {
        count += 1;

        if count == 3 {
            println!("three");

            // skip rest iteration
            continue;
        }

        println!("{}", count);

        if count == 5 {
            println!("OK, that's enough");

            // exit loop
            break;
        }
    }
}
```

### 8.2.1 nesting and label

```rust
#![allow(unreachable_code, unused_labels)]

fn main() {
    'outer: loop {
        println!("Entered the outer loop");

        'inner: loop {
            println!("Entered the inner loop");

            // break only the inner loop
            //break;

            // breaks the outer loop
            break 'outer;
        }

        println!("This point will never be reached");
    }

    println!("Exited the outer loop");
}
```

### 8.2.2 returning from loop

```rust
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    assert_eq!(result, 20);
}
```

## 8.3 while

```rust
fn main() {
    // counter variable
    let mut n = 1;

    // loop while `n` is less than 101
    while n < 101 {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }

        // increment counter
        n += 1;
    }
}
```

## 8.4 for

`for and range`

```rust
fn main() {
    // `n` take the values: 1, 2, ..., 100 in each iteration
    for n in 1..101 {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }
    }

    // `n` take the values: 1, 2, ..., 100 in each iteration
    for n in 1..=100 {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }
    }
}
```

`for and iterator`

```rust
fn for_iter() {
    let names = vec!["Bob", "Frank", "Ferris"];
    for name in names.iter() {
        match name {
            &"Ferris" => println!("There is a rustacean among us!"),
            _ => println!("Hello {}", name),
        }
    }

    for name in names.iter() {
        match *name {
            "Ferris" => println!("There is a rustacean among us!"),
            _ => println!("Hello {}", name),
        }
    }

    println!("names: {:?}", names);
}

fn for_iter_mut() {
    let mut names = vec!["Bob", "Frank", "Ferris"];
    for name in names.iter_mut() {
        *name = match name {
            &mut "Ferris" => "There is a rustacean among us!",
            _ => "Hello",
        }
    }
    println!("names: {:?}", names);
}

fn for_into_iter() {
    let names = vec!["Bob", "Frank", "Ferris"];

    for name in names.into_iter() {
        match name {
            "Ferris" => println!("There is a rustacean among us!"),
            _ => println!("Hello {}", name),
        }
    }

    // error; value borrowed here after move
    // println!("names: {:?}", names);
}

fn main() {
    for_iter();
    for_iter_mut();
    for_into_iter();
}
```

## 8.5 match

```rust
fn main() {
    let number = 13;
    // different values for `number`

    println!("Tell me about {}", number);
    match number {
        // match single value
        1 => println!("One!"),
        // match several value
        2 | 3 | 5 | 7 | 11 => println!("This is a prime"),
        // match inclusive range
        13..=19 => println!("A teen"),
        // handle rest of case
        _ => println!("Ain't special"),
    }

    let boolean = true;
    // match expression too
    let binary = match boolean {
        false => 0,
        true => 1,
    };

    println!("{} -> {}", boolean, binary);
}
```

### 8.5.1 destructuring

#### 8.5.1.1 destructuring tuple

```rust
fn main() {
    let triple = (0, -2, 3);
    // different values for `triple`

    println!("Tell me about {:?}", triple);
    // match can be used to destructure tuple
    match triple {
        // destructure second and third element
        (0, y, z) => println!("First is `0`, `y` is {:?}, and `z` is {:?}", y, z),
        (1, ..) => println!("First is `1` and the rest doesn't matter"),
        (.., 2) => println!("last is `2` and the rest doesn't matter"),
        (3, .., 4) => println!("First is `3`, last is `4`, and the rest doesn't matter"),
        // `..` can be used to ignore the rest of the tuple
        _ => println!("It doesn't matter what they are"),
        // `_` means don't bind the value to a variable
    }
}
```

#### 8.5.1.2 destructuring array and slice

```rust
fn main() {
    let array = [1, -2, 6];
    // change value in the array, or make slice

    match array {
        // dind second and third elements to the respective variable
        [0, second, third] => println!("array[0] = 0, array[1] = {}, array[2] = {}", second, third),

        // single value can be ignored with _
        [1, _, third] => println!(
            "array[0] = 1, array[2] = {} and array[1] was ignored",
            third
        ),

        // bind some and ignore the rest
        [-1, second, ..] => println!(
            "array[0] = -1, array[1] = {} and all the other ones were ignored",
            second
        ),

        // code below would not compile
        // [-1, second] => ...
        [3, second, tail @ ..] => println!(
            "array[0] = 3, array[1] = {} and the other elements were {:?}",
            second, tail
        ),

        [first, middle @ .., last] => println!(
            "array[0] = {}, middle = {:?}, array[2] = {}",
            first, middle, last
        ),
    }
}
```

-[@ binding](https://rustwiki.org/en/rust-by-example/flow_control/match/binding.html)

-[@ identifier pattern](https://doc.rust-lang.org/reference/patterns.html#identifier-patterns)

#### 8.5.1.3 destructuring enum

```rust
#[allow(dead_code)]
enum Color {
    // specified solely name
    Red,
    Blue,
    Green,
    // likewise tie `u32` tuple to different name: color model
    RGB(u32, u32, u32),
    HSV(u32, u32, u32),
    HSL(u32, u32, u32),
    CMY(u32, u32, u32),
    CMYK(u32, u32, u32, u32),
}

fn main() {
    let color = Color::RGB(122, 17, 40);
    // different variant for `color`

    println!("What color is it?");
    // `enum` can be destructured using `match`.
    match color {
        Color::Red => println!("The color is Red!"),
        Color::Blue => println!("The color is Blue!"),
        Color::Green => println!("The color is Green!"),
        Color::RGB(r, g, b) => println!("Red: {}, green: {}, and blue: {}!", r, g, b),
        Color::HSV(h, s, v) => println!("Hue: {}, saturation: {}, value: {}!", h, s, v),
        Color::HSL(h, s, l) => println!("Hue: {}, saturation: {}, lightness: {}!", h, s, l),
        Color::CMY(c, m, y) => println!("Cyan: {}, magenta: {}, yellow: {}!", c, m, y),
        Color::CMYK(c, m, y, k) => println!(
            "Cyan: {}, magenta: {}, yellow: {}, key (black): {}!",
            c, m, y, k
        ),
    }
}
```

#### 8.5.1.4 destructuring pointer / ref

- dereferencing use `*`

- destructuring use `&`, `ref`, `ref mut`

```rust
fn main() {
    // assign reference of type `i32`
    // `&` signifies reference being assigned.
    let reference = &4;

    match reference {
        &val => println!("Got a value via destructuring: {:?}", val),
    }

    match *reference {
        val => println!("Got a value via dereferencing: {:?}", val),
    }

    let _not_a_reference = 3;
    let ref _is_a_reference = 3;

    let value = 5;
    let mut mut_value = 6;

    // `ref` create reference
    match value {
        ref r => println!("Got a reference to a value: {:?}", r),
    }

    // `ref mut` similarly
    match mut_value {
        ref mut m => {
            *m += 10;
            println!("We added 10. `mut_value`: {:?}", m);
        }
    }
}
```

#### 8.5.1.5 destructuring dtructure

```rust
fn main() {
    struct Foo {
        x: (u32, u32),
        y: u32,
    }

    // change value struct
    let foo = Foo { x: (1, 2), y: 3 };

    match foo {
        Foo { x: (1, b), y } => println!("First of x is 1, b = {},  y = {} ", b, y),
        Foo { y: 2, x: i } => println!("y is 2, i = {:?}", i),
        Foo { y, .. } => println!("y = {}, we don't care about x", y),
        // this will give an error: pattern does not mention field `x`
        // Foo { y } => println!("y = {}", y),
    }

    let faa = Foo { x: (1, 2), y: 3 };

    let Foo { x: x0, y: y0 } = faa;
    println!("Outside: x0 = {x0:?}, y0 = {y0}");
}
```

### 8.5.2 guard

```rust
#[allow(dead_code)]
enum Temperature {
    Celsius(i32),
    Fahrenheit(i32),
}

fn main() {
    let temperature = Temperature::Celsius(35);
    match temperature {
        Temperature::Celsius(t) if t > 30 => println!("{}C is above 30 Celsius", t),
        // `if condition` part ^ is guard
        Temperature::Celsius(t) => println!("{}C is below 30 Celsius", t),

        Temperature::Fahrenheit(t) if t > 86 => println!("{}F is above 86 Fahrenheit", t),
        Temperature::Fahrenheit(t) => println!("{}F is below 86 Fahrenheit", t),
    }

    let number: u8 = 4;
    match number {
        i if i == 0 => println!("Zero"),
        i if i > 0 => println!("Greater than zero"),
        _ => unreachable!("Should never happen."),
    }
}
```

### 8.5.3 binding

```rust
// function `age` return `u32`.
fn age() -> u32 {
    15
}

fn some_number() -> Option<u32> {
    Some(42)
}

fn main() {
    println!("Tell me what type of person you are");
    match age() {
        0 => println!("I haven't celebrated my first birthday yet"),
        // `match` 1 ..= 12
        n @ 1..=12 => println!("I'm a child of age {:?}", n),
        n @ 13..=19 => println!("I'm a teen of age {:?}", n),
        // nothing bound, return result
        n => println!("I'm an old person of age {:?}", n),
    }

    match some_number() {
        // `Some` variant, match if value, bound to `n` is equal to 42
        Some(n @ 42) => println!("The Answer: {}!", n),
        // match any other number
        Some(n) => println!("Not interesting... {}", n),
        // match anything else (`None` variant)
        _ => (),
    }
}
```

## 8.6 if let

```rust
#![allow(unused)]

fn main() {
    // `optional` of type `Option<i32>`
    let optional = Some(7);

    match optional {
        Some(i) => {
            println!("This is a really long string and `{:?}`", i);
        }
        _ => {}
    };
}
```

```rust
fn main() {
    // type `Option<i32>`
    let number = Some(7);
    let letter: Option<i32> = None;
    let emoticon: Option<i32> = None;

    // `if let` construct reads: "if `let` destructures `number` into `Some(i)`
    if let Some(i) = number {
        println!("Matched {:?}!", i);
    }

    if let Some(i) = letter {
        println!("Matched {:?}!", i);
    } else {
        println!("Didn't match a number. Let's go with a letter!");
    }

    let i_like_letters = false;

    if let Some(i) = emoticon {
        println!("Matched {:?}!", i);
    } else if i_like_letters {
        println!("Didn't match a number. Let's go with a letter!");
    } else {
        println!("I don't like letters. Let's go with an emoticon :)!");
    }
}
```

```rust
enum Foo {
    Bar,
    Baz,
    Qux(u32),
}

fn main() {
    // variable
    let a = Foo::Bar;
    let b = Foo::Baz;
    let c = Foo::Qux(100);

    // match Foo::Bar
    if let Foo::Bar = a {
        println!("a is foobar");
    }

    // not match Foo::Bar, print nothing
    if let Foo::Bar = b {
        println!("b is foobar");
    }

    // match Foo::Qux has value
    // similar Some()
    if let Foo::Qux(value) = c {
        println!("c is {}", value);
    }

    // binding works with `if let`
    if let Foo::Qux(value @ 100) = c {
        println!("c is one hundred");
    }
}
```

```rust
enum Foo {
    Bar,
}

fn main() {
    let a = Foo::Bar;

    // variable match Foo::Bar
    if let Foo::Bar = a {
        println!("a is foobar");
    }
}
```

## 8.7 let else

```rust
use std::str::FromStr;

fn get_count_item(s: &str) -> (u64, &str) {
    let mut it = s.split(' ');
    let (Some(count_str), Some(item)) = (it.next(), it.next()) else {
        panic!("Can't segment count item pair: '{s}'");
    };
    let Ok(count) = u64::from_str(count_str) else {
        panic!("Can't parse integer: '{count_str}'");
    };
    (count, item)
}

fn main() {
    assert_eq!(get_count_item("3 chairs"), (3, "chairs"));
}
```

## 8.8 while else

```rust
#![allow(unused)]

fn main() {
    // `optional` type `Option<i32>`
    let mut optional = Some(0);

    loop {
        match optional {
            // `optional` destructure, evaluate block
            Some(i) => {
                if i > 9 {
                    println!("Greater than 9, quit!");
                    optional = None;
                } else {
                    println!("`i` is `{:?}`. Try again.", i);
                    optional = Some(i + 1);
                }
            }
            // quit loop when destructure fails
            _ => {
                break;
            }
        }
    }
}
```

```rust
fn main() {
    // `optional` type `Option<i32>`
    let mut optional = Some(0);

    // "while `let` destructure `optional` into
    // `Some(i)`, evaluate block (`{}`). Else `break`.
    while let Some(i) = optional {
        if i > 9 {
            println!("Greater than 9, quit!");
            optional = None;
        } else {
            println!("`i` is `{:?}`. Try again.", i);
            optional = Some(i + 1);
        }
    }
}
```
