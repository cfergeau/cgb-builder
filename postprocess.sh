# This script reorders the fields in the generated json file so that they are
# in the same order as in arkhamdb json
cat $1 |jq  --indent 4 '[.[] | { code, name, text, traits, slot }]'  >$1.jq && mv $1.jq $1
