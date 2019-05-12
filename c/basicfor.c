#include <omp.h>
#include <stdio.h>
#include<time.h>
#define N 100000
 #define CHUNKSIZE 100

int main(int argc, char const *argv[]) {
  clock_t t1=clock();

  int i, chunk;
  float a[N], b[N], c[N];

  /* Some initializations */
  for (i=0; i < N; i++)
  a[i] = b[i] = i * 1.0;
  chunk = CHUNKSIZE;

  #pragma omp parallel shared(a,b,c,chunk) private(i)
  {

  #pragma omp for schedule(auto) nowait
  for (i=0; i < N; i++)
  c[i] = a[i] + b[i];

  }   /* end of parallel region */

  clock_t t2=clock();
  printf("The time taken is.. %g \n", (double) (t2-t1));
  return 0;
}
