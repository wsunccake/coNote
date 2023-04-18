fn main() {
    println!("{} days", 31);
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

    // named argument
    println!(
        "{subject} {verb} {object}",
        object = "the lazy dog",
        subject = "the quick brown fox",
        verb = "jumps over"
    );

    // specify format character
    println!("Base 10:               {}", 69420); // 69420
    println!("Base 2 (binary):       {:b}", 69420); // 10000111100101100
    println!("Base 8 (octal):        {:o}", 69420); // 207454
    println!("Base 16 (hexadecimal): {:x}", 69420); // 10f2c
    println!("Base 16 (hexadecimal): {:X}", 69420); // 10F2C

    // justify or adjust
    println!("{number:0}", number = 1); // left,"1"
    println!("{number:0<5}", number = 1); // left, "10000"
    println!("{number:>5}", number = 1); // right, "    1"
    println!("{number:0>width$}", number = 1, width = 5); // right, "00001"

    // number argument
    println!("My name is {0}, {1} {0}", "Bond", "James");

    // implement fmt::Display
    #[allow(dead_code)]
    struct Structure(i32);
    use std::fmt;
    impl fmt::Display for Structure {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            write!(f, "Structure: {}", self.0)
        }
    }
    println!("This struct `{}` won't print...", Structure(3));

    let number: f64 = 1.0;
    let width: usize = 5;
    println!("{number:>width$}");

    let pi = 3.141592;
    println!("Pi is roughly {:.3}", pi);
}
