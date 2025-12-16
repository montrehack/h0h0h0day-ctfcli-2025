#include<stdlib.h>
#include<stdio.h>
#include<string.h>
#include<unistd.h>
#include<signal.h>
#include<time.h>
#include<sys/time.h>

char wishes[1024];

#define IS_BLANK(c) ((c) == ' ' || (c) == '\t' || (c) == '\n' || (c) == '\r')

void trim(char *s) {
	char *l = strlen(s) + s;
	while (l --> s && IS_BLANK(*l)) *l = '\0';
}

void xmas(int i) {
	char buf[128] = "lump of coal";
	FILE *f = fopen("flag.txt", "r");
	if(f) {
		fgets(buf, sizeof(buf), f);
		trim(buf);
		fclose(f);
	} else {
		perror("flag.txt");
	}
	printf("Sorry, no %s for you, you got a %s instead. Maybe next year.\n", wishes, buf);
	signal(SIGALRM, xmas); // for next year...
}

void setup(void) {
	signal(SIGALRM, xmas);

	FILE * tty = fopen("/dev/tty", "r+");
	if(!tty) {
		perror("Santa hates naughty children who don't have a tty");
		return;
	}
	while (wishes[0] == '\0') {
		fprintf(tty, "Enter your Santa wish list: ");
		fflush(tty);
		fgets(wishes, sizeof(wishes), tty);
		trim(wishes);
	}
	fclose(tty);
	
	time_t today;
	time(&today);
	struct tm tmoday;
	localtime_r(&today, &tmoday);

	struct tm xmas_day = {0, 0, 0, 25, 11, 0, 0, 0, 0, 0, 0};
       	xmas_day.tm_year = tmoday.tm_year;
       	xmas_day.tm_gmtoff = tmoday.tm_gmtoff;
       	xmas_day.tm_zone = tmoday.tm_zone;
       	xmas_day.tm_isdst = 0;
	time_t xas_day = mktime(&xmas_day);

	if (today + 604800 > xas_day) {
		printf("A little too late for your wishlist. Maybe next year.\n");
		xmas_day.tm_year++;
		xas_day = mktime(&xmas_day);
	}

	time_t diff = xas_day - today;
	struct itimerval titi = {{31536000,0},{diff,0}};
	setitimer(ITIMER_REAL, &titi, NULL);
}

[[noreturn]] void waiting(void) {
	char *emo[] = {"\n", "🎄","🎅","🤶","🦌","⛄","❄️","🎁","🕯️","🌟","✨","🍪","🥛","🎶","🔔","🛷","🎀","🧦","🪅","🪞","🪵","☠️"};
	for(;;) {
		printf("%s", emo[rand()%(sizeof(emo)/sizeof(emo[-1]))]);
		fflush(stdout);
		usleep(250000+rand()%500000);
	}
}

int main(int argc, char **argv) {
	setup();
	waiting();
}
