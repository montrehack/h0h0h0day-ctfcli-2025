# Santa's helper
##### Author: Atlantis11656

## Solution
We are given 3 options: List the addresses, Unlock an address, and Leave a present.

When we decompile the binary and we check the "write_present" function, we see that we ask for an "address", aka a file descriptor, and that we use "dup2" to choose that file descriptor. We are not allowed to write to file descriptors 0, 1, and 2 (stdin, stdout, stderr), and any other file descriptor needs to already be opened. 

The "unlock_address" function allows us to open a file descriptor at a given address. However, the only file that we have access to is "/dev/null".

That only leaves us with the option of writing to stdout or stderr if we want to see the flag.

We have clues throughout the binary that can lead us to see that the "address" we choose is stored in a 8-bit integer. This means that the maximum value we can use is 255. However, we do not have any checks to prevent us from using values larger than that.

So by exploiting an integer overflow, we can choose the address 257 (100000001 trucated to 8 bits -> 00000001) to write the flag to file descriptor 1 (stdout).