# 1. hello world

```rust
// hello.rc
fn main() {
    println!("Hello World!"); // 顯示 Hello World!
}
```

```bash
linux:~ $ rustc hello.rs

linux:~ $ ./hello
```

```rust
// activity
fn main() {
    println!("Hello World!");
    println!("I'm a Rustacean!");
}
```

---

## 1.1 comment

`regular comment`

- // line comments which go to the end of the line.
- /_ block comments which go to the closing delimiter. _/

`doc comment`

- /// generate library docs for the following item.
- //! generate library docs for the enclosing item.

```rust
fn main() {
    // This is an example of a line comment.
    // There are two slashes at the beginning of the line.
    // And nothing written after these will be read by the compiler.

    // println!("Hello, world!");

    // Run it. See? Now try deleting the two slashes, and run it again.

    /*
     * This is another type of comment, a block comment. In general,
     * line comments are the recommended comment style. But block comments
     * are extremely useful for temporarily disabling chunks of code.
     * /* Block comments can be /* nested, */ */ so it takes only a few
     * keystrokes to comment out everything in this main() function.
     * /*/*/* Try it yourself! */*/*/
     */

    /*
    Note: The previous column of `*` was entirely for style. There's
    no actual need for it.
    */

    // You can manipulate expressions more easily with block comments
    // than with line comments. Try deleting the comment delimiters
    // to change the result:
    let x = 5 + /* 90 + */ 5;
    println!("Is `x` 10 or 100? x = {}", x);
}
```

---

## 1.2 formatted print

`format!`: write formatted text to String

`print!`: same as format! but the text is printed to the console (io::stdout).

`println!`: same as print! but a newline is appended.

`eprint!`: same as print! but the text is printed to the standard error (io::stderr).

`eprintln!`: same as eprint! but a newline is appended.

```rust
fn any_argument() {
    println!("{{}}: any argument");
    println!("{} days", 31);
}

fn positional_argument() {
    println!("{{n}}: positional argument");
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");
}

fn named_argument() {
    println!("{{name}}: positional argument");
    println!(
        "{subject} {verb} {object}",
        object = "the lazy dog",
        subject = "the quick brown fox",
        verb = "jumps over"
    );
}

fn carry_format() {
    println!("{{:c}}: carry format");
    println!("Base 10:               {}", 69420); // 69420
    println!("Base 2 (binary):       {:b}", 69420); // 10000111100101100
    println!("Base 8 (octal):        {:o}", 69420); // 207454
    println!("Base 16 (hexadecimal): {:x}", 69420); // 10f2c
    println!("Base 16 (hexadecimal): {:X}", 69420); // 10F2C
}

fn justify_format() {
    println!("{{:>}}: right justify");
    println!("{number:>5}", number = 1); // "    1"
    println!("{number:0>5}", number = 1); // 00001

    println!("{{:<}}: left justify");
    println!("{number:<5}", number = 1); // "1    "
    println!("{number:0<5}", number = 1); // 10000

    println!("{{name:<width$}}: named argument with justify");
    println!("{number:0>width$}", number = 1, width = 5); // 00001
}

fn no_implement_display() {
    println!("no implement display ");
    #[allow(dead_code)]
    struct Structure(i32);
    // println!("This struct `{}` won't print...", Structure(3)); // error
    // `Structure` don't implement fmt::Display
}

fn main() {
    any_argument();
    positional_argument();
    named_argument();

    carry_format();
    justify_format();

    no_implement_display();
}
```

```rust
// activity
fn main() {
    // println!("My name is {0}, {1} {0}", "Bond");
    // fix to
    println!("My name is {0}, {1} {0}", "James", "Bond");

    let pi = 3.141592;
    println!("Pi is roughly {:.3}", pi);
}
```

- [Module std::fmt](https://doc.rust-lang.org/std/fmt/)
- [macro_rules!](https://rustwiki.org/en/rust-by-example/macros.html)

### 1.2.1 debug

```rust
// 無法使用 {}, 因為 no implement fmt::Display or fmt::Debug
struct UnPrintable(i32);

// 使用 derive(Debug),  自動產生 fmt::Debug
#[derive(Debug)]
struct DebugPrintable(i32);
```

```rust
#[derive(Debug)]
struct Structure(i32);

#[derive(Debug)]
struct Deep(Structure);

#[derive(Debug)]
struct Person<'a> {
    name: &'a str,
    age: u8,
}

fn print_debug() {
    println!("print with {{:?}}");
    println!("{:?} months in a year.", 12);
    println!(
        "{1:?} {0:?} is the {actor:?} name.",
        "Slater",
        "Christian",
        actor = "actor's"
    );
}

fn print_debug_struct() {
    println!("print with struct {{:?}}");
    println!("Now {:?} will print!", Structure(3));

    // The problem with `derive` is there is no control over how
    // the results look. What if I want this to just show a `7`?
    println!("Now {:?} will print!", Deep(Structure(7)));
    println!("Now {:?} will print!", Deep(Structure(7)).0);
    println!("Now {:?} will print!", Deep(Structure(7)).0 .0);
}

fn print_debug_pretty() {
    println!("print with struct {{:#?}}");
    let name = "Peter";
    let age = 27;
    let peter = Person { name, age };

    println!("{:#?}", peter);
}

fn main() {
    print_debug();
    print_debug_struct();
    print_debug_pretty();
}
```

### 1.2.2 display

```rust
use std::fmt;

struct Structure(i32);

// implement fmt::Display
impl fmt::Display for Structure {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.0)
    }
}
```

```rust
use std::fmt;

#[derive(Debug)]
struct MinMax(i64, i64);

impl fmt::Display for MinMax {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {})", self.0, self.1)
    }
}

fn print_min_max() {
    let minmax = MinMax(0, 14);

    println!("Compare structures:");
    println!("Display: {}", minmax);
    println!("Debug: {:?}", minmax);

    let big_range = MinMax(-300, 300);
    let small_range = MinMax(-3, 3);

    println!(
        "The big range is {big} and the small is {small}",
        small = small_range,
        big = big_range
    );
}

#[derive(Debug)]
struct Point2D {
    x: f64,
    y: f64,
}

impl fmt::Display for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "x: {}, y: {}", self.x, self.y)
    }
}

impl fmt::Binary for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "x: {}, y: {}", self.x, self.y)
    }
}

fn print_point2d() {
    let point = Point2D { x: 3.3, y: 7.2 };

    println!("Compare points:");
    println!("Display: {}", point);
    println!("Debug: {:?}", point);

    // need implement fmt::Binary
    println!("What does Point2D look like in binary: {:b}?", point);
}

fn main() {
    print_min_max();
    print_point2d();
}
```

```rust
// activity
use std::fmt;

#[derive(Debug)]
struct Complex {
    real: f64,
    imag: f64,
}

impl fmt::Display for Complex {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{} + {}i", self.real, self.imag)
    }
}

fn main() {
    let c = Complex {
        real: 3.3,
        imag: 7.2,
    };
    println!("Display: {}", c);
    println!("Debug: {:?}", c);
}
```

#### 1.2.2.1 testcase: list

```rust
write!(f, "{}", value)?;
// same, not recommend =>
try!(write!(f, "{}", value));
```

- [? question mark operator](https://doc.rust-lang.org/reference/expressions/operator-expr.html#the-question-mark-operator)

```rust
use std::fmt;

struct List(Vec<i32>);

impl fmt::Display for List {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let vec = &self.0;

        write!(f, "[")?;

        for (count, v) in vec.iter().enumerate() {
            if count != 0 {
                write!(f, ", ")?;
            }
            // write!(f, "{}", v)?;
            write!(f, "{}: {}", count, v)?; // activity
        }

        write!(f, "]")
    }
}

fn main() {
    let v = List(vec![1, 2, 3]);
    println!("{}", v);
}
```

### 1.2.3 formatting

```text
format!("{}", foo) -> "3735928559"
format!("0x{:X}", foo) -> "0xDEADBEEF"
format!("0o{:o}", foo) -> "0o33653337357"
```

```rust
use std::fmt::{self, Display, Formatter};

struct City {
    name: &'static str,
    lat: f32, // latitude
    lon: f32, // longitude
}

impl Display for City {
    // `f` is a buffer, and this method must write the formatted string into it.
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        let lat_c = if self.lat >= 0.0 { 'N' } else { 'S' };
        let lon_c = if self.lon >= 0.0 { 'E' } else { 'W' };

        // `write!` is like `format!`, but it will write the formatted string
        // into a buffer (the first argument).
        write!(
            f,
            "{}: {:.3}°{} {:.3}°{}",
            self.name,
            self.lat.abs(),
            lat_c,
            self.lon.abs(),
            lon_c
        )
    }
}

#[derive(Debug)]
struct Color {
    red: u8,
    green: u8,
    blue: u8,
}

// activity
impl Display for Color {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        let rgb: u32 = (self.red as u32) * 65536 + (self.green as u32) * 256 + self.blue as u32; // RGB = (R*65536)+(G*256)+B
        write!(
            f,
            "RGB( {}, {}, {}) 0x{:06X}",
            self.red, self.green, self.blue, rgb
        )
    }
}

fn main() {
    for city in [
        City {
            name: "Dublin",
            lat: 53.347778,
            lon: -6.259722,
        },
        City {
            name: "Oslo",
            lat: 59.95,
            lon: 10.75,
        },
        City {
            name: "Vancouver",
            lat: 49.25,
            lon: -123.1,
        },
    ] {
        println!("{}", city);
    }
    for color in [
        Color {
            red: 128,
            green: 255,
            blue: 90,
        },
        Color {
            red: 0,
            green: 3,
            blue: 254,
        },
        Color {
            red: 0,
            green: 0,
            blue: 0,
        },
    ] {
        // println!("{:?}", color);
        println!("{}", color);
    }
}
```
