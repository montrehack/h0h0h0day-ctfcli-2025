# Advent Calendar, a classic operating-system tale

The source and binary are given, and a low-privileged user.
Getting the flag thought natural ways requite waiting for the next Christmas, that might be inconvenient in a time-limited CTF environment.
There seem to be no other approaches given the program: it's short, tight and robust.
The C fanciness in the code are just here to keep people entertained.

If people RTFM correctly, they might read man setitimer(3) and find

> A child created via fork(2) does not inherit its parent's interval timers.  Interval timers are preserved across an execve(2).

A likely strategy is thus to setup a sort-timed itimer, alarm or whatever, then execve the program. Linux preserve the itimer even for a setuid executable.

Problem, the program override the timer with a new one, so we still have to wait, unless we are naughty and don't have a tty.
The simplest way to detach from the controlling terminal is to use setsid.

```
setsid ./sol ./advent-calendar
```
