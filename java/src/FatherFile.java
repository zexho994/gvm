
public class FatherFile {

    public static final boolean FATHER_true = true;
    public byte FATHER_123 = 1;
    public static char FATHER_x = 'X';

    public FatherFile(){
        FATHER_123 = 2;
    }

    protected Object father_method1(){
        return null;
    }

    public int father_method2(int i){
        return i + 1;
    }

}
