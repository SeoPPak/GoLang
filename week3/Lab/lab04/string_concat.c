#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* string_concat(const char* str1, const char* str2){
    int len1 = strlen(str1);
    int len2 = strlen(str2);

    char* result = (char*)malloc(len1 + len2 + 1);

    strcpy(result, str1);
    strcpy(result+len1, str2);

    return result;
}