#include <unistd.h>
#include <fcntl.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <stdlib.h>

int read_address() {
    char buf[128];

    if (!fgets(buf, sizeof buf, stdin)){
        fprintf(stderr, "fgets error: %s\n", strerror(errno));
        exit(1);
    }
    char *end;
    errno = 0;
    long val = strtol(buf, &end, 10);
    if (end == buf){
        fprintf(stderr, "No digits found\n");
        exit(1);
    }
    
    if ((errno == ERANGE)){
        fprintf(stderr, "Value out of range\n");
        exit(1);
    }
    
    if (val < 3){
        fprintf(stderr, "Cannot leave a present at an address under 3\n");
        exit(1);
    }
    while (*end == ' ' || *end == '\t' || *end == '\n')
        end++;
    if (*end != '\0'){
        fprintf(stderr, "Extra characters after number\n");
        exit(1);
    }

    return val;
}

char* read_flag() {
    FILE *f = fopen("/flag", "r");
    if (!f) {
        fprintf(stderr, "Failed to open flag file: %s\n", strerror(errno));
        exit(1);
    }
    char *flag = malloc(42);
    if (!fgets(flag, 42, f)) {
        fprintf(stderr, "Failed to read flag: %s\n", strerror(errno));
        exit(1);
    }
    fclose(f);
    return flag;
}


int fd_is_open(int fd) {
    return fcntl(fd, F_GETFD) != -1;
}

void write_present(){
    char* flag = read_flag();
    printf("Enter the address you want to leave a present to!\n");
    __uint8_t chosen_address = read_address();
    if(!fd_is_open(chosen_address)){
        fprintf(stderr, "Address %d is locked\n", chosen_address);
        exit(1);
    }
    int fd = open("/dev/null", O_WRONLY, 0600);
    dup2(chosen_address,fd);
    int size = write(fd,flag, strlen(flag));
    if(size == -1){
        fprintf(stderr, "Failed to leave present at address %d: %s\n", chosen_address, strerror(errno));
        exit(1);
    }
    free(flag);
    flag = NULL;
    printf("\nPresent left at address %d! Merry christmas!\n", chosen_address);
    exit(0);
}

void unlock_address(){
    printf("Unlocking a new address for you...\n");
    int fd = open("/dev/null", O_WRONLY);
    if(fd == -1 || fd >= 255){
        fprintf(stderr, "Failed to unlock a new address");
        exit(1);
    }
    printf("New address unlocked: %d\n", fd);
}

void print_addresses(){
    printf("Current addresses:\n");
    for(int fd = 0; fd < 255; fd++){
        if(fd_is_open(fd)){
            printf(" - %d | OPEN\n", fd);
        }else{
             printf(" - %d | CLOSED\n", fd);
        }
    }
}

void menu(){
    printf("HOHOHO! Choose an option\n");
    printf("1) List the addresses\n");
    printf("2) Unlock an address\n");
    printf("3) Leave a present\n");
    printf("4) Exit\n");
    printf("Your choice: ");
    char buf[4];
    if (!fgets(buf, sizeof buf, stdin)){
        fprintf(stderr, "fgets error: %s\n", strerror(errno));
        exit(1);
    }
    int choice = atoi(buf);
    switch(choice){
        case 1:
            print_addresses();
            break;
        case 2:
            unlock_address();
            break;
        case 3:
            write_present();
            break;
        default:
            printf("Goodbye!\n");
            exit(0);
            break;
    }
}
int main(void) {
    while (1){    
        menu();
    }
    return 0;
}