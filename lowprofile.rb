class Lowprofile < Formula
  desc "Allows for easy switching between AWS profiles"
  homepage "https://github.com/kellyp/lowprofile"
  url "https://github.com/kellyp/lowprofile/archive/v0.5.0.tar.gz"
  version "0.5.0"
  sha256 "af7133ffdbe8a39696d14f9d592fd46661f72eb0c7a5d4ffaddcb4d5ba2b2823"

  depends_on "jq"
  depends_on "grep"

  def install
    bin.install Dir["bin/lowprofile"]
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
