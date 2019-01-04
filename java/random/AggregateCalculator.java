
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Collections;
import java.util.Comparator;
import java.util.Date;
import java.util.List;
import java.util.Locale;
import java.util.TimeZone;

//Container to hold necessary session info from csv file.
class Session {

  Date startTime;
  String status;
  double duration;

  public Session(Date startTime, String status, double duration) {
    this.startTime = startTime;
    this.status = status;
    this.duration = duration;
  }

  public Date getStartTime() {
    return startTime;
  }

  public String getStatus() {
    return status;
  }

  public double getDuration() {
    return duration;
  }

}

public class AggregateCalculator {

  private static String FILE_PATH = "C:/Arajit/Personal/interview-1-master/session_history.csv";;
  private static String SEPARATOR = ",";
  private static String PASS = "passed";
  static double BUILD_FAILURE_THRESOLD = 20;
  static double AVG_BUILD_DURATION_THRESOLD = 500;

  /**
   * Converts an ISO-8601 formatted UTC timestamp.
   *
   * @return The parsed {@link Date}, or null.
   */
  public static Date convertUTCStringToDate(String isoUtcString) {
    DateFormat isoUtcFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss", Locale.getDefault());
    isoUtcFormat.setTimeZone(TimeZone.getTimeZone("UTC"));
    try {
      return isoUtcFormat.parse(isoUtcString);
    } catch (ParseException e) {
      e.printStackTrace();
      return null;
    }
  }

  /**
   * Format and print to console.
   * 
   * @param week               int current week.
   * @param weeklyCount        int total number of session for the current week.
   * @param weeklyFailureCount int
   * @param weeklyBuildTime    double weekly build time
   */
  public static void print(int week, int weeklyCount, int weeklyFailureCount, double weeklyBuildTime) {
    StringBuilder outBuf = new StringBuilder();

    double failurePercet = (weeklyFailureCount * 100.0d / weeklyCount);
    double weeklyBuildAvgTime = weeklyBuildTime / weeklyCount;

    outBuf.append("Week").append(week).append("-> Failure Count:").append(weeklyFailureCount).append(", Pass:")
        .append((weeklyCount - weeklyFailureCount) * 100.0d / weeklyCount).append("%").append(", Fail:")
        .append(failurePercet).append("%");

    if (failurePercet >= BUILD_FAILURE_THRESOLD) {
      outBuf.append(", ABNORMAL_FAILURE");
    }

    if (weeklyBuildAvgTime >= AVG_BUILD_DURATION_THRESOLD) {
      outBuf.append(", ABNORMAL_TIME");
    }

    System.out.println(outBuf.toString());

  }

  public static void main(String[] args) {
    String line = "";
    int lineCount = -1;
    List<Session> sessionList = new ArrayList<>();
    Calendar cal = Calendar.getInstance();

    try (BufferedReader br = new BufferedReader(new FileReader(FILE_PATH))) {

      while ((line = br.readLine()) != null) {
        lineCount++;

        // Skip header line
        if (lineCount != 0) {
          String[] session = line.split(SEPARATOR);

          Session sl = new Session(convertUTCStringToDate(session[2].substring(1, session[2].length() - 1)),
              session[3].substring(1, session[3].length() - 1),
              Double.valueOf(session[4].substring(1, session[4].length() - 1)));

          sessionList.add(sl);
        }

      }

    } catch (IOException e) {
      e.printStackTrace();
    }

    // Sort the session info in natural order of build start time
    Collections.sort(sessionList, new Comparator<Session>() {
      @Override
      public int compare(Session o1, Session o2) {
        return o1.getStartTime().compareTo(o2.getStartTime());
      }
    });

    // Extract starting point and add 7 days to crate 7 day marker
    // variable(sevenDayEndMarker)
    cal.setTime(sessionList.get(0).getStartTime());
    cal.add(Calendar.DATE, 7);
    Date sevenDayEndMarker = cal.getTime();

    int week = 1, weeklyCount = 0, weeklyFailureCount = 0, noOfRecordsProcessed = 0, totalFailure = 0;
    ;
    double weeklyBuildTime = 0;

    /*
     * Iterate over the sorted collection and find out weekly total count, build
     * time, failure , success etc using sevenDayEndMarker marker
     */
    for (int i = 0; i < lineCount; i++) {

      Session s = sessionList.get(i);

      /*
       * If current session start time is greater than the sevenDayEndMarker, it means
       * all the records for current week has been processed
       */ 
      if (s.getStartTime().compareTo(sevenDayEndMarker) > 0) {

        noOfRecordsProcessed += weeklyCount;
        totalFailure += weeklyFailureCount;

        // Print required info to console
        print(week, weeklyCount, weeklyFailureCount, weeklyBuildTime);

        // Re-evaluate sevenDayEndMarker for next iteration
        cal.setTime(s.getStartTime());
        cal.add(Calendar.DATE, 7);
        sevenDayEndMarker = cal.getTime();

        // Reset counters for next iteration
        weeklyCount = 1;
        weeklyFailureCount = !s.getStatus().equalsIgnoreCase(PASS) ? 1 : 0;
        week++;
        weeklyBuildTime = 0;
      } else {
        /*
         * When current start time is less than sevenDayEndMarker, calculate
         * weeklyCount,weeklyBuildTime, weeklyBuildTime, weeklyFailureCount etc
         */
        weeklyCount++;
        weeklyBuildTime += s.getDuration();
        if (!s.getStatus().equalsIgnoreCase(PASS)) {
          // Anything other than "passed" has been considered as failure.
          weeklyFailureCount++;
        }
      }

    }

    /*
     * If total no of records processed is less than lineCount, print
     * (lineCount-noOfRecordsProcessed)
     */
    if (noOfRecordsProcessed < lineCount) {
      totalFailure += weeklyFailureCount;
      print(week, weeklyCount, weeklyFailureCount, weeklyBuildTime);
    }

    System.out.println("Total Pass/Fail-> " + ((lineCount - totalFailure) * 100.0d) / lineCount + "%,"
        + (totalFailure * 100.0d) / lineCount + "%");

  }

}
