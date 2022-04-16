#!/bin/sh

TOPIC="$1"
TOPIC_PATH="is_$TOPIC"
TITLE=""
TAGS=""

curl --request GET --url https://wordsapiv1.p.rapidapi.com/words/hatchback/typeOf --header "X-RapidAPI-Host: $WORDS_HOST" --header "X-RapidAPI-Key: $WORDS_KEY"
RAW_TAGS=$(curl "https://wordsapiv1.p.mashape.com/words/$TOPIC" | jq )

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
echo -e "Synonyms: [$ALL_ARGS]" >> $TOPIC_PATH/README.md
echo -e "TypeOf: [$ALL_ARGS]" >> $TOPIC_PATH/README.md
echo -e "Subtopics:\n" >> $TOPIC_PATH/README.md
