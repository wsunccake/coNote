# 4. variable binding

```rust
fn main() {
    let an_integer = 1u32;
    let a_boolean = true;
    let unit = ();

    // copy `an_integer` into `copied_integer`
    let copied_integer = an_integer;

    println!("An integer: {:?}", copied_integer);
    println!("A boolean: {:?}", a_boolean);
    println!("Meet the unit value: {:?}", unit);

    // variable name with an prefix underscore can be silenced (compiler warns about unused)
    let _unused_variable = 3u32;

    // let noisy_unused_variable = 2u32;
    // FIXME ^ Prefix with an underscore to suppress the warning
}
```

## 4.1 mutability

```rust
fn main() {
    let _immutable_binding = 1;
    let mut mutable_binding = 1;

    println!("Before mutation: {}", mutable_binding);

    // Ok
    mutable_binding += 1;

    println!("After mutation: {}", mutable_binding);

    // error! cannot assign new value to immutable variable
    // _immutable_binding += 1;
}
```

## 4.2 scope and shadowing

```rust
fn main() {
    let long_lived_binding = 1;

    {
        let short_lived_binding = 2;

        println!("inner short: {}", short_lived_binding);
    }

    // error! `short_lived_binding` doesn't exist in this scope
    // println!("outer short: {}", short_lived_binding);

    println!("outer long: {}", long_lived_binding);
}
```

```rust
fn main() {
    let shadowed_binding = 1;

    {
        println!("before being shadowed: {}", shadowed_binding);

        let shadowed_binding = "abc";

        println!("shadowed in inner block: {}", shadowed_binding);
    }
    println!("outside inner block: {}", shadowed_binding);

    let shadowed_binding = 2;
    println!("shadowed in outer block: {}", shadowed_binding);
}
```

## 4.3 declare first

```rust
fn main() {
    // declare variable binding
    let a_binding;

    {
        let x = 2;

        // initialize binding
        a_binding = x * x;
    }

    println!("a binding: {}", a_binding);

    let another_binding;

    // error! use of uninitialized binding
    // println!("another binding: {}", another_binding);

    another_binding = 1;

    println!("another binding: {}", another_binding);
}
```

## 4.4 freezing

```rust
fn main() {
    let mut _mutable_integer = 7i32;

    {
        // shadowing by immutable `_mutable_integer`
        let _mutable_integer = _mutable_integer;

        // error! `_mutable_integer` is frozen in this scope
        // _mutable_integer = 50;
        // `_mutable_integer` goes out of scope
    }

    // Ok! `_mutable_integer` is not frozen in this scope
    _mutable_integer = 3;
}
```
