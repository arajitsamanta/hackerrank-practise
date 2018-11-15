#include <assert.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// https://www.hackerrank.com/challenges/abbr/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dynamic-programming
// Complexity: O(n^2)
char *abbreviation(char *a, char *b)
{
    int a_len = get_length(a), b_len = get_length(b), i, j, temp;
    int row = a_len + 1, col = b_len + 1;
    int dp[row][col];
    for (int i = 0; i < row; i++)
    {
        for (int j = 0; j < col; j++)
        {
            dp[i][j] = 0;
        }
    }
    dp[0][0] = 1;

    for (i = 0; i < a_len; i++)
    {
        for (j = 0; j <= b_len; j++)
        {
            // Delete lowercase char from a
            if (a[i] >= 97 && a[i] <= 122)
            {
                //printf("found lower\n");
                dp[i + 1][j] |= dp[i][j];
            }

            // Match characters, make sure char from a is upper case
            temp = (a[i] >= 65 && a[i] <= 90) ? a[i] : a[i] - 32;
            if (j < b_len && temp == b[j])
            {
                dp[i + 1][j + 1] |= dp[i][j];
            }
        }
    }

    return dp[a_len][b_len] ? "YES" : "NO";
}

int main()
{
    FILE *fptr = fopen("out.txt", "w");
    FILE *infptr = fopen("in.txt", "r");

    char *q_endptr;
    char *q_str = readline(infptr);
    int q = strtol(q_str, &q_endptr, 10);

    if (q_endptr == q_str || *q_endptr != '\0')
    {
        exit(EXIT_FAILURE);
    }
    int q = 2;

    for (int q_itr = 0; q_itr < q; q_itr++)
    {
        char *a = readline();

        char *b = readline();

        char *result = abbreviation(a, b);

        fprintf(fptr, "%s\n", result);
    }

    fclose(fptr);

    return 0;
}

char *readline()
{
    size_t alloc_length = 1024;
    size_t data_length = 0;
    char *data = malloc(alloc_length);

    while (true)
    {
        char *cursor = data + data_length;
        char *line = fgets(cursor, alloc_length - data_length, fptr);

        if (!line)
        {
            break;
        }

        data_length += strlen(cursor);

        if (data_length < alloc_length - 1 || data[data_length - 1] == '\n')
        {
            break;
        }

        size_t new_length = alloc_length << 1;
        data = realloc(data, new_length);

        if (!data)
        {
            break;
        }

        alloc_length = new_length;
    }

    if (data[data_length - 1] == '\n')
    {
        data[data_length - 1] = '\0';
    }

    data = realloc(data, data_length);

    return data;
}
