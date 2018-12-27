
import java.util.Stack;
import java.util.Scanner;

class LargestRectanglularAreaInHistogram {

    private static final Scanner scanner = new Scanner(System.in);

    // https://www.hackerrank.com/challenges/largest-rectangle/problems
    // Complexit: O(n)
    static int largestRectangle(int[] hist) {
        final Stack<Integer> s = new Stack<>();

        int maxArea = 0;
        int tp;
        int areaWithTop;

        int i = 0;
        while (i < hist.length) {
            if (s.empty() || hist[s.peek()] <= hist[i]) {
                s.push(i++);
            } else {
                tp = s.pop();
                int w = s.empty() ? i : i - s.peek() - 1;
                areaWithTop = hist[tp] * w;

                if (maxArea < areaWithTop)
                    maxArea = areaWithTop;
            }
        }

        while (!s.empty()) {
            tp = s.pop();
            int w = s.empty() ? i : i - s.peek() - 1;
            areaWithTop = hist[tp] * w;

            if (maxArea < areaWithTop)
                maxArea = areaWithTop;
        }

        return maxArea;
    }

    public static void main(String[] args) {
        int n = Integer.parseInt(scanner.nextLine().trim());
        int[] hist = new int[n];
        for (int i = 0; i < n; i++) {
            hist[i] = scanner.nextInt();
        }
        System.out.println(largestRectangle(hist));
    }

}
