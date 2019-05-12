/*
  ================================
  gcc -fopenmp -o name name.c -lm
  ================================
*/
#include <stdlib.h>
#include <stdio.h>
#include <time.h>
#include <omp.h>
#include <math.h>

int max;

void usage(const char *name) {
  printf("Usage: %s maxNum\n", name);
  exit(0);
}

void escriu(int *p) {
  for (int i=2; i<max; i++) {
    if(!p[i]) printf("%d ", i);
  }
  printf("\n");
}

void markPrimes(int *primers) {
  int i, j;
  int m = (int) (sqrt(max)+0.5);
    #pragma omp parallel for private(j) schedule(dynamic, 1)
    for (i=2; i<=m; i++) {
      if(!primers[i]) {
        for (j=i*i; j<=max; j+=i) primers[j]=1;
      }
    }
}

int guillem(int *primers) {
  for(int i=5; i<max; i++) {
    if (!primers[i] && (long long unsigned int)(i*i-1)%24!=0) {
      printf("%i\n", i);
      return 0;
    }
  }
  return 1;
}

int main(int argc, char const *argv[]) {
  if (argc!=2) usage(argv[0]);

  max = atoi(argv[1]);
  int *primers;
  primers = calloc(max, sizeof(int));
  if (primers==NULL) {
    fprintf(stderr,"Error allocating memory\n");
    exit(-1);
  }

  double begin = omp_get_wtime();
  markPrimes(primers);
  double end = omp_get_wtime();

  //escriu(primers);
  if (guillem(primers)) printf("El guillem tenia rao!!\n");
  printf("time: %f\n", end-begin);
  free(primers);
  return 0;
}
