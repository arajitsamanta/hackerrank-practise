import java.util.Scanner;

//https://www.hackerrank.com/challenges/game-of-two-stacks/problem
public class GameOfTwoStacks {

  private static final Scanner scanner = new Scanner(System.in);

  static int twoStacks(int x, int a[], int b[]) {

    int i = 0, j = 0, n = a.length, m = b.length, count = 0, sum = 0;

    while (i < n && sum + a[i] <= x) {
      sum += a[i];
      i++;
    }

    count = 1;

    while (j < m && i >= 0) {
      sum += b[j];
      j++;

      while (sum > x && i > 0) {
        i--;
        sum -= a[i];
      }

      if (sum <= x && i + j > count)
        count = i + j;
    }

    return count;
  }

  public static void main(String[] args) {

    //BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

    int g = Integer.parseInt(scanner.nextLine().trim());

    for (int gItr = 0; gItr < g; gItr++) {
      String[] nmx = scanner.nextLine().split(" ");

      int n = Integer.parseInt(nmx[0].trim());

      int m = Integer.parseInt(nmx[1].trim());

      int x = Integer.parseInt(nmx[2].trim());

      int[] a = new int[n];

      String[] aItems = scanner.nextLine().split(" ");

      for (int aItr = 0; aItr < n; aItr++) {
        int aItem = Integer.parseInt(aItems[aItr].trim());
        a[aItr] = aItem;
      }

      int[] b = new int[m];

      String[] bItems = scanner.nextLine().split(" ");

      for (int bItr = 0; bItr < m; bItr++) {
        int bItem = Integer.parseInt(bItems[bItr].trim());
        b[bItr] = bItem;
      }

      int result = twoStacks(x, a, b);

      System.out.println(result);

      //bufferedWriter.write(String.valueOf(result));
      //bufferedWriter.newLine();
    }
    //bufferedWriter.close();
  }

}