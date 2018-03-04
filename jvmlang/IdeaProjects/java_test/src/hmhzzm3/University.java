package hmhzzm3;

public class University{
    String Uname;
    String Ulocation;

    public University(String uname, String ulocation) {
        Uname = uname;
        Ulocation = ulocation;
    }

    @Override
    public String toString() {
        return "University{" +
                "Uname='" + Uname + '\'' +
                ", Ulocation='" + Ulocation + '\'' +
                '}';
    }
}
