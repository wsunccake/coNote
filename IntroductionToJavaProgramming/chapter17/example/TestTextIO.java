package chapter17.example;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.PrintWriter;
import java.util.Scanner;

public class TestTextIO {
    public static void main(String[] args) throws FileNotFoundException {
        try (PrintWriter output = new PrintWriter("temp.txt")) {
            output.print("Java 101");
            output.close();
        }

        Scanner input = new Scanner(new File("temp.txt"));
        System.out.println(input.nextLine());
    }
}
