#include <stdio.h>
//#include <string.h>
#include "libsoe.h"

int main(){
    char p[100];
    int num[10]={1,2,3,4,5,6,7,8,9,10};

    scanf("%s",&p);
//    for(int i=0;i<10;i++){
//       scanf("%d",&num[i]);
//    }
    int a = handler(1,2);
    printf("%d\n",a);
    printChar(p);
    numbers(&num,10);
    return 0;

}