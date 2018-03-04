package hmhzzm1;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        float[] arr = new float[5];
        float sum = 0;
        float avg = 0;

        for (int i = 0; i < arr.length; i++) {
            arr[i] = sc.nextFloat();
            sum = sum + arr[i];
        }

        avg = sum/arr.length;
        System.out.printf("입력받은 숫자의 평균은 : " + avg);
    }
}
