# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class AwsSso < Formula
  desc "AWS account credentials using SSO"
  homepage "https://github.com/nanih98/aws-sso"
  version "0.0.3"
  license "Apache 2.0 license"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.3/aws-sso_0.0.3_Darwin_arm64.tar.gz"
      sha256 "ce79ad4c089baeb2c5a0802acb98614f674b30cb67d5c9cbae5251c67a01009d"

      def install
        bin.install "aws-sso"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.3/aws-sso_0.0.3_Darwin_x86_64.tar.gz"
      sha256 "a04a49579d6dc585b57ff18afdc31db0f78ceb76362a841b664065f2188fbf3c"

      def install
        bin.install "aws-sso"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.3/aws-sso_0.0.3_Linux_x86_64.tar.gz"
      sha256 "d82c52f2852e792796368a8dc0ef0944feb3c67dd0ca349567efd058e3b8a6fd"

      def install
        bin.install "aws-sso"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.3/aws-sso_0.0.3_Linux_arm64.tar.gz"
      sha256 "12170a54623f9a446ab2c985a5b6a8bc0e00321926ceaab87db553b4ef175f8b"

      def install
        bin.install "aws-sso"
      end
    end
  end
end
