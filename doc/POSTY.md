# Low Profile - The AWS Profile Activator

Hi, we here at DualSpark feel your pain if you've ever worked with multiple AWS
accounts.  That's why we've come up with a really simple tool for management of multiple AWS profiles from the cli.  Low Profile (lowprofile) isn't magical it simply manages your ~/.aws/ files and associated environment variables.  Nothing you couldn't do manually just a simplified interface.  Today we are happy to be releasing and open sourcing Low Profile.  Feel free to create issues, features or just use it.  We hope it makes life a little easier when working with multiple AWS environments.

## Source

The project is hosted on [Github](project link) and is written in [Go](https://golang.org).  Using Go gave me a great reason to dig in and start using Go.

## Installation

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

which for example can done without editing the .zshrc file with the following:

```bash

echo ". /usr/local/etc/lowprofile" >> ~/.zshrc

```

or `.bash_profile`:

```bash

echo ". /usr/local/etc/lowprofile" >> ~/.bash_profile

```

you can also add the `AWS_DEFAULT_PROFILE` variable to your environment file if it
isn't there already (lowprofile will do this automatically the first time you activate
a profile).

```bash

echo "export AWS_DEFAULT_PROFILE=default" >> ~/.zshrc

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

### Prompt v2

As you can see lowprofile makes it super easy to manage profiles from the command line.  But that's not all, as an added bonus we are adding prompt support.  So you can see your current AWS profile in your prompt.  Simply use Low Profile to update the prompt for you.  

First make sure you have an active profile.
```bash
$ lowprofile activate-profile --profile default

```

Next just let Low Profile trick out your prompt for you.

```bash
$ lowprofile activate-prompt
Updating your prompt.

aws(default) $ lowprofile activate-profile --profile different-profile
...
aws(different-profile) $ lowprofile deactivate-profile
...
aws() $ lowprofile deactivate-prompt
Updating your prompt.

$
```

The goal of Low Profile is to make it easier to set your active AWS profile and
easily see which one is selected.  We think we have achieved that goal, please let us know what you think.

##License

Free and Open 2015
