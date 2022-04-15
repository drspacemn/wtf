# WTF is wtf/...?

<div align="center">
    <img src="./assets/nictf.jpeg" style="width: 300px">
</div>

### Overview
<hr>
WTF is a collection of simple, from first principle building block exercises to help understand the core concepts that constitute a large one.

### Create New Topic
<hr>
Enter the topic you want to learn about snake_case followed by all applicable tags you can think of(also snake_case)

Usage: scripts/wtf_is.sh <TOPIC_NAME> ...<TOPIC_TAGS>
```bash
$ bash scripts/wtf_is.sh data_availability ethereum zk zero_knowledge rollups encoding data_structures serialization
```

*NOTE: include a if applicable (e.g. scripts/wtf_is.sh a_runtime ...)*

### Create New Sub Topic
<hr>
Enter the over-arching topic this will go under, and the name of the sub topic in snake_case followed by the programming language for your demo. Defaults to golang.

Usage: scripts/subtopic_wtf_is.sh <TOPIC_NAME> <SUB_TOPIC_NAME> [golang|rust|python|all]
```bash
$ bash scripts/subtopic_wtf_is.sh data_availability erasure_coding all
```
