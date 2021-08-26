public class Foreach{
    public static void main(String[] args) {
        int[] ints = new int[]{1};
        for(int i : ints){
            GvmOut.to(i);
        }
    }
}