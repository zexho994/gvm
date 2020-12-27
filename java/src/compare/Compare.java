public class Compare{
	public static void main(String[] args) {
		Compare compare = new Compare();
		compare.if_inst();
		compare.for_inst();
		compare.while_inst();
	}


	public void if_inst(){
		int sun = 0;
		int i = 5;
		if ( i > 10){
			sun = 100;
		}else{
			sun = 10;
		}
	}

	public void for_inst(){
		int sun = 0;
		for(int i = 0;i < 10;i++){
			++sun;
		}
	}

	public void while_inst(){
		int sun = 0;
		while(sun < 10){
			++sun;
		}
	}

}