![Logo](https://github.com/KS-the-visionary/Krystal-Script/blob/main/Logo.png)


# Krystal Script
## A Modern Alternative To Shell Scripting

### About Krystal Script
Krystal Script is a modern alternative to Bash scripting on linux and powershell for windows.
Krystal Script is not a replacement for either of the shells, but as an easier alternative for beginners.
A lot of beginners are scared to use the terminal.
Krystal Script aims to solve that issue.
In Krystal Script, you write code in plain english, and the Krystal Transpiler magically generates respective script files for the target Operating system.
It produces a shell script on linux(`.sh`) and a powershell script on windows(`.ps1`)
It makes development easier, because you can write code once-in Krystal Script and target both the powershell and bash at the same time.

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

## History of Krystal Script
It was during the pandemic in 2020 that I decided to switch to Linux. Ubuntu was the first distribution I tried out. I had no idea what I was getting into. Whenever I encountered an obstacle on my Linux journey, I would search for solutions online, and most of them seemed to involve using the terminal. People rarely presented the GUI (Graphical User Interface) way of doing things. I was hesitant to use the terminal because I was afraid I might accidentally damage my computer if I made a mistake.

Only later did I realize why Linux tutorials often explain things using the terminal instead of showing the GUI method, even though a GUI option may exist. Linux has numerous Desktop Environments, and what works on one distribution might not work on others. By using the terminal, the solutions remain consistent across the multitude of distributions.

I decided to address this problem once and for all. I knew that the only way to eliminate the fear of using the terminal for people was to make it more appealing and user-friendly. No complicated syntax or unfamiliar symbols, just simple and easy-to-use commands. That's why I embarked on creating my own version of shell scripting.

The initial alpha version of Krystal Script was called "bashX" and was written in C++, targeting only the Linux shell.

Later on, I realized that many Windows users were also afraid of using the PowerShell or the Command Prompt. This intrigued me, and I thought things could be even simpler if I could develop a unified shell scripting language that worked out of the box on both Linux and Windows. Consequently, I abandoned bashX and decided to drop C++.

That's when Krystal Script was born—a modern, unified alternative for Windows and Linux that is extremely user-friendly. Believe it or not, you can learn all of it in just a weekend!
