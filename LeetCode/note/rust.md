# rust

## hashmap

```rust
map1 == map2
```

```rust
let mut char_map = std::collections::HashMap::new();
let s: String

// slow
if char_map.contains_key(&c) {
    char_map.insert(c, *char_map.get(&c).unwrap() + 1);
} else {
    char_map.insert(c, 1);
}

// fast
*char_map.entry(c).or_insert(0) += 1;
```

```rust
let mut char_map = std::collections::HashMap::new();

// slow
for (_k, v) in char_map.into_iter() {
    if v != 0 {
        return false;
    }
}

// fast
for (_, v) in counter {
    if v != 0 {
        return false;
    }
}
```

---

## hashmap and vector

```rust
use std::collections::HashMap;

// slow
let nums: Vec<i32> = vec![]
let mut num_map: HashMap<i32, i32> = HashMap::new();

for n in nums {
    if num_map.contains_key(&n) {
        ...
    }
}

// fast
let nums: Vec<i32> = vec![]
let mut num_map: HashMap<&i32, i32> = HashMap::new();

for n in &nums {
    if num_map.contains_key(n) {
        ...
    }
}
```

---

## hashmap and string

```rust
// slow
let mut char_map = std::collections::HashMap::new();
let mut n: usize = 0;
let mut c: u8;
let s: String;

while n < len(s) {
    c = s.as_bytes()[n];
    if char_map.contains_key(&c) {
        char_map.insert(c, *char_map.get(&c).unwrap() + 1);
    } else {
        char_map.insert(c, 1);
    }
    n += 1;
}

// normal
let mut char_map = std::collections::HashMap::new();
let mut c: &u8;
let s: String;

for c in s.as_bytes() {
    *char_map.entry(c).or_insert(0) += 1;
}

// normal
let s: String;
let mut char_map: HashMap<char, i32> = HashMap::new();

s.chars().for_each(|c| {
    char_map
        .entry(c)
        .and_modify(|counter| *counter += 1)
        .or_insert(1);
});
```

---

## string

str: immutable sequence of UTF-8 bytes of dynamic length

String: dynamic heap string type

```rust

let s0: &str = "Hello, world!";

let mut s1 = String::from("hello");
let s2: &String = &s1;
```

---

## vector

```rust
vec1 == vec2
```

```rust
for i in 0..vec1.len() {
    println!("{:?}", i);
}

for v in vec1.iter() {
    println!("{:?}", v);
}

for (pos, value) in vec1.iter().enumerate() {
        println!("{:?}, {:?}", pos, value);
}
```

```rust
let mut dp: Vec<Vec<i32>> = vec![vec![0; n]; m];
```
