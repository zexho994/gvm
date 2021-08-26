public class Person<T>{

    T name;

    public static void main(String[] args) {
        Person<String> p = new Person<String>();
        p.name = "张三";
        GvmOut.to(p.name);
    }
}