#!/bin/sh

TOPIC="$1"
TOPIC_PATH="is_$TOPIC"
shift
ALL_ARGS="$@"
ALL_ARGS=$(echo "$ALL_ARGS" | tr " " ",")
TITLE=""

mkdir -p $TOPIC_PATH
touch $TOPIC_PATH/README.md

if [[ $1 == *"a"* ]]; then
    SPLIT=(${TOPIC_PATH//_/ })
    TITLE="${SPLIT[1]}"
    TITLE+=" ${SPLIT[2]^}"
else
    for i in ${TOPIC//_/ }
    do
        TITLE+=" ${i^}"
    done
fi

echo -e "## WTF is$TITLE?\n\n" >> $TOPIC_PATH/README.md 
echo -e "#### Overview:\n\n" >> $TOPIC_PATH/README.md 
echo -e "#### Notes:\n\n" >> $TOPIC_PATH/README.md 
echo -e "#### Sources:\n\n" >> $TOPIC_PATH/README.md 
echo -e "#### Examples:\n\n<hr>" >> $TOPIC_PATH/README.md 
echo -e "Subtopics: []\n" >> $TOPIC_PATH/README.md
echo -e "Tags: [$ALL_ARGS]" >> $TOPIC_PATH/README.md
