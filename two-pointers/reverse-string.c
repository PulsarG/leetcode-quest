// https://leetcode.com/problems/reverse-string/description/

void reverseString(char* s, int sSize) {
  for (int i = 0; i < sSize; i++) {
    char temp = s[i];
    s[i] = s[sSize - 1 - i];
    s[sSize - 1 - i] = temp;
    if ((sSize - 1 - i) - i < 2) break;
  }
}