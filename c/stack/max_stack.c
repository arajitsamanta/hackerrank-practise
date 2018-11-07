#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MIN_INT -1 << (8 * sizeof(int)) - 1
#define SIZE 100

//https://www.hackerrank.com/challenges/maximum-element/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
int main(int argc, char const *argv[])
{
  int n, type, val, top = -1, max = MIN_INT, value, no_of_vals_read, pop_elem;
  scanf("%d\n", &n);
  int stack[n];
  char line[SIZE], *linePtr;
  char delims[] = " \t\r\n";

  while (n--)
  {

    if (fgets(line, SIZE, stdin) == NULL)
    {
      fprintf(stderr, "No input\n");
      return 1;
    }

    linePtr = strtok(line, delims);
    no_of_vals_read = sscanf(linePtr, "%d", &type);

    while (no_of_vals_read > 0)
    {
      linePtr = strtok(NULL, delims);
      no_of_vals_read = (linePtr == NULL) ? 0 : sscanf(linePtr, "%d", &val);
    }

    switch (type)
    {
    case 1:
      if (top == -1)
      {
        stack[++top] = val;
        max = val;
      }
      else
      {
        if (val > max)
        {
          stack[++top] = 2 * val - max;
          max = val;
        }
        else
        {
          stack[++top] = val;
        }
      }
      break;
    case 2:
      pop_elem = stack[top];
      if (pop_elem > max)
      {
        max = 2 * max - pop_elem;
      }
      --top;
      break;
    case 3:
      printf("%d\n", max);
      break;
    default:
      printf("Invalid operation");
    }
  }

  return 0;
}