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
    // Finding prefix sum.
    prefix = (prefix + i) % m;

    // Finding maximum of prefix sum.
    maximum = max(prefix, maximum);

    cout << prefix << " " << maximum << endl;

    // Finding iterator pointing to the first
    // element that is not less than value
    // "prefix + 1", i.e., greater than or
    // equal to this value.
    auto it = S.lower_bound(prefix + 1);

    cout << "lower bound:" << (*it) << endl;

    if (it != S.end())
    {
      maximum = max(maximum, prefix - (*it) + m);
    }

    // // Inserting prefix in the set.
    S.insert(prefix);
  }
  return maximum;
}

int main(int argc, char const *argv[])
{
  long arr[] = {3, 3, 9};
  int m = 7;

  vector<long> a(5);
  for (int i = 0; i < 5; i++)
  {
    a[i] = arr[i];
  }

  cout << maximumSum(a, m) << endl;
  return 0;
}
