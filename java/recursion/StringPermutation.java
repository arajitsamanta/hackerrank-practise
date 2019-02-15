
public class StringPermutation{
  public static void main(String[] args){
    permute("abc");
  }

  static void permute(String s){
    permuteHelper("",s);
  }

  static void permuteHelper(String soFar,String rest){
    indent(soFar.length());
    System.out.println("permuteHelper(\""+ soFar + "\",\""+rest+"\")");

    if(rest.length()==0){
      System.out.println(soFar);
    }else{
      for(int i=0;i<rest.length();i++){
        String temp=rest.substring(0,i) + rest.substring(i+1);
        permuteHelper(soFar + rest.charAt(i), temp);
      }
    }

  }

  static void indent(int n){
    for(int i=0;i<n;i++){
      System.out.print("    ");
    }
  }

}