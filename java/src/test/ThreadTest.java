/**
 * @author Zexho
 * @date 2021/1/8 7:04 下午
 */
public class ThreadTest {

    public static void main(String[] args) {
        Soap<String, String> s = str -> "hello soap";
        s.to();

//         Soap<Integer, Integer> s2 = ThreadTest::soap1;
//         s2.to(1,2,3);
//
//         ThreadTest threadTest = new ThreadTest();
//         Soap<Object, Object> s3 = threadTest::soap2;
//         s3.to(new Object());
    }

//     public static Integer soap1(Integer... ints) {
//         int i1 = 1;
//         int i2 = 2;
//         return i1 + i2 + ints[0];
//     }
//
//     public Object soap2(Object... objs) {
//         return objs[0];
//     }
}
