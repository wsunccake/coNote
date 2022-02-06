package chapter9.example;

public class SimpleCircle2 {
    public static void main(String[] args) {
        SimpleCircle2 circle = new SimpleCircle2();
        System.out.println("The area of the circle of radius " + circle.radius + " is " + circle.getArea());
        SimpleCircle2 circle2 = new SimpleCircle2(25);
        System.out.println("The area of the circle of radius" + circle2.radius + " is " + circle2.getArea());
        SimpleCircle2 circle3 = new SimpleCircle2(125);
        System.out.println("The area of the circle of radius" + circle3.radius + " is " + circle3.getArea());
        circle2.radius = 100;
        System.out.println("The area of the circle of radius" + circle2.radius + " is " + circle2.getArea());
    }

    double radius ;

    SimpleCircle2(){
        radius = 1;
    }
    SimpleCircle2(double newRadius){
        radius = newRadius;
    }
    double getArea(){
        return radius*radius*Math.PI;
    }
    double getPerimeter(){
        return 2*radius*Math.PI;
    }
    void  setRadius(double newRadius){
        radius = newRadius;
    }
}