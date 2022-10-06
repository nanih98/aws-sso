```bash
$ aws-sso profile
```

Once your credentials are configured (with the previous symlink) inside your **.aws/credentials** file, this command will read and promt each profile in an interactive terminal list. Then, the profile you selected, will be copied in your clipboard and then you will execute the command **export AWS_PROFILE=yourprofile** You can export the variable directly from a child process. Something like this issue <https://stackoverflow.com/questions/1506010/how-to-use-export-with-python-on-linux>

> Note: in the interactive terminal, you can search for your profiles using SHIFT+/

This command at the end will do nothing, since we can export a variable from a child process. So, other solutions once you have all your credentials in inside `.aws/credentials` is the following:
```bash
aws-profile () {
	PROFILE=$(cat ~/.aws/credentials|grep "^\["|sed "s/]$//"|sed "s/^\[//"| fzf)
	export AWS_PROFILE=$PROFILE
}
```
Save this function inside your `.bashrc` or `.zshrc`.

* Requirements: `fzf` (`brew install fzf`)

Is an interactive menu terminal to change `AWS_PROFILE` at the moment in your terminal.

https://raw.githubusercontent.com/nanih98/aws-sso/main/docs/profile.md

