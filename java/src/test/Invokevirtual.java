/**
 * @author Zexho
 * @date 2021/8/21 11:26 上午
 */
public class Invokevirtual {

    public static void main(String[] args) {
        Invokevirtual invokevirtual = new Invokevirtual();
        int res = invokevirtual.additive(1, 2);
        GvmOut.to(res);
    }

    public int additive(int x, int y) {
        int r = x + y;
        return r;
    }

}
