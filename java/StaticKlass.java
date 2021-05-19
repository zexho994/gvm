/**
 * @author Zexho
 * @date 2021/5/19 2:36 下午
 */
public class StaticKlass {
    public static PrintOut sk;

    public StaticKlass(){ }

    public StaticKlass(PrintOut s){ sk = s; }

    public void print(String s) {
        System.out.println(s);
    }

    public static void main(String[] args) {
        sk.print("111");
    }

}
