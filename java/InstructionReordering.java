package java;

/**
 * @author : zexho
 * @created : 2021-08-27
**/
public class InstructionReordering{
  public static void main(String[] args){
    int a = 1;
    int b = 2;
    int c = a + 1;
    int d = 1 + b;
    sys.GvmOut.to(d);
  } 
}
