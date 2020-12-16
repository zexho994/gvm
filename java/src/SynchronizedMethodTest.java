/**
 * @author Zexho
 */
public class SynchronizedMethodTest {
    private Object lock = new Object();
    private int i = 0;

    public static void main(String[] args) {
        SynchronizedMethodTest syn = new SynchronizedMethodTest();
        syn.lock();
    }

    public int lock(){
        synchronized(lock){
            i++;
        }
        return i;
    }
}