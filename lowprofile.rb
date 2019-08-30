class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://github.com/kellyp/lowprofile/archive/v0.2.0.tar.gz"
  version "0.2.0"
  sha256 "42e015d4977a27298d58bd6d3735a456af9e77b78bbacc0a37752c9c3c6ff7cf"

  depends_on "jq"

  def install
    bin.install Dir["etc/lowprofile"]
  end

  def caveats; <<-EOS
    Add the following to your bash_profile or zshrc to complete the install:

      . #{HOMEBREW_PREFIX}/bin/lowprofile

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
