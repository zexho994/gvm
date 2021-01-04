public class GaussTest {

    public int gauss(int initNum, int loopSize) {
        int start = initNum;
        int sum = 0;
        for (int i = 1; i <= loopSize; i++) {
            start += i;
        }
        System.out.println(sum);
        return sum;
    }
}
