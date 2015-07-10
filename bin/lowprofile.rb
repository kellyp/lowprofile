# Documentation: https://github.com/Homebrew/homebrew/blob/master/share/doc/homebrew/Formula-Cookbook.md
#                /usr/local/Library/Contributions/example-formula.rb

class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://s3-us-west-2.amazonaws.com/performance-tires/releases/dev/lowprofile-0.1.tar.gz"
  version "0.1"
  sha256 "05dbd93af64562179289b3ada346efe09895a4be31939a309a1b550c052e8e0f"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath

    system "go", "get", "-d", "github.com/kellyp/lowprofile/lowprofile"
    # Build and install lowprofile
    system "go", "build", "-v", "-o", "./bin/lowprofile-#{version}", "github.com/kellyp/lowprofile/lowprofile"

    bin.install Dir["bin/*"]
    etc.install Dir["etc/*"]
  end

  def caveats; <<-EOS.undent
    Add the following to your bash_profile or zshrc to complete the install:

      . /usr/local/etc/lowprofile

    and source the file to pick up the change.

    if you don't already have it in there feel free to add (if not lowprofile
    will append it for you):

      export AWS_DEFAULT_PROFILE=default


    that's it lowprofile with take it from there!
    You can now switch AWS profiles simply by typing

      lowprofile activate-profile --profile new-profile

    EOS
  end

  test do
    system "false"
  end
end
