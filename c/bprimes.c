#include <stdlib.h>
#include <stdio.h>
#include <time.h>

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
  for (int i=2; i*i<max; i++) {
    if(!primers[i])
      for (int j=i*2; j<max; j+=i) primers[j]=1;
  }
}

int main(int argc, char const *argv[]) {
  if (argc!=2) usage(argv[0]);

  max = atoi(argv[1]);
  int *primers;
  primers = malloc(sizeof(int)*max);
  if (primers==NULL) {
    fprintf(stderr,"Error allocating memory\n");
    exit(-1);
  }

  clock_t begin = clock();
  markPrimes(primers);
  clock_t end = clock();

  //escriu(primers);
  printf("time: %f\n", (double) (end-begin)/CLOCKS_PER_SEC);

  return 0;
}
