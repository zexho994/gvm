public class Ldc{
	public static final int static_1 = 1;
	public static final String static_str1 = "static str1";

	public static void main(String[] args){
		int int1 = 1;
		String str1 = "static str1";
		String str2 = "static str2";
	}	
	class ldc {
		int k;
		public void fun(){
			k = Ldc.static_1;	
		}
	}	
}
