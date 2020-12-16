/**
 * @author Zexho
 * @created 2020/9/17 8:45 上午
 */
public class InvokeDemo implements Runnable {
    public static void main(String[] args) {
        new InvokeDemo().test();
    }

    public void test() {
        InvokeDemo.staticMethod();
        InvokeDemo demo = new InvokeDemo();
        demo.instanceMethod();
        super.equals(null);
        this.run();
        ((Runnable) demo).run();
    }

    public static void staticMethod() {
    }

    private void instanceMethod() {
    }

    @Override
    public void run() {
    }

}
