package kevinTv.modern.eight;

public class OopAnotherExample {
    public static void main(String[] args) {
        final CalculatorService calculatorService = new CalculatorService(new Addition());
        final int addtionResult = calculatorService.calculate(33,2);
        System.out.println(addtionResult);

        final CalculatorService calculatorService1 = new CalculatorService(new Substation());
        final int substationResult = calculatorService1.calculate(33,2);
        System.out.println(substationResult);

        Calculation calculation = new Addition();
        final int addtionalResult = calculation.calculate(13,3);
        System.out.println(addtionalResult);
    }
}

interface Calculation{
    int calculate(int num1, int num2);
}

class Addition implements Calculation{
    @Override
    public int calculate(int num1, int num2) {
        return num1 + num2;
    }
}

class Substation implements Calculation{
    @Override
    public int calculate(int num1, int num2) {
        return num1 - num2;
    }
}


class CalculatorService{
    private final Calculation calculation;

    public CalculatorService(Calculation calculation) {
        this.calculation = calculation;
    }

    public int calculate(int num1, int num2){
        if(num1>10 && num2 < num1){
            return calculation.calculate(num1,num2);
        }else{
            throw new IllegalArgumentException("Invalid input num1: "+num1+" num2: " +num2 );
        }
    }
}
