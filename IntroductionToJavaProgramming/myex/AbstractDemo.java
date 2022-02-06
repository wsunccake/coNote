package myex;

public class AbstractDemo {
    public static void main(String[] args) {

//       Vehicle v0 = new Vehicle(); // abstract class can't new instance

        System.out.println("v1 status:");
        Vehicle v1 = new Vehicle() {
            @Override
            public void run() { System.out.println("anonymous Vehicle run"); }
        };
        v1.voice();
        v1.run();

        System.out.println("c1 status:");
        Car c1 = new Car();
        c1.voice();
        c1.run();

        System.out.println("s1 status:");
        Sport s1 = new Sport();
        s1.voice();
        s1.run();

        System.out.println("t1 status:");
        Truck t1 = new Truck();
        t1.voice();
        t1.run();

//      abstract class up casting
        System.out.println("");

        System.out.println("v2 status:");
        Vehicle v2 = new Car();
        v2.voice();
        v2.run();

        System.out.println("v3 status:");
        Vehicle v3 = new Sport();
        v3.voice();
        v3.run();

        System.out.println("v4 status:");
        Vehicle v4 = new Truck();
        v4.voice();
        v4.run();

//      class up casting
        System.out.println("");

        System.out.println("c2 status:");
        Car c2 = new Sport();
        c2.voice();
        c2.run();

        System.out.println("c3 status:");
        Car c3 = new Truck();
        c3.voice();
        c3.run();
//      c3.concatenate();
        ((Truck) c3).concatenate();


    }
}

abstract class Vehicle {
    public void voice () { System.out.println("Vehicle voice"); }

    public abstract void run ();
}

class Car extends Vehicle {

    @Override
    public void run() { System.out.println("Car run"); }
}


class Sport extends Car {

    @Override
    public void run() { System.out.println("Sport run"); }

    @Override
    public void voice() { System.out.println("Sport voice"); }
}

class Truck extends Car {

    public void concatenate () { System.out.println("Truck concatenate"); }
}