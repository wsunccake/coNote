# 7. expression

```rust
fn main() {
    // variable binding
    let x = 5u32;

    // expression;
    x;
    x + 1;
    15;

    let y = {
        let x_squared = x * x;
        let x_cube = x_squared * x;

        // expression will be assigned to `y`
        x_cube + x_squared + x
    };

    let z = {
        // semicolon suppresses this expression and `()` is assigned to `z`
        2 * x;
    };

    println!("x is {:?}", x);
    println!("y is {:?}", y);
    println!("z is {:?}", z);
}
```
