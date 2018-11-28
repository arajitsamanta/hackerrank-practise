#include <assert.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *readline();
char **split_string(char *);

//https://www.hackerrank.com/challenges/game-of-two-stacks/problem
int twoStacks(int x, int a_count, int *a, int b_count, int *b)
{
    int sum = 0, idxA = 0, idxB = 0, count = 0;

    //considering only first stack and calculating count
    while (idxA < a_count  && sum+a[idxA]<=x )
    {
        sum += a[idxA];
        idxA++;
    }
    count = idxA;

    while (idxB < b_count && idxA >= 0)
    { 
        //now adding one element of second stack at a time and subtracting the top element of first stack and calculating the count
        sum += b[idxB];
        idxB++;
        while (sum > x && idxA > 0)
        {
            idxA--;
            sum -= a[idxA];
        }
        if (sum <= x && idxA + idxB > count)
            count = idxA + idxB;
    }
    return count;
}

int main()
{
    //FILE* fptr = fopen(getenv("OUTPUT_PATH"), "w");

    char *g_endptr;
    char *g_str = readline();
    int g = strtol(g_str, &g_endptr, 10);

    if (g_endptr == g_str || *g_endptr != '\0')
    {
        exit(EXIT_FAILURE);
    }

    for (int g_itr = 0; g_itr < g; g_itr++)
    {
        char **nmx = split_string(readline());

        char *n_endptr;
        char *n_str = nmx[0];
        int n = strtol(n_str, &n_endptr, 10);

        if (n_endptr == n_str || *n_endptr != '\0')
        {
            exit(EXIT_FAILURE);
        }

        char *m_endptr;
        char *m_str = nmx[1];
        int m = strtol(m_str, &m_endptr, 10);

        if (m_endptr == m_str || *m_endptr != '\0')
        {
            exit(EXIT_FAILURE);
        }

        char *x_endptr;
        char *x_str = nmx[2];
        int x = strtol(x_str, &x_endptr, 10);

        if (x_endptr == x_str || *x_endptr != '\0')
        {
            exit(EXIT_FAILURE);
        }

        char **a_temp = split_string(readline());

        int a[n];

        for (int a_itr = 0; a_itr < n; a_itr++)
        {
            char *a_item_endptr;
            char *a_item_str = a_temp[a_itr];
            int a_item = strtol(a_item_str, &a_item_endptr, 10);

            if (a_item_endptr == a_item_str || *a_item_endptr != '\0')
            {
                exit(EXIT_FAILURE);
            }

            a[a_itr] = a_item;
        }

        char **b_temp = split_string(readline());

        int b[m];

        for (int b_itr = 0; b_itr < m; b_itr++)
        {
            char *b_item_endptr;
            char *b_item_str = b_temp[b_itr];
            int b_item = strtol(b_item_str, &b_item_endptr, 10);

            if (b_item_endptr == b_item_str || *b_item_endptr != '\0')
            {
                exit(EXIT_FAILURE);
            }

            b[b_itr] = b_item;
        }

        int result = twoStacks(x, n, a, m, b);

        //fprintf(fptr, "%d\n", result);
        printf("result %d\n", result);
    }

    // fclose(fptr);

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
        char *line = fgets(cursor, alloc_length - data_length, stdin);

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

char **split_string(char *str)
{
    char **splits = NULL;
    char *token = strtok(str, " ");

    int spaces = 0;

    while (token)
    {
        splits = realloc(splits, sizeof(char *) * ++spaces);
        if (!splits)
        {
            return splits;
        }

        splits[spaces - 1] = token;

        token = strtok(NULL, " ");
    }

    return splits;
}
