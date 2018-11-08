
import java.io.FileReader;
import java.io.BufferedReader;
import java.io.IOException;
import java.util.Arrays;

public class PrintLastKLines {

  /**
   * Complexity ========= Time: O(n). where n is total number of lines in the
   * file. Space: O(k). k is number of lines to be printed
   *
   * @param path The file path
   * @param k    last k lines to be printed from the file
   * @throws IOException
   */
  static void tailLastKLines(String path, int k) throws IOException {
    String[] temp = new String[k];
    int i, count = 1;
    try (BufferedReader br = new BufferedReader(new FileReader(path))) {

      for (String line; (line = br.readLine()) != null;) {
        // Temp array holds only those lines that are to be printed. But we don't know
        // when line Nth from last is reached.
        if (count == (k + 1)) {
          count = 1;
        }
        temp[count - 1] = line;
        count++;
      }
    }

    // At the end of the loop varialbe count-1) indicates the index in the array
    // from where you can start printing the lines.
    int startPrintFrom = count - 1;

    // Print start position to k.
    if (count <= k) {
      for (i = startPrintFrom; i < k; i++) {
        System.out.println(temp[i]);
      }
    }

    // Print rest(from index 0 to start position)
    for (i = 0; i < startPrintFrom; i++) {
      System.out.println(temp[i]);
    }

  }

  public static void main(String[] args) throws IOException {
    tailLastKLines("./temp.txt", 5);
  }

}