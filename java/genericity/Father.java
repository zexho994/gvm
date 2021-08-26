public class Father<T> {
    T t;

    public static void main(String[] args) {
        Father<? super Son> f = new Father<Son>();
    }
}