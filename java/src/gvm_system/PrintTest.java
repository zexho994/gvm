public class PrintTest{

  static class Gvm {
    public static void out(int i){};
  }

  public static void main(String []args){
    Gvm.out(1);
  }

}
