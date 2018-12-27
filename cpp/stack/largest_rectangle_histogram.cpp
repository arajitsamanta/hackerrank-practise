#include <iostream>
#include <stack>
#include <bits/stdc++.h>

using namespace std;

// https://www.hackerrank.com/challenges/largest-rectangle/problems
// Complexit: O(n)
int largestRectangle(vector<int> hist)
{
    // Create an empty stack. The stack holds indexes of hist[] array. The bars stored in stack are always in increasing order of their heights.
    stack<int> s;

    int max_area = 0;  // Initalize max area
    int top;           // To store top of stack
    int area_with_top; // To store area with top bar as the smallest bar

    // Run through all bars of given histogram
    int i = 0,n=hist.size();
    while (i < n)
    {
        // If this bar is higher than the bar on top stack, push it to stack
        if (s.empty() || hist[s.top()] <= hist[i])
        {
            s.push(i++);
        }
        else
        {
            // If this bar is lower than top of stack, then calculate area of rectangle with stack top as the smallest (or minimum height) bar.
            // 'i' is 'right index' for the top and element before top in stack is 'left index'
            top = s.top(); // store the top index
            s.pop();       // pop the top

            // Calculate the area with hist[tp] stack as smallest bar
            area_with_top = hist[top] * (s.empty() ? i : i - s.top() - 1);

            // update max area, if needed
            if (max_area < area_with_top)
                max_area = area_with_top;
        }
    }

    // Now pop the remaining bars from stack and calculate
    // area with every popped bar as the smallest bar
    while (s.empty() == false)
    {
        top = s.top();
        s.pop();
        area_with_top = hist[top] * (s.empty() ? i : i - s.top() - 1);

        if (max_area < area_with_top)
            max_area = area_with_top;
    }

    return max_area;
}

int main()
{
    int hist[] = {6, 2, 5, 4, 5, 1, 6};
    vector<int> vec (hist, hist + sizeof(hist) / sizeof(hist[0]) );
    //int n = sizeof(hist) / sizeof(hist[0]);
    cout << "Maximum area is " << largestRectangle(vec);
    return 0;
}