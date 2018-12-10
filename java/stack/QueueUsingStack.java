import java.util.Scanner;
import java.util.Stack;

//https://www.hackerrank.com/challenges/queue-using-two-stacks/problem
//Time complexity O(q)
public class QueueUsingStack {
  static Scanner sc = new Scanner(System.in);

  public static void main(String[] args) {
    
    Stack<Integer> stack_insert = new Stack<>();
    Stack<Integer> stack_delete = new Stack<>();
    int q = sc.nextInt();
    while (q-- > 0) {
      int type = sc.nextInt();
      if (type == 1) {
        int x = sc.nextInt();
        stack_insert.push(x);
      } else if (type == 2) {
        if (stack_delete.isEmpty()) {
          while (!stack_insert.isEmpty()) {
            stack_delete.push(stack_insert.pop());
          }
        }
        stack_delete.pop();
      } else {
        if (stack_delete.isEmpty()) {
          while (!stack_insert.isEmpty()) {
            stack_delete.push(stack_insert.pop());
          }
        }
        System.out.println(stack_delete.peek());
      }
    }
  }

}
