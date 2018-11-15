import java.util.Map;
import java.util.HashMap;
import java.util.Stack;

//https://www.hackerrank.com/challenges/balanced-brackets/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
class BalancedBrackets {

  static boolean isBalanced(String s) {

    Map<Character, Character> braces = new HashMap<>() {
      {
        put('(', ')');
        put('[', ']');
        put('{', '}');
        put('<', '>');
      }
    };

    int len = s.length();
    char ch;

    if (len % 2 != 0)
      return false;

    Stack<Character> halfBraces = new Stack<Character>();
    for (int i = 0; i < len; i++) {
      ch = s.charAt(i);
      if (braces.containsKey(ch)) {
        halfBraces.push(ch);
      } else if (halfBraces.isEmpty() || braces.get(halfBraces.pop()) != ch) {
        return false;
      }

    }

    return halfBraces.isEmpty() ? true : false;
  }

  public static void main(String args[]) {

    /*
     * }][}}(}][))]
     * [](){()}
     * ()
     * ({}([][]))[]()
     * {)[](}]}]}))}(())(
     * ([[)
     *
     * NO
     * YES
     * YES
     * YES
     * NO
     * NO
     *
     */

    System.out.println(isBalanced(("}][}}(}][))]")));
    System.out.println(isBalanced(("[](){()}")));
    System.out.println(isBalanced(("()")));
    System.out.println(isBalanced(("({}([][]))[]()")));
    System.out.println(isBalanced(("{)[](}]}]}))}(())(")));
    System.out.println(isBalanced(("([[)")));


  }
}