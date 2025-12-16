#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/utsname.h>

void print_header(void) {
  printf("  _________              __       ____   ____            .__   __   \n");
  printf(" /   _____/____    _____/  |______\\   \\ /   /____   __ __|  |_/  |_ \n");
  printf(" \\_____  \\\\__  \\  /    \\   __\\__  \\\\   Y   /\\__  \\ |  |  \\  |\\   __\\\n");
  printf(" /        \\/ __ \\|   |  \\  |  / __ \\\\     /  / __ \\|  |  /  |_|  |  \n");
  printf("/_______  (____  /___|  /__| (____  /\\___/  (____  /____/|____/__|  \n");
  printf("        \\/     \\/     \\/          \\/             \\/                 \n");
}

// my_password_is_hunter2?!
const char password_check[] = {0x5b,0x57,0x6b,0x5e,0x50,0x45,0x5e,0x47,0x5f,0x42,0x55,0x6d,0x44,0x14,0x6b,0x51,0x44,0x5a,0x11,0x04,0x44,0x53,0x0c,0x10};
// flag-i_use_linux_compiled_in_wasm_in_my_browser_btw
const char enc_flag[] = {0x0b,0x15,0x3e,0x17,0x4c,0x1a,0x2c,0x02,0x1c,0x17,0x3b,0x33,0x00,0x1d,0x2a,0x10,0x2a,0x0d,0x1b,0x08,0x02,0x5b,0x53,0x44,0x09,0x26,0x36,0x1e,0x3e,0x04,0x12,0x04,0x02,0x2d,0x0d,0x31,0x36,0x1e,0x26,0x37,0x17,0x1c,0x1b,0x12,0x01,0x57,0x4d,0x7e,0x0f,0x0d,0x28};

void _start(void) {
  print_header();
  struct utsname uts;
  uname(&uts);

  int is_correct = 1;
  char password[25];
  char dec_flag[51];

  while (1) {
    printf("[*] ENTER PASSWORD> ");
    fflush(stdout);
    fgets(password, 25, stdin);

    if (strlen(password) != 24) {
      printf("[!] INCORRECT PASSWORD\n");
      continue;
    }

    for (int i = 0; i < strlen(password); i++) {
      if ((password_check[i] ^ password[i]) != uts.release[i]) {
        printf("[!] INCORRECT PASSWORD\n");
        is_correct = 0;
        break;
      }
      else {
        is_correct = 1;
      }
    }

    if (!is_correct) {
      continue;
    }

    for (int i = 0; i < 51; i++) {
      dec_flag[i] = enc_flag[i] ^ password[i % strlen(password)];
    }

    printf("[+] CORRECT PASSWORD, FLAG: %s\n", dec_flag);
    break;
  }

  while (1) {
    getchar();
  }
}
