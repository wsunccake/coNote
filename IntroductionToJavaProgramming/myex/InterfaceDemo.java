package myex;

public class InterfaceDemo {
    public static void main(String[] args) {

        Plane p1 = new Plane();
        p1.fly();

        Bird b1 = new Bird();
        b1.fly();
        b1.voice();

        IFly i1 = new Plane();
        i1.fly();
        ((ITurbo) i1).turbo();
        ((Plane) i1).turbo();

        IFly i2 = new Bird();
        i2.fly();
//        i2.voice();
        ((Animal) i2).run();
        ((Bird) i2).run();

        IFly i3 = new IFly() {
            @Override
            public void fly() { System.out.println("IFly fly"); }
        };
        i3.fly();
    }
}

interface IFly {
    void fly();
}

interface ITurbo {
    void turbo();
}

class Plane extends Vehicle implements IFly, ITurbo{

    @Override
    public void fly() { System.out.println("Plane fly"); }

    @Override
    public void run() { System.out.println("Plane run"); }

    @Override
    public void turbo() { System.out.println("Plane turbo"); }
}

class Bird extends Animal implements IFly {

    @Override
    public void fly() { System.out.println("Bird fly"); }

    public void run() { System.out.println("Bird run");}
}