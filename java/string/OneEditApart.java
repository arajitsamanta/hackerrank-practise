
public class OneEditApart {

    static boolean isOneEditApart(String s1, String s2) {

        //Start alway's with shorter string
        if (s1.length() > s2.length()) {
            String temp = s1;
            s1 = s2;
            s2 = temp;
        }

        //If length of longer string - shorter it can never be one edit awyay
        if (s2.length() - s1.length()>1) {
            return false;
        }

        boolean foundDiff = false;

        for (int idx1 = 0, idx2 = 0; idx1 < s1.length() && idx2<s2.length();) {

            if (s1.charAt(idx1) != s2.charAt(idx2)) {
                //return false if is second difference
                if (foundDiff)
                    return false;
                foundDiff=true;
                
				if (s1.length() == s2.length()) { //If lengths are equal, move shorter
					idx1++;
				}
            } else {
               idx1++; // If matching, move shorter pointer
            }
            idx2++; //Always increase longer string index

        }

        return true;

    }

    public static void main(String[] args) {
        System.out.println(isOneEditApart("cat", "dog"));
        System.out.println(isOneEditApart("cats", "cat"));
        System.out.println(isOneEditApart("cat", "cut"));
        System.out.println(isOneEditApart("cat", "cast"));
        System.out.println(isOneEditApart("cat", "at"));
        System.out.println(isOneEditApart("cat", "act"));
    }
}
