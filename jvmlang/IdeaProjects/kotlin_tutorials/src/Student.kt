data class Student(val id: Int, var name: String, var email: String)


fun main(args: Array<String>) {
    val student = Student(1, "김현진", "kklath@naver.com");
    print(student)
}