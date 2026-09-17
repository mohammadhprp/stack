#!/bin/sh
# Installer for the stack CLI.
#
# Usage:
#   curl -fsSL https://github.com/mohammadhprp/stack/releases/latest/download/install.sh | sh
#   sh install.sh [--dir <dir>]
#   STACK_VERSION=vX.Y.Z sh install.sh
#
# Environment:
#   STACK_VERSION      Release tag to install (default: the latest release).
#   STACK_INSTALL_DIR  Install directory (alternative to --dir).
set -eu

REPO="mohammadhprp/stack"
BINARY="stack"

INSTALL_DIR="${STACK_INSTALL_DIR:-}"

usage() {
	cat <<'EOF'
Install the stack CLI from GitHub Releases.

Usage:
  curl -fsSL https://github.com/mohammadhprp/stack/releases/latest/download/install.sh | sh
  sh install.sh [--dir <dir>]
  STACK_VERSION=vX.Y.Z sh install.sh

Options:
  --dir <dir>   Install directory (default: $HOME/.local/bin)
  -h, --help    Show this help

Environment:
  STACK_VERSION      Release tag to install (default: latest)
  STACK_INSTALL_DIR  Install directory (alternative to --dir)
EOF
}

while [ $# -gt 0 ]; do
	case "$1" in
	--dir)
		if [ $# -lt 2 ]; then
			echo "error: --dir requires a value" >&2
			exit 2
		fi
		INSTALL_DIR="$2"
		shift 2
		;;
	--dir=*)
		INSTALL_DIR="${1#--dir=}"
		shift
		;;
	-h | --help)
		usage
		exit 0
		;;
	*)
		echo "error: unknown argument: $1" >&2
		usage >&2
		exit 2
		;;
	esac
done

if [ -z "$INSTALL_DIR" ]; then
	: "${HOME:?HOME is not set; pass --dir or set STACK_INSTALL_DIR}"
	INSTALL_DIR="$HOME/.local/bin"
fi

case "$(uname -s)" in
Darwin) os="darwin" ;;
Linux) os="linux" ;;
*)
	echo "error: unsupported OS: $(uname -s). stack supports darwin and linux." >&2
	exit 1
	;;
esac

case "$(uname -m)" in
x86_64 | amd64) arch="amd64" ;;
arm64 | aarch64) arch="arm64" ;;
*)
	echo "error: unsupported architecture: $(uname -m). stack supports amd64 and arm64." >&2
	exit 1
	;;
esac

if [ -n "${STACK_VERSION:-}" ]; then
	base_url="https://github.com/${REPO}/releases/download/${STACK_VERSION}"
else
	base_url="https://github.com/${REPO}/releases/latest/download"
fi

archive="${BINARY}_${os}_${arch}.tar.gz"
archive_url="${base_url}/${archive}"
checksums_url="${base_url}/checksums.txt"

if command -v curl >/dev/null 2>&1; then
	download() { curl -fsSL "$1" -o "$2"; }
elif command -v wget >/dev/null 2>&1; then
	download() { wget -q "$1" -O "$2"; }
else
	echo "error: need curl or wget to download stack." >&2
	exit 1
fi

tmp="$(mktemp -d "${TMPDIR:-/tmp}/stack.XXXXXX")"
trap 'rm -rf "$tmp"' EXIT HUP INT TERM

echo "Downloading ${archive}..."
download "$archive_url" "$tmp/$archive"
download "$checksums_url" "$tmp/checksums.txt"

# Find the expected checksum for this archive (tolerate the "hash *file" form).
expected="$(awk -v f="$archive" '{ n = $2; sub(/^\*/, "", n); if (n == f) print $1 }' "$tmp/checksums.txt")"
if [ -z "$expected" ]; then
	echo "error: ${archive} not listed in checksums.txt" >&2
	exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
	actual="$(sha256sum "$tmp/$archive" | awk '{ print $1 }')"
elif command -v shasum >/dev/null 2>&1; then
	actual="$(shasum -a 256 "$tmp/$archive" | awk '{ print $1 }')"
else
	echo "error: need sha256sum or shasum to verify the download." >&2
	exit 1
fi

if [ "$expected" != "$actual" ]; then
	echo "error: checksum mismatch for ${archive}" >&2
	echo "  expected: ${expected}" >&2
	echo "  actual:   ${actual}" >&2
	exit 1
fi
echo "Checksum verified."

tar -xzf "$tmp/$archive" -C "$tmp"

src="$tmp/$BINARY"
if [ ! -f "$src" ]; then
	src="$(find "$tmp" -type f -name "$BINARY" | head -n 1)"
fi
if [ -z "$src" ] || [ ! -f "$src" ]; then
	echo "error: ${BINARY} binary not found in ${archive}" >&2
	exit 1
fi

mkdir -p "$INSTALL_DIR"
cp "$src" "$INSTALL_DIR/$BINARY"
chmod +x "$INSTALL_DIR/$BINARY"

installed="$INSTALL_DIR/$BINARY"
echo "Installed ${BINARY} to ${installed}"

case ":${PATH:-}:" in
*":$INSTALL_DIR:"*) ;;
*)
	echo ""
	echo "Note: ${INSTALL_DIR} is not on your PATH."
	echo "Add it with:"
	echo "  export PATH=\"${INSTALL_DIR}:\$PATH\""
	;;
esac
