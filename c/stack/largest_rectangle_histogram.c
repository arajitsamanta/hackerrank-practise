#include <assert.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* readline();
char** split_string(char*);

int min(int a,int b){
    return a < b ? a : b;
}

long max(long a,long b){
    return a > b ? a : b;
}

// https://www.hackerrank.com/challenges/largest-rectangle/problem
// Complexity: O(n^2)
long largestRectangle_BruteForce(int h_count, int *h)
{

    int i, j, height = 0, width;
    long max_area = INT_MIN;

    for (i = 0; i < h_count; i++)
    {
        //Get current height
        height = h[i];
        for (j = i-1; j >= 0; j--)
        {   
            width = (i - j + 1);  

            //Height should be the minimum between i and j           
            height = min(h[j],height);

            //Calculate area
            max_area = max(max_area,height * width);
        }
    }
    return max_area;
}

int main()
{
    FILE* fptr = fopen(getenv("OUTPUT_PATH"), "w");

    char* n_endptr;
    char* n_str = readline();
    int n = strtol(n_str, &n_endptr, 10);

    if (n_endptr == n_str || *n_endptr != '\0') { exit(EXIT_FAILURE); }

    char** h_temp = split_string(readline());

    int* h = malloc(n * sizeof(int));

    for (int i = 0; i < n; i++) {
        char* h_item_endptr;
        char* h_item_str = *(h_temp + i);
        int h_item = strtol(h_item_str, &h_item_endptr, 10);

        if (h_item_endptr == h_item_str || *h_item_endptr != '\0') { exit(EXIT_FAILURE); }

        *(h + i) = h_item;
    }

    int h_count = n;

    long result = largestRectangle_BruteForce(h_count, h);

    fprintf(stdout, "%ld\n", result);

    //fclose(fptr);

    return 0;
}

char* readline() {
    size_t alloc_length = 1024;
    size_t data_length = 0;
    char* data = malloc(alloc_length);

    while (true) {
        char* cursor = data + data_length;
        char* line = fgets(cursor, alloc_length - data_length, stdin);

        if (!line) { break; }

        data_length += strlen(cursor);

        if (data_length < alloc_length - 1 || data[data_length - 1] == '\n') { break; }

        size_t new_length = alloc_length << 1;
        data = realloc(data, new_length);

        if (!data) { break; }

        alloc_length = new_length;
    }

    if (data[data_length - 1] == '\n') {
        data[data_length - 1] = '\0';
    }

    data = realloc(data, data_length);

    return data;
}

char** split_string(char* str) {
    char** splits = NULL;
    char* token = strtok(str, " ");

    int spaces = 0;

    while (token) {
        splits = realloc(splits, sizeof(char*) * ++spaces);
        if (!splits) {
            return splits;
        }

        splits[spaces - 1] = token;

        token = strtok(NULL, " ");
    }

    return splits;
}
