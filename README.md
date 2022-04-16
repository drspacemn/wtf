<div align="center">
    <h1>WTF is wtf?</h1>
    <img src="./assets/nictf.jpeg" style="width: 300px">
</div>

## Overview
WTF is a compendium of useless information. A collection of simple, from first principle building block exercises to help understand the core concepts that constitute a large one.
<br><br>

### Create New Topic
Enter the topic you want to learn about snake_case followed by all applicable tags you can think of(also snake_case)

```bash
// Usage: scripts/wtf_is.sh <TOPIC_NAME> ...<TOPIC_TAGS>
$ bash scripts/wtf_is.sh data_availability ethereum zk zero_knowledge rollups encoding data_structures serialization
```

*NOTE: include a if applicable (e.g. scripts/wtf_is.sh a_runtime ...)*
<br><br>

### Create New Sub Topic
Enter the over-arching topic this will go under, and the name of the sub topic in snake_case followed by the programming language for your demo. Defaults to golang.

```bash
// Usage: scripts/subtopic_wtf_is.sh <TOPIC_NAME> <SUB_TOPIC_NAME> [golang|rust|python|all]
$ bash scripts/subtopic_wtf_is.sh data_availability erasure_coding all
```
