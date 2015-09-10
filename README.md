# Low Profile

Simple profile management for AWS.

[![Build Status](https://ci.kellyp.com/api/badge/github.com/kellyp/lowprofile/status.svg?branch=master)](https://ci.kellyp.com/github.com/kellyp/lowprofile)

## Install

Installation is super simple, we are using [Brew](http://brew.sh) on OS X.

NOTE: While in development the Low Profile Brew formula is being hosted outside of the Brew project.

```bash
$ brew update
$ brew install https://s3-us-west-2.amazonaws.com/performance-tires/brew/latest/lowprofile.rb
```

After the brew install finishes you will be advised to update your zshrc or bash_profile with the following:

```bash

  . /usr/local/etc/lowprofile

```

which for example can done without editing the `.zshrc` file with the following:

```bash

echo ". /usr/local/etc/lowprofile" >> ~/.zshrc

```

or `.bash_profile`:

```bash

echo ". /usr/local/etc/lowprofile" >> ~/.bash_profile

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

### Activate

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
