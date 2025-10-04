#!/usr/bin/env bash
set -e

PKG_MANAGER=$(command -v dnf yum|head -n1)
${PKG_MANAGER} -y install glib2-devel glibc-static glibc-static gcc make golang rpkg go-rpm-macros
