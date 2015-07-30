require "language/go"

class Lowprofile < Formula
  desc ""
  homepage ""
  url "https://github.com/dualspark/lowprofile/archive/b162b9f19313ee879442562b8b06911311faa52a.tar.gz"
  version "0.1"
  sha256 "64109b05badf2fe7000e9e73ce02f3cdbaf4066946c1c193e0e513633e40d021"

  depends_on "go" => :build

  go_resource "github.com/dualspark/lowprofile" do
    url "https://github.com/dualspark/lowprofile.git", :revision => "b162b9f19313ee879442562b8b06911311faa52a"
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
