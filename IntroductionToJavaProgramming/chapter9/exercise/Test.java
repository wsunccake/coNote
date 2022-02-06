package chapter9.exercise;

public class Test {
    public int factorial (int n) {
        int result = 1;
        for (int i = 1; i<= n ; i++)
            result *=i;
        return result;
    }

    public static void main(String[] args) {
         Test a = new Test();
         a.factorial(3);
    }
}
