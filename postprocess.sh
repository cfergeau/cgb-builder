# This script reorders the fields in the generated json file so that they are
# in the same order as in arkhamdb json
cat $1 |jq  --indent 4 '[.[] | { code, name, text, traits, slot }]'  >$1.jq && mv $1.jq $1

sed -i 's/ \[per_investigator\]/[per_investigator]/' ./translations/fr/pack/tcu/fgg_encounter.json
sed -i 's/(→ <b>/<b>(→/g' ./translations/fr/pack/tcu/fgg_encounter.json
sed -i 's,([skull] / [cultist] / [tablet] / [elder_thing]),([skull]/[cultist]/[tablet]/[elder_thing]),g' ./translations/fr/pack/tcu/fgg_encounter.json
sed -i 's,"<i>\(.*\)</i>","\1",' ./translations/fr/pack/tcu/fgg_encounter.json

[action]: -> [action] :
<b>Discussion</b>. <b>Discussion.</b> 
</b>. -> .</b>
