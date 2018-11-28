
# https://www.hackerrank.com/challenges/game-of-two-stacks/problem
def twoStacks(x, a, b):
    i = 0
    j = 0
    n = len(a)
    m = len(b)
    sum = 0
    count = 0

    while i < n and (sum+a[i]) <= x:
        sum += a[i]
        i = i+1

    count = i

    while j < m and i >= 0:
        sum += b[j]
        j = j+1

        while sum > x and i > 0:
          sum -= a[i]
          i = i-1

        if sum <= x and i + j > count:
          count = i + j

    return count


x = 10
a = [4, 2, 4, 6, 1]
b = [2, 1, 5, 8]
print("result", twoStacks(x, a, b))
