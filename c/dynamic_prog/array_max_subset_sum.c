#include <stdio.h>

int max(int a, int b);
int maxSubsetSum(int n, int *arr);

int max(int a, int b)
{
  return a > b ? a : b;
}

//https://www.hackerrank.com/challenges/max-array-sum/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dynamic-programming
//Complexity: O(n)
int maxSubsetSum(int n, int *arr)
{

  if (n == 0)
    return 0;
  arr[0]=max(0,arr[0]);
  if (n == 1)
    return arr[0];
  arr[1]=max(arr[0],arr[1]);

  for (int i = 2; i < n; i++)
  {
    arr[i] = max(arr[i-1],arr[i]+arr[i-2]);
  }

  return arr[n-1];
}

int main()
{
  int arr[] = {-3, 5, 2, 0, 4, -6};
  //int arr[] = {1,2,3,4,5,6,7,8,9,10};
  //int arr[] = {3, 7, 4, 6, 5};
  //int arr[]={3, 5, -7, 8, 10};

  int n = sizeof(arr) / sizeof(arr[0]);

  printf("\nFInal Max sum = %d", maxSubsetSum(n, arr));
  return 0;
}