package chapter19.example;

public class Max {
    public static Comparable max(Comparable o1, Comparable o2) {
        if (o1.compareTo(o2) > 0)
            return o1;
        else
            return o2;
    }

//    public static void main(String[] args) {
//        Max.max("Welcome", 23);
//    }
}
