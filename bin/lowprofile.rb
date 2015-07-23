require "language/go"

class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://github.com/kellyp/lowprofile/archive/1e2af55a5102650a4ccc2c296654f713c56ff00a.tar.gz"
  version "0.1"
  sha256 "11ea8e8ef3185276ac9592a10e783e693cfcc27cf227dc4628c52afdf6913635"

  depends_on "go" => :build

  go_resource "github.com/kellyp/lowprofile" do
    url "git@github.com:kellyp/lowprofile.git", :revision => "1e2af55a5102650a4ccc2c296654f713c56ff00a"
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
