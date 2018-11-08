import java.util.TreeSet;
import java.util.SortedSet;

//https://www.hackerrank.com/challenges/maximum-subarray-sum/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
public class MaximumSubarrayModuloSum {

  static long max(long a, long b) {
    return a > b ? a : b;
  }

  static long maximumSum(long[] arr, long m) {

    long prefix = 0, maximum = 0, n = arr.length;

    TreeSet<Long> ts = new TreeSet<>();

    ts.add(0L);

    for (int i = 0; i < n; i++) {

      //Find prefix at index i
      prefix = (prefix + arr[i]) % m;

      //Find max between prefix and max
      maximum = max(maximum, prefix);

      //Find starting point from the tree set where value >= prefix+1
      SortedSet<Long> tail = ts.tailSet(prefix + 1);

      //If tail is not empty re calculate max.
      if (!tail.isEmpty()) {
        maximum = max(maximum, prefix - tail.first() + m);
      }

      //Insert prefix to tree set.
      ts.add(prefix);
    }

    return maximum;

  }

  public static void main(String[] args) {

    long[] arr = { 3, 3, 9, 9, 5 };
    System.out.println(maximumSum(arr, 7));
  }
}