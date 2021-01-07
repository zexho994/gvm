/**
 * @author : zexho
 * @created : 2021-01-05
**/
public class AlgorithmTest {
  
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

    AlgorithmTest.add();
    AlgorithmTest.div();
    AlgorithmTest.mul();
    AlgorithmTest.sub();
  }

  public static void add(){
    int i1 = 10;
    int i2 = 20;
    GvmOut.to(i1 + i2);

  }

  public static void div(){
    int i1 = 99;
    int i2 = 33;
    GvmOut.to(i1 / i2);
  }

  public static void mul(){
    int i1 = 100;
    int i2 = 4;
    GvmOut.to(i1 * i2);
  }

  public static void sub(){
    int i1 = 4;
    int i2 = 104;
    GvmOut.to(i1 - i2);
  }

}
