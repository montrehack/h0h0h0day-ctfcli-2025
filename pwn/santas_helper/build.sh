gcc -std=c11 -O2 \
    -Wall -Wextra -Wpedantic -Werror \
    -Wshadow -Wsign-conversion -Wformat=2 \
    -Wnull-dereference -Wduplicated-cond -Wduplicated-branches \
    -Wlogical-op -Wvla -Walloca \
    -fstack-protector-strong \
    -D_FORTIFY_SOURCE=3 \
    -fPIE -pie \
    -fno-common \
    -fcf-protection=full \
    -Wl,-z,relro,-z,now,-z,noexecstack \
    chal.c -o public/chal