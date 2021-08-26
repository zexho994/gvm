/**
 * @author Zexho
 * @date 2021/5/18 8:04 下午
 */
public class PrintOut {

    public static void main(String[] args) {
        int i = 1;
        GvmOut.to(i);

        double d = 2d;
        GvmOut.to(d);

        float f = 3f;
        GvmOut.to(f);

        String str = "4";
        GvmOut.to(str);

        boolean z = true;
        GvmOut.to(z);

        long l = 5l;
        GvmOut.to(l);
    }

}
