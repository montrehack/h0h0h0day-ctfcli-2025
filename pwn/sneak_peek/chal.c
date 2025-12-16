#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stddef.h>
#include <unistd.h>
#include <errno.h>

#include <sys/prctl.h>
#include <linux/seccomp.h>
#include <linux/filter.h>
#include <sys/syscall.h>

#include <fcntl.h>
#include <sys/mman.h>

#if !defined filename
#define filename "flag.txt"
#endif

#if !defined limit
#define limit 60
#endif

#define SYS_READ        0
#define SYS_WRITE       1
#define SYS_EXIT_GROUP	231
#define SYS_PROCESS_VM_READV	310
#define SYS_PROCESS_VM_WRITEV	311

struct sock_filter filter[] = {
    BPF_STMT(BPF_LD | BPF_W | BPF_ABS, offsetof(struct seccomp_data, nr)),
    BPF_JUMP(BPF_JMP | BPF_JEQ | BPF_K, SYS_PROCESS_VM_WRITEV, 5, 0),
    BPF_JUMP(BPF_JMP | BPF_JEQ | BPF_K, SYS_PROCESS_VM_READV, 4, 0),
    BPF_JUMP(BPF_JMP | BPF_JEQ | BPF_K, SYS_EXIT_GROUP,       3, 0),
    BPF_JUMP(BPF_JMP | BPF_JEQ | BPF_K, SYS_READ,             2, 0),
    BPF_JUMP(BPF_JMP | BPF_JEQ | BPF_K, SYS_WRITE,            1, 0),
    BPF_STMT(BPF_RET | BPF_K, SECCOMP_RET_ERRNO | (EPERM & SECCOMP_RET_DATA)),
    BPF_STMT(BPF_RET | BPF_K, SECCOMP_RET_ALLOW),
};


struct sock_fprog prog = {
    .len = (unsigned short)(sizeof(filter) / sizeof(filter[0])),
    .filter = filter,
};

void open_and_idle ()
{
	FILE *f;
	f = fopen(filename, "r");
	if (!f) {
		perror("open");
		exit(1);
	}

	for (int i = 0; i < 10; i++) sleep(limit / 10);

	return;
}

void setup_seccomp()
{
	if (prctl(PR_SET_NO_NEW_PRIVS, 1, 0, 0, 0)) {
		perror("prctl(PR_SET_NO_NEW_PRIVS)");
		exit(1);
	}

	if (prctl(PR_SET_SECCOMP, SECCOMP_MODE_FILTER, &prog)) {
		perror("prctl(PR_SET_SECCOMP)");
		exit(1);
	}
}

void read_and_call_shellcode ()
{
	void *buf = malloc(2048);
	if (buf == NULL) {
		perror("malloc");
		exit(1);
	}

	memset(buf, 0, 2048);
	ssize_t ret = read(0, buf, 2048);
	if (ret == 0) {
		perror("read");
		exit(1);
	}

	int res = mprotect((void *)((size_t)buf & 0xfffffffffffff000), 4096, PROT_READ|PROT_WRITE|PROT_EXEC);
	if (res != 0) {
		perror("mprotect");
		exit(1);
	}

	setup_seccomp();
	((void (*) (void)) buf) ();

	return;
}

int main(int argc, char *argv[])
{
	if (argc == 1) {
		execl("/usr/bin/strace",
		      "/usr/bin/strace",
		      "-f",
		      realpath(argv[0], NULL),
		      "neige",
		      NULL);
	}

	setvbuf(stdin, NULL, _IONBF, 0);
	setvbuf(stdout, NULL, _IONBF, 0);

	pid_t pid = fork();
	if (pid == -1) {
		perror("fork failed");
		exit(1);
	}

	if (pid == 0) {
		open_and_idle();
		exit(0);
	}

	read_and_call_shellcode();
	exit(0);
}
