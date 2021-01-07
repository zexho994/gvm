/**
 * @author : zexho
 * @created : 2021-01-07
**/
public class PrintStaticTest{
  static int i = 994;
  static float f = 994f;
  static double d = 994d;
  static boolean b = true;
  static long l = 994l;
  static String s = "hello world";

  public static void main(String[] args){
    GvmOut.to(i); 
    GvmOut.to(f); 
    GvmOut.to(d); 
    GvmOut.to(b); 
    GvmOut.to(l);
    GvmOut.to(s);
  }
}
