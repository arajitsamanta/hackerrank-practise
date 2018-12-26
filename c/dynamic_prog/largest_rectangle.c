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

// https://www.hackerrank.com/challenges/largest-rectangle/problem
long largestRectangleBruteForce(int h_count, int* h) {

    int i,j,height,width;
    long max_area=INT_MIN,area=0;
    //3 , 2 ,3
    for(i=0;i<h_count;i++){
        height=h[i];
        for(j=i;j>=0;j--){
            width=j-i+1;
            area=height * width;
            max_area = area > max_area ? area  : max_area; 
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

    long result = largestRectangle(h_count, h);

    fprintf(fptr, "%ld\n", result);

    fclose(fptr);

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
