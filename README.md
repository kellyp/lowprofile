# Low Profile

Simple profile management for AWS.

[![Build Status](https://ci.dualspark.com/api/badge/github.com/DualSpark/lowprofile/status.svg?branch=master)](https://ci.dualspark.com/github.com/DualSpark/lowprofile)

## Install

Installation is super simple, we are using [Brew](http://brew.sh).

NOTE: While in development the Low Profile Brew formula is being hosted outside of the Brew project.

```bash
$ brew update
$ brew install https://s3-us-west-2.amazonaws.com/performance-tires/brew/latest/lowprofile.rb
```

After the brew install finishes you will be advised to update your zshrc or bash_profile with the following:

```bash

  . $(brew --repository)/etc/lowprofile

```

which for example can done without editing the `.zshrc` file with the following:

```bash

echo ". $(brew --repository)/etc/lowprofile" >> ~/.zshrc

```

or `.bash_profile`:

```bash

echo ". $(brew --repository)/etc/lowprofile" >> ~/.bash_profile

```

you can also add the `AWS_PROFILE` variable to your environment file if it
isn't there already (lowprofile will do this automatically the first time you activate
a profile).

```bash

echo "export AWS_PROFILE=default" >> ~/.zshrc

```

As always you can use the [AWS cli](http://aws.amazon.com/cli/) to add or update profiles.  

```bash

$ aws configure --profile different-profile

AWS Access Key ID [None]: lettersandnumbers
AWS Secret Access Key [None]: morelettersandnumbers
Default region name [None]: us-west-2
Default output format [None]: json

```

## Usage

Now that Low Profile is installed and you have a default AWS profile congfigured, you can now describe-profiles, activate-profile, describe-active-profile and deactivate-profile the currently active  profile.  


```bash
$ lowprofile describe-profiles
default
different-profile

$ lowprofile describe-active-profile

$ lowprofile activate-profile --profile default
activating profile default

$ lowprofile describe-active-profile
current profile is default

$ lowprofile activate-profile --profile different-profile
activating profile different-profile

$ lowprofile describe-active-profile
current profile is different-profile

$ lowprofile deactivate-profile
deactiving profile different-profile

$ lowprofile describe-active-profile
there is currently no active profile

```

## What's going on?

Interaction with the [AWS APIs](https://aws.amazon.com/documentation/) can be done by calling them directly, through a library or from the command line.  The libraries and command line support use of several environment variables for auth'ing with AWS (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SECURITY_TOKEN and AWS_PROFILE).  Low Profile uses the AWS_PROFILE variable for activating profiles inside the ~/.aws/crendentials file.  You can still use the AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SECURITY_TOKEN variables to override what's in the credentials file.  Low Profile uses your shell's login init file, currently `~/.bash_profile` or `~/.zshrc`.  After a `lowprofile` command is run your shell session will be reset to pick change to the AWS_PROFILE variable.
