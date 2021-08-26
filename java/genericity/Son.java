public class Son<T extends Person<String>>{

    T father;

    public static void main(String[] args) {
        Son<Person<String>> son = new Son<>();
        son.father = new Person();
        son.father.name = "李四";
        GvmOut.to(son.father.name);
    }

    public T name(T t){
        return t;
    }
}