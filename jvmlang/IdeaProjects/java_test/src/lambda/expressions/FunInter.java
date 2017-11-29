package lambda.expressions;


//기능인터페이스
//람다 (Lambdas)는 단 하나의 추상 메소드를 가진 인터페이스 인 기능적 인터페이스에서만 작동 할 수 있습니다.
// 기능 인터페이스에는 default 또는 static 메소드가 여러 개있을 수 있습니다.
//(이러한 이유 때문에 때때로 Single Abstract Method Interfaces 또는 SAM Interfaces라고도 함).
public class FunInter {

}

interface Foo1 {
    void bar();
}

interface Foo2 {
    int bar(boolean baz);
}

interface Foo3 {
    String bar(Object baz, int mink);
}

interface Foo4 {
    default String bar() { // default so not counted
        return "baz";
    }
    void quux();
}

@FunctionalInterface
interface Foo5 {
    void bar();
}

@FunctionalInterface
interface BlankFoo1 extends Foo3 { // inherits abstract method from Foo3

}

@FunctionalInterface
interface Foo6 {
    void bar();
    boolean equals(Object obj); // overrides one of Object's method so not counted
}

interface BadFoo {
    void bar();
    void quux(); // <-- Second method prevents lambda: which one should
    // be considered as lambda?
}

interface BlankFoo2 { }