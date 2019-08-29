class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://github.com/kellyp/lowprofile/archive/v0.2.0.tar.gz"
  version "0.2.0"
  sha256 "fa4e89e04bc9b2b71785278bbd26109714371d87b14b6657fe5f1da907035a4b"

  def install
    etc.install Dir["etc/*"]
  end

  def caveats; <<-EOS.undent
    Add the following to your bash_profile or zshrc to complete the install:

      . #{HOMEBREW_PREFIX}/etc/lowprofile

    and source the file to pick up the change.

    if you don't already have it in there feel free to add (if not lowprofile
    will append it for you):

      export AWS_PROFILE=default


    that's it lowprofile with take it from there!
    You can now switch AWS profiles simply by typing

      lowprofile activate --profile new-profile

    EOS
  end
end
