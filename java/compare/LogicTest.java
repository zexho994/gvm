/**
 * @author : a994
 * @created : 2021-01-05
**/
public class LogicTest{

  public static void main(String[] args){
    LogicTest.if_test();  
    LogicTest.for_test();
    LogicTest.while_test();
  }

  public static void if_test(){
      int n = 10;
      GvmOut.to("10 > 9?");
      if (n > 9){
        GvmOut.to("yes");
      } else{
        GvmOut.to("no");
      }
  }

  public static void for_test(){
    GvmOut.to("for : ");
    for(int i = 0 ; i < 5 ; i++){
      GvmOut.to(i);
    }
  }
  
  public static void while_test(){
    GvmOut.to("while: ");
    int x = 100;
    while(x < 105){
      GvmOut.to(x++);
    }
  }

}
