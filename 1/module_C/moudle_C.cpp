#include <iostream>

extern "C"{
    #include "demo.h"
}

void SayHello(const char* s){
    std :: cout << s;
}