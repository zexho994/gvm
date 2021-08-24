import test.GvmOut;

/**
 * @author : zexho
 * @created : 2021-01-07
**/
public class PrintFieldsTest{
    public int i = 994;
    public float f = 994f;
    public double d = 994d;
    public boolean b = true;
    public long l = 994l;
    public String s = "hello world";
  
    public static void main(String[] args){
        PrintFieldsTest c = new PrintFieldsTest();
        GvmOut.to(c.i);
        GvmOut.to(c.f);
        GvmOut.to(c.d);
        GvmOut.to(c.b);
        GvmOut.to(c.l);
        GvmOut.to(c.s);
    }

}
