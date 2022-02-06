package chapter19.example;

import chapter13.example.Circle;
import chapter13.example.GeometricObject;
import chapter13.example.Rectangle;

public class BoundedTypeDemo {
    public static void main(String[] args) {
        Rectangle rectangle = new Rectangle(2, 2);
        Circle circle = new Circle(2);

        System.out.println("Same area? " + equalArea(rectangle, circle));
    }

    public static <E extends GeometricObject> boolean equalArea( E object1, E object2) {
        return object1.getArea() == object2.getArea();
    }
}
