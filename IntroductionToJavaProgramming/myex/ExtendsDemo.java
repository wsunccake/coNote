package myex;

public class ExtendsDemo {
    public static void main(String[] args) {

        System.out.println("a1 status:");
        Animal a1 = new Animal();
        a1.voice();
        a1.run();

        System.out.println("d1 status:");
        Dog d1 = new Dog();
        d1.voice();
        d1.run();

        System.out.println("c1 status:");
        Cat c1 = new Cat();
        c1.voice();
        c1.run();
        c1.jump();

//      up casting
        System.out.println("");

        System.out.println("a2 status:");
        Animal a2 = new Dog();
        a2.voice();
        a2.run();

        System.out.println("a3 status:");
        Animal a3 = new Cat();
        a3.voice();
        a3.run();
//        a3.jump(); // parent class no child class method
        ((Cat) a3).jump();

//      down casting
        System.out.println("");

        System.out.println("d2 status:");
//        Dog d2 = new Animal(); // Compile
//        d2.voice();

        System.out.println("d3 status:");
//        Dog d3 = (Dog) new Animal(); // RRTI
//        c2.voice();

//      first up casting, last down casting
        System.out.println("");

        System.out.println("c4 status");
        Animal a4 = new Cat();
        Cat c4 = (Cat) a4;
        c4.voice();
        c4.run();
        c4.jump();

        System.out.println("c5 status");
//        Cat c5 = new Dog();
    }

}

class Animal {

    public void voice () { System.out.println("Animal voice"); }

    public void run () { System.out.println("Animal run"); }
}


class Dog extends Animal {
    public void run () { System.out.println("Dog run"); }

    public void voice () { System.out.println("Dog voice"); }
}

class Cat extends Animal {
    public void run () { System.out.println("Cat run"); }

    public void jump () { System.out.println("Cat jump"); }
}