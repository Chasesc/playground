#include <stdio.h>

int main() {
    int x = 0;
    int *p = &x;
    for(int i = 0; i < 4096; ++i) {
        printf("%d: p=%p, %d\n", i + 1, p++, *p);
    }
}
