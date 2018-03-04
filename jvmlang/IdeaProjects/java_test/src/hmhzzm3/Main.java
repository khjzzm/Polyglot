package hmhzzm3;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);


        System.out.print("대학교이름:");
        String uname = sc.nextLine();
        System.out.print("대학교장소");
        String ulocation = sc.nextLine();
        University obj = new University(uname,ulocation);

        System.out.print("이름:");
        String name = sc.nextLine();
        System.out.print("학년:");
        int grade = sc.nextInt();sc.nextLine();
        System.out.print("성별:");
        String gender = sc.nextLine();
        System.out.print("전공:");
        String major = sc.nextLine();

        UniversityStudent hmh = new UniversityStudent(name,grade,gender,major,obj);

        System.out.println(hmh);
    }
}
