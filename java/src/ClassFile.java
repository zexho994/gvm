
public class ClassFile implements MyInterface {

    public static final boolean FLAG = true;
    public static final byte BYTE = 123;
    public static final char X = 'X';
    public static final short SHORT = 12345;
    public final int INT = 120;
    public static long LONG = 12345678901L;
    public float PI = 3.14f;
    public static final double E = 2.71828;
    public Object obj = new Object();

    public static void main(String[] args) throws RuntimeException {
        System.out.println("Hello, World!");
    }

    public void method1(){
    }

    public void method2(int i){
    }

    public int method3(Object obj){
        return 100;
    }

    public Object method4(MyObject mc){
        return null;
    }

}
