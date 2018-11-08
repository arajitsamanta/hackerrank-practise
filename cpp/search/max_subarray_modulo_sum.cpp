//https://www.hackerrank.com/challenges/maximum-subarray-sum/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
// C++ program to find sub-array having maximum
// sum of elements modulo m.
#include <bits/stdc++.h>
using namespace std;

long maximumSum(vector<long> a, long m)
{

  long prefix = 0, maximum = 0, n = a.size();
  set<long> S;

  S.insert(0);
  for (long i : a)
  {
    prefix = (prefix + i) % m;
    maximum = max(prefix, maximum);

    auto it = S.lower_bound(prefix + 1);
    if (it != S.end())
    {
      maximum = max(maximum, prefix - (*it) + m);
    }
    S.insert(prefix);
  }
  return maximum;
}

int main(int argc, char const *argv[])
{
  int arr[] = {3, 3, 9, 9, 5};
  int n = sizeof(arr) / sizeof(arr[0]);
  int m = 7;
  cout << maxSubarray(arr, n, m) << endl;

  vector<long> a(5);
  for (int i = 0; i < n; i++)
  {
    //long a_item =arr[i];
    a[i] = arr[i];
  }

  cout << maximumSum(a, m) << endl;
  return 0;
}
