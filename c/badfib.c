#include <stdlib.h>
#include <time.h>
#include <stdio.h>


long long int fib(int n) {
  if(n<=2) return 1;
  return fib(n-1)+fib(n-2);
}

int main(int argc, char const *argv[]) {

  int num = atoi(argv[1]);
  long long int res;
  clock_t begin = clock();
  res = fib(num, 4);
  clock_t end = clock();

  printf("num: %lld,\ttime: %f\n", res, (double) (end-begin)/CLOCKS_PER_SEC);

  return 0;
}
