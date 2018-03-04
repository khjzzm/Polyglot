package hmhzzm3;

public class UniversityStudent extends Student{
    String major;
    University UI;

    public UniversityStudent(String name, int grade, String gender, String major, University UI) {
        super(name, grade, gender);
        this.major = major;
        this.UI = UI;
    }

    @Override
    public String toString() {
        return "UniversityStudent{" +
                "major='" + major + '\'' +
                ", UI=" + UI +
                ", name='" + name + '\'' +
                ", grade=" + grade +
                ", gender='" + gender + '\'' +
                '}';
    }
}
