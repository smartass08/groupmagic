A simple program written in GoLang that allows you to add your emails to gsuite's google group
= 


- Supports adding single email to respected group
- Support from adding all your emails obtained from json provided by google api

 
 Instructions :-
-
 - For single email :-
    +  `./groupmagic single --email "Your@EmailAddress" --group "group@domain.tld"`
        + Where `"Your@EmailAddress"`   is your email address which needs to be added to group & `"group@domain.tld"` is your gsuite group email address
 
-  For multiple emails you have to specify the folder which has all the `.json` files as follows :-
     +  `./groupmagic multiple --SaPath "YourAccountsfolder" --group "group@domain.tld"`
        + Where `"YourAccountsfolder"` is your path to the `.json` files  (default is "accounts" in same folder) & `"group@domain.tld"` is your gsuite group email address

- You are to use `./groupmagic --help` to view all this info again in program.



