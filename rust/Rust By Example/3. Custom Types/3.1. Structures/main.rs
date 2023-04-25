// An attribute to hide warnings for unused code.
#![allow(dead_code)]

#[derive(Debug)]
struct Person {
    name: String,
    age: u8,
}

// A unit struct
struct Unit;

// A tuple struct
#[derive(Debug)]
struct Pair(i32, f32);

// A struct with two fields
#[derive(Debug)]
struct Point {
    x: f32,
    y: f32,
}

// Structs can be reused as fields of another struct
#[derive(Debug)]
struct Rectangle {
    // A `` can be specified by where the top left and bottom right
    // corners are in space.
    top_left: Point,
    bottom_right: Point,
}

fn rect_area(rectangle: Rectangle) -> f32 {
    let Rectangle {
        top_left: Point { x: x1, y: y1 },
        bottom_right: Point { x: x2, y: y2 },
    } = rectangle;
    let x = x1 - x2;
    let y = y1 - y2;
    let area = x * y;
    area.abs()
}

fn square(point: Point, length: f32) -> Rectangle {
    let p = Point {
        x: point.x + length,
        y: point.y + length,
    };
    Rectangle {
        top_left: point,
        bottom_right: p,
    }
}

fn main() {
    // Create struct with field init shorthand
    let name = String::from("Peter");
    let age = 27;
    let peter = Person { name, age };

    // Print debug struct
    println!("{:?}", peter);

    // Instantiate a `Point`
    let point: Point = Point { x: 10.3, y: 0.4 };

    // Access the fields of the point
    println!("point coordinates: ({}, {})", point.x, point.y);

    // Make a new point by using struct update syntax to use the fields of our
    // other one
    let bottom_right = Point { x: 5.2, ..point };

    // `bottom_right.y` will be the same as `point.y` because we used that field
    // from `point`
    println!("second point: ({}, {})", bottom_right.x, bottom_right.y);

    // Destructure the point using a `let` binding
    let Point {
        x: left_edge,
        y: top_edge,
    } = point;

    let _rectangle = Rectangle {
        // struct instantiation is an expression too
        top_left: Point {
            x: left_edge,
            y: top_edge,
        },
        bottom_right: bottom_right,
    };

    // Instantiate a unit struct
    let _unit = Unit;

    // Instantiate a tuple struct
    let pair = Pair(1, 0.1);

    // Access the fields of a tuple struct
    println!("pair contains {:?} and {:?}", pair.0, pair.1);

    // Destructure a tuple struct
    let Pair(integer, decimal) = pair;

    println!("pair contains {:?} and {:?}", integer, decimal);

    let rect = Rectangle {
        top_left: Point { x: 0.0, y: 0.0 },
        bottom_right: Point { x: 1.0, y: 1.0 },
    };
    println!("rectangle area {}", rect_area(rect));

    let s: Rectangle = square(Point { x: 2.0, y: 2.0 }, 3.0);
    println!("rectangle: {:?}", s);
}
