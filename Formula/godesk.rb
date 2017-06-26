class Godesk < Formula
  desc "A quick Desk workspace launcher with fuzzy filtering"
  homepage "https://github.com/hamin/godesk"
  url "https://github.com/hamin/godesk/releases/download/v0.1.0/godesk_0.1.0_darwin_amd64.tar.gz"
  version "0.1.0"
  sha256 "c2558e602e767303ccae791199d35e0082ad2c0f639ebe893f2455ebfdecf480"

  def install
    bin.install "godesk"
  end
end
