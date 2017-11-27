public class SumFrom1To{

    public static void main(String[] args) {
        System.out.println(sumFrom1To(99));
        System.out.println(sumFrom1ToGaussRecursive(99));
    }

    public static int sumFrom1To(int x){
        int result = 0;
        for(int i=0; i<=x; i++){
            result += i;
        }
        return result;
    }

    public static int sumFrom1ToGauss(int x){
        int result = (1+x) * (x/2);
        return result;
    }

    public static int sumFrom1ToGaussOdd(int x){
        int result = 0;
        if(x%2==0){
            result += (1+x) * (x/2);
        }else{
            result += x * ((x-1)/2) + x;
        }
        return result;
    }

    public static int sumFrom1ToGaussRecursive(int x){
        int result = 0;
        if(x%2==0){
            result += (1+x) * (x/2);
        }else{
            result += sumFrom1ToGaussRecursive(x-1) + x;
        }
        return result;
    }

}
