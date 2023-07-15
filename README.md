![Logo](https://github.com/KS-the-visionary/Krystal-Script/blob/main/Logo.png)


# Krystal Script
## A Modern Alternative To Shell Scripting

### About Krystal Script
Krystal Script is a modern alternative to Bash scripting on linux and powershell for windows.
Krystal Script is not a replacement for either of the shells, but as an easier alternative for beginners.
A lot of beginners are scared to use the terminal.
Krystal Script Aims to solve that issue.
In Krystal Script, you write code in plain english, and the Krystal Transpiler magically generates respective script files for the target Operating system.
It produced a shell script on linux(`.sh`) and a powershell script on windows(`.ps1`)
It makes development easier, because you can write code once in Krystal Script and target both windows and linux shells at the same time.

## Features
- English like syntax
- Write once, transpile to both linux and windows, Krystal Transpiler can produce both `.sh` and `.ps1` files
- Supports direct interop with regular bash and powershell commands
- Written in Go
- Fully Open Source 😉

## Limitations
Krystal Script is only a psuedo-language. What I mean is that-Krystal Script uses Regex to parse the code and translate it into shell script.
And because of that the parsing capabilities of the transpiler are somewhat limited. Hence, the syntax has to strictly be followed.
Terse one-liners aren't really doable in Krystal Script.
But, sticking to the syntax improves redability and hence is a fair trade. 🙃
