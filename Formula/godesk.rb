class Godesk < Formula
  desc "A quick Desk workspace launcher with fuzzy filtering"
  homepage "https://github.com/hamin/godesk"
  url "https://github.com/hamin/godesk/releases/download/v0.1.1/godesk_0.1.1_darwin_amd64.tar.gz"
  version "0.1.1"
  sha256 "55f7554426a4d072c92054298e67ff0210fc1e4a845364342cbfa647a11c6058"

  def install
    bin.install "godesk"
  end
end
