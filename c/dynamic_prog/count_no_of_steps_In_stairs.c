#include <stdio.h>
#include <stdlib.h>

//https://www.hackerrank.com/challenges/ctci-recursive-staircase/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=recursion-backtracking
// Complexity: O(n)
int stepPermsMemo(int n,int memo[]){
    printf("n: %d memo: %d\t",n,memo[n]);
    if(n<0){
        return 0;
    }else if(n==0){
        return 1;
    }else if(memo[n]>-1){
        return memo[n];
    }else{
        memo[n]=stepPermsMemo(n-1,memo)+stepPermsMemo(n-2,memo)+stepPermsMemo(n-3,memo);
        return memo[n];
    }
}

// Complete the stepPerms function below.
int stepPerms(int n) {
    int memo[n+1];
    for(int i=0;i<=n;i++){
        memo[i]=-1;
    }
    return stepPermsMemo(n,memo);
}

int main()
{
  printf("%d\n",stepPerms(3));
}