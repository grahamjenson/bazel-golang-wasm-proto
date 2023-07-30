The symlinks in this directory are required to have `go mod tidy` work on the
bazel directory structure.

For as long as bazel does not change the internal paths for the generated
protobufs, this should work.
