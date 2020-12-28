/**
 * @author : zexho
 * @created : 2020-12-28
**/
public class CreateArray{

      public static void main(String []args){
         CreateArray c = new CreateArray();
         c.createIntArray();
      }

      public void createIntArray(){
        int[] intarr = new int[18];
        for (int i = 0 ; i < intarr.length ; i++){
          intarr[i] = i + 10;
        }
      }

}
