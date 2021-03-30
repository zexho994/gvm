/**
 * @author Zexho
 * @created 2020/9/18 7:34 下午
 */
public class FibonacciTest {
    public static void main(String[] args) {
        long x = fibonacci(1);
    }

    private static long fibonacci(long n) {
        if (n <= 1) {
            return n;
        }

        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}
