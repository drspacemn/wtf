#!/bin/sh


SUB_TOPIC="is_$1"
SUB_TOPIC="$SUB_TOPIC/$2"
LANGUAGE="$3"
WRK_DIR=$(pwd)

mkdir -p $SUB_TOPIC
SPLIT=(${SUB_TOPIC//// })

CURR_SUBS=$(grep 'Subtopics' "${SPLIT[0]}/README.md" | cut -d " " -f 2-)

if [[ ! "$CURR_SUBS" == *"${SPLIT[1]}"* ]]; then
    NEW_SUBS="\\${CURR_SUBS:P:L-1}"
    if [ ${#CURR_SUBS} -ge 3 ]; then
        NEW_SUBS+=","
    fi

    NEW_SUBS+="${SPLIT[1]}\]"
    SUB_SUBS="sed -i 's/Subtopics: \\${CURR_SUBS:P:L-1}\\]$/Subtopics: $NEW_SUBS/' ${SPLIT[0]}/README.md"
    bash -c "$SUB_SUBS"
fi

initGolang() {
    bash -c "sed -i  '/^#### Examples:.*/a - \[golang\](./$SUB_TOPIC/go)' ${SPLIT[0]}/README.md"
    mkdir -p $SUB_TOPIC/go
    cd $SUB_TOPIC/go; touch main.go
    bash -c "go mod init github.com/drspacemn/wtf/$SUB_TOPIC/go"
    bash -c "go mod tidy"
    echo -e 'package main\n\nimport (\n\t"fmt"\n)\n\nfunc main() {\n\n}' > main.go
    cd $WRK_DIR
}

initRustlang() {
    bash -c "sed -i  '/^#### Examples:.*/a - \[rust\](./$SUB_TOPIC/rust)' ${SPLIT[0]}/README.md"
    cd $SUB_TOPIC
    cargo new rust
    cd $WRK_DIR
}

initPython() {
    bash -c "sed -i  '/^#### Examples:.*/a - \[python\](./$SUB_TOPIC/python)' ${SPLIT[0]}/README.md"
    mkdir -p $SUB_TOPIC/python
    touch $SUB_TOPIC/python/main.py
}

case "${LANGUAGE,,}" in 
    "all")
        initGolang
        initRustlang
        initPython
        ;;
    "rust")
        initRustlang
        ;;
    "python")
        initPython
        ;;
    *)
        initGolang
        ;;
esac
