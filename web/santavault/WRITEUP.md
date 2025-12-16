# SantaVault

The website is a running a linux kernel compiled in WASM. You can see the source of the image in the file `initramfs.cpio.gz`. In the archive, the file `/init` mounts a few directories before executing the file `/sbin/init`.

Since is it WASM, you can decompile it using `wabt` with the command `wasm-decompile --enable-threads [FILE]`. This gives us a C-like code output. There is a lot of things that we do not care for, so here's the main thing we need:

```c
// CUT

export global password_check:int = 2784;
export global enc_flag:int = 2816;

// CUT

data rodata_L_str_6(passive) = "nan\00[*] ENTER PASSWORD> \00";
data rodata_L_str_7(passive) = "NAN\00";
data rodata_L_str_8(passive) =
".\00[+] CORRECT PASSWORD, FLAG: %s\0a\00";
data rodata_password_check(passive) =
"[Wk^PE^G_BUmD\14kQDZ\11\04DS\0c\10";
data rodata_enc_flag(passive) =
  "\0b\15>\17L\1a,\02\1c\17;3\00\1d*\10*\0d\1b\08\02[SD\09&6\1e>\04\12\04"
  "\02-\0d16\1e&7\17\1c\1b\12\01WM~\0f\0d(";
data rodata_Lstr(passive) =
  "  _________              __       ____   ____            .__   __   \00"
  "";
data rodata_Lstr_9(passive) =
  " /   _____/____    _____/  |______\   \ /   /____   __ __|  |_/  |_ \00"
  "";
data rodata_Lstr_10(passive) =
  " \_____  \\__  \  /    \   __\__  \\   Y   /\__  \ |  |  \  |\   __\\00"
  "";
data rodata_Lstr_11(passive) =
  " /        \/ __ \|   |  \  |  / __ \\     /  / __ \|  |  /  |_|  |  \00"
  "";
data rodata_Lstr_12(passive) =
  "/_______  (____  /___|  /__| (____  /\___/  (____  /____/|____/__|  \00"
  "";
data rodata_Lstr_13(passive) =
  "        \/     \/     \/          \/             \/                 \00"
  "";

// CUT

export function print_header() {
  var a:int = memory_base;
  puts(a + 2867);
  puts(a + 2936);
  puts(a + 3005);
  puts(a + 3074);
  puts(a + 3143);
  puts(a + 3212);
}

export function start() {
  var f:int;
  var e:int;
  var a:int_ptr = stack_pointer - 496;
  stack_pointer = a;
  var b:int = memory_base;
  puts(b + 2867);
  puts(b + 2936);
  puts(b + 3005);
  puts(b + 3074);
  puts(b + 3143);
  puts(b + 3212);
  uname(a + 106);
  var c:byte_ptr = GOT_data_internal_stdin[0]:int;
  var d:int = GOT_data_internal_stdout[0]:int;
  loop L_a {
    printf(memory_base + 2719, 0);
    fflush(d);
    fgets(a + 80, 25, c);
    if (strlen(a + 80) == 24) goto B_b;
    loop L_c {
      b = memory_base;
      puts(b + 3281);
      printf(b + 2719, 0);
      fflush(d);
      fgets(a + 80, 25, c);
      if (strlen(a + 80) != 24) continue L_c;
    }
    label B_b:
    b = -24;
    loop L_f {
      e = a + 80 + b;
      if (
        ((e + 24)[0]:ubyte ^ (memory_base + 2784 + b + 24)[0]:ubyte) !=
        ((f = a + 106 + b) + 154)[0]:ubyte) goto B_e;
      if (
        ((e + 25)[0]:ubyte ^ (memory_base + 2784 + b + 25)[0]:ubyte) !=
        (f + 155)[0]:ubyte) goto B_e;
      b = b + 2;
      if (eqz(b)) goto B_d;
      continue L_f;
    }
    label B_e:
    puts(memory_base + 3281);
    continue L_a;
    label B_d:
  }
  b = 0;
  e = strlen(a + 80);
  loop L_h {
    f = memory_base;
    c = a + 16 + b;
    c[0] = (a + 80 + b % e)[0]:ubyte ^ (f + 2816 + b)[0]:ubyte;
    if (b == 50) goto B_g;
    (c + 1)[0]:byte =
      (a + 80 + (b + 1) % e)[0]:ubyte ^
      (memory_base + 2816 + b + 1)[0]:ubyte;
    b = b + 2;
    continue L_h;
  }
  label B_g:
  b = memory_base;
  a[0] = a + 16;
  printf(b + 2746, a);
  loop L_i {
    getchar();
    continue L_i;
  }
}
```

We can see that the code does a few things:
  - Validate if the password is 24 chars long.
  - Compare the password using the `password_check` variable and the uname's version.
  - If the comparison is valid, XOR the password with the `enc_flag` variable and print it.

We can find the uname information by grepping `Linux version` in the `vmlinux.wasm` file.

The password is `my_password_is_hunter2?!` which gives the flag `flag-i_use_linux_compiled_in_wasm_in_my_browser_btw`