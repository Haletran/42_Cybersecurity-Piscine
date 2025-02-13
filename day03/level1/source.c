#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *ft_strcpy(char *dest, char *src)
{
    int s;

    s = 0;
    while (src[s] != '\0')
    {
        dest[s] = src[s];
        s++;
    }
    dest[s] = '\0';
    return (dest);
}

int main()
{
    char test[1024];
    char *value;

    value = malloc(strlen("__stack_check") + 1);
    ft_strcpy(value, "__stack_check");
    printf("Please enter a key: ");
    scanf("%s", test);
    if (strcmp(value, test) != 0)
    {
        printf("Nope\n");
    }
    else
    {
        printf("Good job. \n");
    }

    return (0);
}
