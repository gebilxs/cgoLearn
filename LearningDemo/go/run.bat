@echo off
gcc -m32 ./c_demo.c -o c_demo.exe -I. -L. -lsoe
c_demo.exe