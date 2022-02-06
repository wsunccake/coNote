package chapter9.exercise;

public class F {
    private int i;
    private static double k = 0;

    F() {
        this(5);
    }

    F(int f) {
        setI(f);
    }

    public void setI(int i){
        this.i = i;
    }

    public void setJ(int j) {
        this.i = j;
    }

    public int get() {
        return this.i;
    }

    public static void main(String[] args) {
        F f1 = new F(10);
        System.out.println(f1.get());
        f1.setI(1);
        System.out.println(f1.get());
        f1.setJ(3);
        System.out.println(f1.get());
    }
}
