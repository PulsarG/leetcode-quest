// https://leetcode.com/problems/valid-palindrome/?envType=problem-list-v2&envId=two-pointers

#include <ctype.h>
#include <stdio.h>

int isPalindrome(char* s) {
  int sizeS = 0;
  int sizeTmp = 0;
  for (int i = 0;; i++) {
    if (s[i] == '\0') {
      break;
    }
    // возможно стоит проверять после перевода в нижний регистр
    if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') ||
        (s[i] >= '0' && s[i] <= '9')) {
      sizeTmp++;
    }
    sizeS++;
  }
  if (sizeTmp < 2) return 1;
  char tmp[sizeTmp];
  int count = 0;
  for (int i = 0; i < sizeS; i++) {
    if (s[i] == '\0') {
      break;
    }
    if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') ||
        (s[i] >= '0' && s[i] <= '9')) {
      tmp[count] = tolower(s[i]);
      count++;
    }
  }

  for (int i = 0; i < sizeTmp; i++) {
    printf("%c\n", tmp[i]);
  }

  for (int i = 0; i < sizeTmp; i++) {
    if (tmp[i] != tmp[sizeTmp - i - 1]) {
      return 0;
    }
    if ((sizeTmp - i - 1) - i < 2) {
      break;
    }
  }
  return 1;
}

int main(void) {
  char s1[3] = {'a', 'b'};
  char s2[5] = {'h', 'e', 'e', 'H'};
  char s3[6] = {'h', 'e', 'd', ':', 'h'};
  char s4[11] = {'r', 'a', 'c', 'e', ' ', 'a', ' ', 'c', 'a', 'r'};

  printf("Result: %d\n", isPalindrome(s1));
  printf("Result: %d\n", isPalindrome(s2));
  printf("Result: %d\n", isPalindrome(s3));
  printf("Result: %d\n", isPalindrome(s4));

  return 0;
}

/* Интересный подход с комментариев на Литкод

    while (left < right) {
  if (s.charAt(left) != s.charAt(right)) {
    return false;
  }
  left++;
  right--;
}
return true;
} */