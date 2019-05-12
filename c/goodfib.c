#include <stdlib.h>
#include <time.h>
#include <stdio.h>
#include <omp.h>

int max, num_th;

long long int fib(int n) {
  if(n<=2) return 1;

    long long int r1, r2;

    #pragma omp task shared(r1) if(n>max-2*num_th)
    {
      r1=fib(n-1);
    }
    #pragma omp task shared(r2) if(n>max-2*num_th)
    {
      r2=fib(n-2);
    }
    #pragma omp taskwait
    return r1+r2;
}

int main(int argc, char const *argv[]) {

  int num = atoi(argv[1]);
  max=num;
  num_th = omp_get_num_threads();
  long long int res;

  clock_t begin = clock();
  #pragma omp parallel firstprivate(num)
  {
    #pragma omp single
    {
      res = fib(num);
    }
  }
  clock_t end = clock();

  printf("num: %lld,\ttime: %f\n", res, (double) (end-begin)/CLOCKS_PER_SEC);

  return 0;
}
