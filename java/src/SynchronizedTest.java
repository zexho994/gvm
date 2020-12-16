/**
 * @author Zexho
 */
public class SynchronizedTest {
    public static Object lock = new Object();
    public static void main(String[] args) {
        int i = 0;
        synchronized(lock){
            i++;
        }
    }
}