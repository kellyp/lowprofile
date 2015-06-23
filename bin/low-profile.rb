# Documentation: https://github.com/Homebrew/homebrew/blob/master/share/doc/homebrew/Formula-Cookbook.md
#                /usr/local/Library/Contributions/example-formula.rb

class LowProfile < Formula
  desc ""
  homepage ""
  url "https://s3-us-west-2.amazonaws.com/performance-tires/releases/latest/low-profile-0.1.tar.gz"
  version "0.1"
  sha256 "ea4529397597dc8989d236b044feaf2a9812574f5e9f1da81ba1cac5b94212ad"


  def install
    bin.install Dir["bin/*"]
    etc.install Dir["etc/*"]
  end

  def caveats; <<-EOS.undent
    Add the following to your bash_profile or zshrc to complete the install:

      . /usr/local/etc/low-profile/low-profile

    and source the file to pick up the change.

    if you don't already have it in there feel free to add (if not low-profile
    will append it for you):

      export AWS_DEFAULT_PROFILE=default


    that's it low-profile with take it from there!
    You can now switch AWS profiles simply by typing

      low-profile activate-profile --profile new-profile

    EOS
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! It's enough to just replace
    # "false" with the main program this formula installs, but it'd be nice if you
    # were more thorough. Run the test with `brew test low-profile`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
