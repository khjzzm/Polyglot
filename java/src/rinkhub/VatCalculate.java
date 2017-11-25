package rinkhub;

public class VatCalculate {
    static final double TAX_RATE = 1.1;
    public static void main(String[] args) {
        double totalValue = 50000;          // 합계금액
        double supplyValue;                 // 공금가액
        double vat;                         // 부가가치세(세액)

        supplyValue = supplyValueCalculate(totalValue);
        vat = vatCalculate1(supplyValue);
        System.out.println("공급가액:"+supplyValue);
        System.out.println("세액1:"+vat);
        System.out.println(supplyValue + vat);
        System.out.println("=================================");

        supplyValue = supplyValueCalculate(totalValue);
        vat = vatCalculate2(totalValue, supplyValue);
        System.out.println("공급가액:"+supplyValue);
        System.out.println("세액2:"+vat);
        System.out.println(supplyValue + vat);
        System.out.println(supplyValue);
    }

    //공급가액 = 합계급액 / 1.1
    public static double supplyValueCalculate(double totalValue){
        return totalValue / TAX_RATE;
    }

    //세액1 = 공급가액 * 0.1
    public static double vatCalculate1(double supplyValue){
        return supplyValue * 0.1;
    }

    // 세액2 = 합계금액 - 공금가액
    public static double vatCalculate2(double totalValue, double supplyValue){
        return totalValue - supplyValue;
    }
}
