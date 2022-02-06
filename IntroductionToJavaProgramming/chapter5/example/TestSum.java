package chapter5.example;

public class TestSum {
    public static void main(String[] args) {
        float sum = 0;

        for (float i = 0.01f; i <= 1.0f; i = i + 0.01f)
            sum += i;
        System.out.println("The sum is " + sum);

        sum = 0.0f;
        double currentValue = 0.01;
        for (int count = 0; count < 100; count++) {
            sum += currentValue;
            currentValue += 0.01;
        }
        System.out.println("The sum is " + sum);
    }
}
