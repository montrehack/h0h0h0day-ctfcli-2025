function SANDBOX()
	return "t", { os = {
		clock = os.clock,
		date = os.date,
		difftime = os.difftime,
		time = os.time,
	} }
end

local function readfile(name)
	local f = io.open(name, "r")
	local s = f:read("*a")
	f:close()
	return s
end

local function edit(s)
	io.write("position : ")
	local p = tonumber(io.read())
	if not p or p < 0 then error("invalid number") end
	io.write("character: ")
	local c = io.read()
	if #c ~= 1 then error("not a character") end

	return string.sub(s, 1, p) .. c .. string.sub(s, p+1)
end

io.output():setvbuf("no")

local sandbox = readfile("sandbox.lua")
print("you may insert 2 characters into your jail: "..sandbox)
for i=1,2 do sandbox = edit(sandbox) end
print("your jail: "..sandbox)
io.write("code to run inside the jail: ")
input = io.read()

if not input or string.match(input, "[.\'\"()]") then
	error("illegal character")
end

assert(load(sandbox, "sandbox", "t"))()
