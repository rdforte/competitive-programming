# Competitive Programming 
Contains my answers to problems solved on variouse competitive programming sites

### Setting up C++ in VSCode
[Using Clang in Visual Studio Code](https://code.visualstudio.com/docs/cpp/config-clang-mac)

### Using bits/stdc++.h
bits/stdc++ is a GNU GCC extension, whereas OSX uses the clang compiler. So we will not have access to this header.
In order to add this header we have to first find out where to put it.

run:
```
echo "" | gcc -xc - -v -E
```
This will give you a large sum of output but what we want is to find where it uses the #include
![finding include](./assets/findingInclude.png)

look for the one with `include` at the end.

Once we have this lets go into that directory and make a new directory called __bits__
```
cd /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include
```
```
sudo mkdir bits
```
now go into bits and add [stdc++.h](https://github.com/gcc-mirror/gcc/blob/master/libstdc%2B%2B-v3/include/precompiled/stdc%2B%2B.h)

run the following command in the same directory as the stdc++.h file. Note that the path you are copying should match the 'include' path
you retrieved earlier.
```
sudo cp stdc++.h /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include/bits/stdc++.h
```

note there where a few issues with the __cplusplus preprocessor macro so a had to move a few of the libraries around but once I did
that it seemed to be working fine.

### Note
We use C++ 17 because LeetCode does not support C++ 20


## DSA - Links
### Arrays & Hashing
- [EASY 217. Contains Duplicate](./LeetCode/Easy/217-contains-duplicate/index.cpp)
- [EASY 242. Valid Anagram](./LeetCode/Easy/242-valid-anagram/index.cpp)

## IntelliJ
Sometimes the Vim keybindings clash with IntelliJ keybindings. We can resolve this by following:
https://www.jetbrains.com/help/idea/using-product-as-the-vim-editor.html
