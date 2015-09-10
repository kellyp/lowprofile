require "language/go"

class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://github.com/kellyp/lowprofile/archive/v0.1.1.tar.gz"
  version "0.1"
  sha256 "fa4e89e04bc9b2b71785278bbd26109714371d87b14b6657fe5f1da907035a4b"

  depends_on "go" => :build

  go_resource "github.com/kellyp/lowprofile" do
    url "https://github.com/kellyp/lowprofile.git", :tag => "v0.1.1"
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

      . #{HOMEBREW_PREFIX}/etc/lowprofile

    and source the file to pick up the change.

    if you don't already have it in there feel free to add (if not lowprofile
    will append it for you):

      export AWS_PROFILE=default


    that's it lowprofile with take it from there!
    You can now switch AWS profiles simply by typing

      lowprofile activate-profile --profile new-profile

    EOS
  end

  test do
    system "#{bin}/lowprofile-#{version}", "--help"
  end
end
