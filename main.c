#include <stdio.h>
#include <ctype.h>

int isPalindrome(char *s)
{
    int sizeS = 0;
    int sizeTmp = 0;
    for (int i = 0;; i++)
    {
        if (s[i] == '\0')
        {
            break;
        }
        if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9'))
        {
            sizeTmp++;
        }
        sizeS++;
    }
    // sizeS++;
    char tmp[sizeTmp];
    int count = 0;
    for (int i = 0; i < sizeS; i++)
    {
        if (s[i] == '\0')
        {
            break;
        }
        if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9'))
        {

            tmp[count] = tolower(s[i]);
            count++;
        }
    }
    for (int i = 0; i < sizeTmp; i++)
    {

        if ((sizeTmp - 1 - i) - i < 2)
        {
            break;
        }
        if (tmp[i] != tmp[sizeTmp - 1 - i])
        {
            return 0;
        }
    }
    return 1;
}

int main(void)
{

    char s1[2] = {' '};
    char s2[5] = {'h', 'e', 'e', 'H'};
    char s3[6] = {'h', 'e', 'e', ':', 'h'};

    printf("Result: %d\n", isPalindrome(s1));
    printf("Result: %d\n", isPalindrome(s2));
    printf("Result: %d\n", isPalindrome(s3));

    return 0;
}