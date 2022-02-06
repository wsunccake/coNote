package chapter9.exercise;

public class WhichWrong {
    int i = 5;
    static int k = 2;

    public static void main(String[] args) {
        int j = k;
        m2(1,2);
        WhichWrong a = new WhichWrong();
        a.m1();
    }
    public void m1(){
        i = i + k + m2(i, k);
    }
    public static int m2(int i, int j){
        return (int)(Math.pow(i,j));
    }
}
