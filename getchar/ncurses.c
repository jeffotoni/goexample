#include <ncurses.h>

// modified halloworld from ncurses tutorial..
// to compile: gcc -o halloinput halloinput.c -lncurses

int main(int argc, char** argv)
{
    char c;

    initscr();
    refresh();
    printw("press a key: ");
    c = getch();
    printw("\nkeycode: %d\n", (int) c);
    printw("press any key to quit..");
    getch();
    endwin();
    return 0;
}
