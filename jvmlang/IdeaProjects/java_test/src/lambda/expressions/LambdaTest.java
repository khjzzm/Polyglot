package lambda.expressions;

import java.lang.FunctionalInterface;

public class LambdaTest {
    public static void main(String[] args) {
        Func add = (int a, int b) -> a+b;
    }
}

@FunctionalInterface
interface Func {
    public int calc(int a, int b);
}
