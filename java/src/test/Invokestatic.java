/**
 * @author Zexho
 * @date 2021/8/24 6:31 下午
 */
public class Invokestatic {
    public static void main (String[] args){
        Invokestatic.foo();
        Invokestatic invoke = new Invokestatic();
        invoke.foo2();
    }

    public static void foo(){

    }

    public final void foo2(){}
}
