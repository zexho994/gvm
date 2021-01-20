/**
 * @author : zexho
 * @created : 2021-01-20
 **/

public class SynchronizedTest{
  static Object lock = new Object();

  public static void main(String[] args){
    synchronized (lock){
      GvmOut.to(1);
    }
  }

  public synchronized void lockMethod(){
    GvmOut.to("lock method!");
  }
}
