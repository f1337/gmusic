class Gmusic.Go < Formula
  desc "A command-line remote for the unofficial Google Play Music Desktop Player."
  homepage "https://github.com/f1337/gmusic.go"
  url "https://github.com/f1337/gmusic.go/releases/download/v0.0.1/gmusic.go_0.0.1_macOS_64-bit.tgz"
  version "0.0.1"
  sha256 "889173a7ae533f0183bd2ee0207a910b09c489b1da3420196f0a85c73a68b34d"

  def install
    bin.install "gmusic.go"
  end
end
