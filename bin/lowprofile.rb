require "language/go"

class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://github.com/kellyp/lowprofile/archive/383c6f77c9f5e7b8e7d2f014b8cece384f62e194.tar.gz"
  version "0.1"
  sha256 "1db22337db30079568690d6f22e267385e28ca5ac0038cb25deb4b5b9e2d5a60"

  depends_on "go" => :build

  go_resource "github.com/kellyp/lowprofile" do
    url "https://github.com/kellyp/lowprofile.git", :revision => "383c6f77c9f5e7b8e7d2f014b8cece384f62e194"
  end

  def install
    ENV["GOPATH"] = buildpath
    Language::Go.stage_deps resources, buildpath/"src"

    # Build and install lowprofile
    system "go", "build", "-v", "-o", "./bin/lowprofile-#{version}", "main.go"

    bin.install Dir["bin/lowprofile-#{version}"]
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
    system "#{bin}/lowprofile-#{version}", "--help"
  end
end
