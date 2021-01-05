public class Add {
    public void add(int n1, int n2) {
        int n3 = n1 + n2;
    }

    public void add1(){
        int i = 0;
        int j = i++;
    }

    public void add2(){
        int i = 0;
        int j = ++i;
    }

    public void add3(){
        int i = 0;
        i = i + 1;
        i++;
        ++i;
    }
}
