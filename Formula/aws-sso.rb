# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class AwsSso < Formula
  desc "AWS account credentials using SSO"
  homepage "https://github.com/nanih98/aws-sso"
  version "0.0.1"
  license "Apache 2.0 license"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.1/aws-sso_0.0.1_Darwin_arm64.tar.gz"
      sha256 "c48e227fad63798cea3837f43ba465774693f2505481e0eb81d5f0b79af50686"

      def install
        bin.install "aws-sso"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.1/aws-sso_0.0.1_Darwin_x86_64.tar.gz"
      sha256 "9bcd1fb37a86614bfe8a51a9326779fd36e392023d26042191b7241e87262ebf"

      def install
        bin.install "aws-sso"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.1/aws-sso_0.0.1_Linux_armv6.tar.gz"
      sha256 "6f59106b865f12240f38f31d4c862e644b5c013f4225da3dff3f7d4a9711cfb5"

      def install
        bin.install "aws-sso"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.1/aws-sso_0.0.1_Linux_x86_64.tar.gz"
      sha256 "c4900b4cd778dc75bc32dc04c2c31fa7c042c8a73910217fc256304cbd0187f1"

      def install
        bin.install "aws-sso"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/nanih98/aws-sso/releases/download/v0.0.1/aws-sso_0.0.1_Linux_arm64.tar.gz"
      sha256 "6ae25d95dade33af89d58d8af105455c74934ef16630a6d6cdebd366743489ca"

      def install
        bin.install "aws-sso"
      end
    end
  end
end
