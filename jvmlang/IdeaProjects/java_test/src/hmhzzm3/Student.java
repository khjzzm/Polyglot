package hmhzzm3;

public class Student {
    String name;
    int grade;
    String gender;

    public Student(String name, int grade, String gender) {
        this.name = name;
        this.grade = grade;
        this.gender = gender;
    }

    public void upgrade(){
        grade++;
    }

}
