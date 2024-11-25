#include <stdio.h>

int *g;

void first()
{
  int a = 1;
  g = &a;
}

void second()
{
  int b = 2;
}

int main(int argc, char **argv)
{
  first();
  second();
  printf("%d\n", *g);
  return 0;
}
