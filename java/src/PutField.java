public class PutField{
	public int int_1  = 1;
	public float float_1 = 1.0f;
	public boolean flag = true;
	public String str_1 = "str1";
	public double double_1 = Double.MAX_VALUE;
	public long long_1 = Long.MAX_VALUE;
	public char char_1 = 'c';
	public byte byte_1 = 'a';

	public static int static_int_1 = 1;

	public static void main(String []args){
		PutField c = new PutField();
		c.int_1 = 2;
		c.float_1 = 2.0f;
		c.flag = false;
		c.str_1 = "str2";
	}

}
