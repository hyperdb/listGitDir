# listGitDir

I enjoy a smooth workflow of jumping quickly between repositories managed collectively by the ghq command using peco. However, as I tried out an editor called Void, I thought keeping all my repositories in the same place might make it unclear which editor I used to edit them. So, I decided to prepare a directory specifically for Void. Naturally, I then wanted a command to output a list of subdirectories in the Void directory.

## What is listGitDir?

listGitDir is a command that searches for subdirectories under a specified directory and displays them to standard output. If a subdirectory does not contain a `.git` directory directly under it, it is excluded from the results.

## How to Use

Build the binary and use as follows. For use with peco, error messages (such as when `targetDirectory` does not exist) are not displayed.

```bash
listGitDir [targetDirectory]
```

## Example Usage with Peco

Here is a sample batch file I use. You can select the repository you want to move to from the launched screen, and it will change the directory to that repository.

```bat
@echo off

for /f "tokens=*" %%x in ('listGitDir "D:\My Projects\Void" ^| peco') do (
  d:
  cd %%x
  break
)
```

## License

The source code for this project is licensed under the MIT License. Texts and other content (except source code) are licensed under CC BY 4.0. Please see the "LICENSE" file for details.
