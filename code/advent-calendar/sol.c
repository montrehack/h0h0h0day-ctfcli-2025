#include<stdlib.h>
#include<stdio.h>
#include<unistd.h>
#include<sys/time.h>

int main(int argc, char **argv) {
	alarm(2);
#if 0 // OR
	struct itimerval titi = {{0,0},{2,0}};
	setitimer(ITIMER_REAL, &titi, NULL);
#endif
	execvp(argv[1], argv+1);
	perror(NULL);
}
