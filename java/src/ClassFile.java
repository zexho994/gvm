
public class ClassFile extends FatherFile implements MyInterface {

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
        ClassFile cf = new ClassFile();
        FatherFile cf2 = new FatherFile();
        FatherFile cf3 = new ClassFile();
        cf.method1();
        cf.method2(111);
        int res3 = cf.method3(new Object());
        Object o = cf.method4(new MyObject());
        int resf2 = cf.father_method2(1);
        cf2.father_method1();
        int resfff2 = cf3.father_method2(3);
        System.out.println("res3 = " + res3);
        System.out.println("resf2 = " + resf2);
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

    public static int staticMethod(){
        int field1 = 1;
        int field2 = 2;
        return field1 + field2;
    }

    public static int staticMethod(int field1,int field2){
        return field1 + field2;
    }

    public native void MyNative();

}
