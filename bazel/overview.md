## Overview

Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users.

### How to use Bazel
1. setup bazel: install it
2. setup a project workspace: a directory where bazel looks for build inputs and BUILD files. It also store build output. 
3. write a BUILD file: tell bazel what to build and how to build. 

### Build process
when running a build or test, bazel does the following:
1. load the BUILD files relevant to the target
2. analyze the input and their dependencies, applies the specified build rules and produce a action graph.
3. execute the build action on the inputs until the final build outputs are produced. 