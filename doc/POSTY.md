# Low Profile - The AWS Profile Activator

Hi, we here at DualSpark feel your pain if you've ever worked with multiple AWS
accounts.  That's why we've come up with a really simple tool for management of multiple AWS profiles from the cli.  Low Profile (low-profile) isn't magical it simply manages your ~/.aws/ files and associated environment variables.  Nothing you couldn't do manually just a simplified interface.  Today we are happy to be releasing and open sourcing Low Profile.  Feel free to create issues, features or just use it.  We hope it makes life a little easier when working with multiple AWS environments.

## Source

The project is host on [Github](project link) and is written in [Go](https://golang.org).  Using Go gave me a great reason to dig in and start using Go.

## Installation

Installation is super simple, we are using [Brew](http://brew.sh) on OS X.

```bash
$ brew update
$ brew install low-profile
```

As always you can use the [AWS cli](http://aws.amazon.com/cli/) to add or update profiles.  

```bash
$ aws configure

AWS Access Key ID [None]: lettersandnumbers
AWS Secret Access Key [None]: morelettersandnumbers
Default region name [None]: us-west-2
Default output format [None]: json
```

## Usage

### Activate

Now that Low Profile is installed and you have a default AWS profile congfigured, you can now activate-profile, describe-active-profile and deactivate-profile the currently active  profile.  


```bash
$ low-profile describe-active-profile

$ low-profile activate-profile --profile default
activating profile default

$ low-profile describe-active-profile
current profile is default

$ low-profile activate-profile --profile different-profile
activating profile different-profile

$ low-profile describe-active-profile
current profile is different-profile

$ low-profile deactivate-profile
deactiving profile different-profile

$ low-profile describe-active-profile
there is currently no active profile

```

### Prompt

As you can see low-profile makes it super easy to manage profiles from the command line.  But that's not all, as an added bonus we are adding prompt support.  So you can see your current AWS profile in your prompt.  Simply use Low Profile to update the prompt for you.  

First make sure you have an active profile.
```bash
$ low-profile activate-profile --profile default

```

Next just let Low Profile trick out your prompt for you.

```bash
$ low-profile activate-prompt
Updating your prompt.

aws(default) $ low-profile activate-profile --profile different-profile
...
aws(different-profile) $ low-profile deactivate-profile
...
aws() $ low-profile deactivate-prompt
Updating your prompt.

$
```

The goal of Low Profile is to make it easier to set your active AWS profile and
easily see which one is selected.  We think we have achieved that goal, please let us know what you think.

##License

Free and Open 2015
