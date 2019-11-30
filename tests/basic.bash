#!/usr/bin/env bash

: << USAGE
# shoud be used with: [basht](github.com/progrium/basht)
# install with go: 
#   go get github.com/progrium/basht
# run the tests from the root dir:

  basht tests/basic.bash

USAGE

T_simple_match() { 
  yaml-patch -o <(cat <<END-OF-PATCH
- op: remove
  path: /bar
END-OF-PATCH
) << EOF
foo: 1
bar: 42
EOF
}

T_simple_non_matching() { 
  yaml-patch -o <(cat <<END-OF-PATCH
- op: remove
  path: /bar
END-OF-PATCH
 ) << EOF
foo: 1
EOF
}

T_multi_doc_with_match() {
yaml-patch -o <(cat <<END-OF-PATCH
- op: remove
  path: /bar
END-OF-PATCH
 ) << EOF
foo: 1
bar: 42
---
foo: 1
bar: 42
EOF
}

T_multi_doc_with_partial_match() {
yaml-patch -o <(cat <<END-OF-PATCH
- op: remove
  path: /bar
END-OF-PATCH
 ) << EOF
foo: 1
bar: 42
---
foo: 1
EOF
}
