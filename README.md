This software is under construction!

# Password CLI - Written in Golang

This is a password CLI- for those tired of creating new passwords everytime. What if you had one (or a bunch) of masterPasswords
and you simple would permute them for different websites.

Nothing is currently stored in anywhere, there will be development for exporting the generated password for a txt file later on. The masterpassword is only in your memory.

## How to Use:

Currently we dont have a binary released so you need to compile the project itself from the source code. In order to do that assure you have the latest version of Golang installed in your operating system. Then run on the root of the project the following command:

```
go build
```

This will create a binary executable for your operating system in the current directory. Now it's just the process of running:

```
./passwordCLI -password=YOUR_MASTER_PASSWORD -path=PATH_TO_EXPORT_TXT -copy=false
```

The copy keyword is if you want to copy to clipboard the first password generated.

## Development Todo (For first stable release):

- [x] Better CLI options (export to txt)
- [ ] Limit on permutation items
- [x] Unitary Test and Coverage Tests
- [ ] CI/CD for releases versioning
- [ ] More randomness to the process (better security)
- [x] Option for 1 permutation only
- [x] Copy to clipboard

Update 04/19/2025 : Added some unitary tests and got some interesting bugs fixed, such as whitespaces on password input. Regex did the trick. The ideia now is that every feature added must come with its sets of unitary tests on regular and edge cases to prove efficiency.

Update 04/21/2025: Added function for creating a directory witch will be a feature for exporting csv-txt files and its respectives unitary tests.

Update 04/29/2025: Guess the feature for exporting to txt is ready. Maybe in future we must implement more test on edge cases to avoid breaking the program as whole. The choice base file is txt for simplicity

update 05/17/2025: Added feature for one password only, this generate one random permutation, also need to add unitary test for those. A bit tricky since we have random numbers being generated
