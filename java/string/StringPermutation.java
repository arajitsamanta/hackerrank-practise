
public class StringPermutation {

    public static void main(String[] args) {
        permute("","abc");
    }

    static void indent(int n){
        for(int i=0;i<n;i++){
            System.out.print(" ");
        }
    }

    static void permute(String soFar, String rest) {
        indent(soFar.length());
        System.out.println("permute("+soFar+","+rest+")");

        if (rest.isEmpty()) {
            indent(soFar.length());
            System.out.println(soFar);
        } else {
            for (int i = 0; i < rest.length(); i++) {
                String remaining = rest.substring(0, i) + rest.substring(i + 1);
                permute(soFar + rest.charAt(i), remaining);
            }
        }
    }

}