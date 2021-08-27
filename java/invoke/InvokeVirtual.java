public class InvokeVirtual {

    public static void main(String[] args) {
        InvokeVirtual invokeVirtual = new InvokeVirtual();
        int res = invokeVirtual.additive(1, 2);
        GvmOut.to(res);
    }

    public int additive(int x, int y) {
        int r = x + y;
        return r;
    }

}