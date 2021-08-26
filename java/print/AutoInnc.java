public class AutoInnc {
    public static void main(String[] args) {
        int i = 1;
        GvmOut.to(i++);
        GvmOut.to(++i);
    }
}