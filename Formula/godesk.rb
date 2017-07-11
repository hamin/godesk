class Godesk < Formula
  desc "A quick Desk workspace launcher with fuzzy filtering"
  homepage "https://github.com/hamin/godesk"
  url "https://github.com/hamin/godesk/releases/download/v0.1.0/godesk_0.1.0_darwin_amd64.tar.gz"
  version "0.1.1"
  sha256 "c0f4f0bf0154e90ab2aba7c4a6b2e78b5a931658"

  def install
    bin.install "godesk"
  end
end
