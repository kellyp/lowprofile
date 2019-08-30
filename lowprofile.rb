class Lowprofile < Formula
  desc "Allows for easy switching between AWS profiles"
  homepage "https://github.com/kellyp/lowprofile"
  url "https://github.com/kellyp/lowprofile/archive/v0.2.1.tar.gz"
  version "0.2.1"
  sha256 "01d96374579a6b3553b3cc8df85675eb3c761bc0329f9905dd209fbdf90a4373"

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
