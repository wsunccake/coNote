#include <stdio.h>
#include <stdbool.h>
#include <string.h>

bool isUnique(char *str)
{
    int length = strlen(str);

    if (length > 128)
        return false;

    bool char_set[128] = {false};

    for (int i = 0; i < length; i++)
    {
        int val = str[i];
        if (char_set[val])
            return false;
        char_set[val] = true;
    }

    return true;
}

int main()
{
    char *inputs[5] = {"abcde", "hello", "apple", "kite", "padle"};
    bool outputs[5] = {true, false, false, true, true};
    for (int i = 0; i < 5; i++)
    {
        if (isUnique(inputs[i]) != outputs[i])
        {
            printf("%s: %s, %s\n", inputs[i], isUnique(inputs[i]) ? "true" : "false", outputs[i] ? "true" : "false");
        }
    }

    return 0;
}