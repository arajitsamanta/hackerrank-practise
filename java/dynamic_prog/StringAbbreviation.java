
import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

// https://www.hackerrank.com/challenges/abbr/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dynamic-programming
// Complexity: O(n^2)
public class StringAbbreviation {

  private static final Scanner scanner = new Scanner(System.in);

  static String abbreviation2(String a, String b) {
    char[] x = a.toCharArray();
    char[] y = b.toCharArray();
    boolean[][] dp = new boolean[x.length + 1][y.length + 1];

    // 0 consumed from a, 0 consumed from b is reachable position
    dp[0][0] = true;

    for (int i = 0; i < x.length; i++) {
      for (int j = 0; j <= y.length; j++) {
        // Delete lowercase char from a
        if (Character.isLowerCase(x[i])) {
          dp[i + 1][j] |= dp[i][j];
        }

        // Match characters, make sure char from a is upper case
        if (j < y.length && Character.toUpperCase(x[i]) == y[j]) {
          dp[i + 1][j + 1] |= dp[i][j];
        }
      }
    }

    return dp[x.length][y.length] ? "YES" : "NO";
  }

  public static void main(String[] args) throws IOException {
    File file = new File("out.txt");
    BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(file));

    int q = scanner.nextInt();
    scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

    for (int qItr = 0; qItr < q; qItr++) {
      String a = scanner.nextLine();

      String b = scanner.nextLine();

      String result = abbreviation2(a, b);
      System.out.println(result);

      System.out.println(1 | 0);

      bufferedWriter.write(result);
      bufferedWriter.newLine();
    }

    bufferedWriter.close();

    scanner.close();
  }
}
