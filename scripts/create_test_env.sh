#!/bin/sh

touch tests/holgersyncfile.yml

cat <<EOT > tests/source_test.json
{
   "name": "Foo",
   "description": "Bar"
}
EOT

cat <<EOT > tests/holgersyncfile.yml
HolgersyncConfig:
   SourceFileConfig:
      filePath: tests/source_test.json

   Targets:
EOT

for ((i = 0; i < $1; i++)); do

   cat <<EOT >> tests/holgersyncfile.yml
   - path: tests/test_folder_$((i+1))
     gitConfig:
     - username: $2
       personalAccessToken: $3
       remote: origin
EOT

   mkdir -p tests/test_folder_$((i+1))
   
done