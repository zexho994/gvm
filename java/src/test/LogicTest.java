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
      if (n > 9){
        GvmOut.to(10);
      } else{
        GvmOut.to(9);

      }
  }

  public static void for_test(){
    for(int i = 0 ; i < 5 ; i++){
      GvmOut.to(i);
    }

  }
  
  public static void while_test(){
    int x = 100;
    while(x < 105){
      GvmOut.to(x++);
    }

    x = 100;
    while(x < 105){
      GvmOut.to(++x);
    }
  }

}
