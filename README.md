![Logo](https://github.com/KS-the-visionary/Krystal-Script/blob/main/Logo.png)


# Krystal Script
## A Modern Alternative To Shell Scripting

### About Krystal Script
Krystal Script is a modern alternative to Bash scripting on linux and powershell for windows.
Krystal Script is not a replacement for either of the shells, but as an easier alternative for beginners.
A lot of beginners are scared to use the terminal.
Krystal Script Aims to solve that issue.
In Krystal Script, you write code in plain english, and the Krystal Transpiler magically generates respective script files for the target Operating system.

### Purpose:
Krystal Script was not created as a replacement for bash.
The whole idea behind this project is to reduce the learning curve for beginners.

In Krystal Script the code is written in english.

The code you write in Krystal Script gets translated to native shell commands and a `.sh` file is produced, if you are on linux, otherwise a `.bat` file is produced.

You can then run these files, or pass it onto others.


## Features

- English like syntax
- Can produce both `.sh` and `.ps1` files
- Supports interop with regular shell commands
- Write once, transpile to both linux and windows
- Open Source 😉

## Limitations
Krystal Script is only a psuedo-language. What I mean is that-Krystal Script uses Regex to parse the code and translate it into shell script.
And because of that the parsing capabilities of the transpiler are somewhat limited. Hence, the syntax has to strictly be followed.
Terse one-liners aren't really doable in Krystal Script.
But, sticking to the syntax improves redability and hence is a fair trade. 🙃
