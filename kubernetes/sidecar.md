## Overview

The sidecar pattern is a way to run multiple containers in a single pod. The sidecar pattern is useful when you want to run a container that does auxiliary tasks for your main application. For example, your main application might be a web server, and the sidecar might provide metrics or logging for the main application.

Applications and services often require related functionality, such as monitoring, logging, configuration, and networking services. These peripheral tasks can be implemented as separate components or services.

- Most of the time these sidecar containers are simple and small that consume fewer resources than the main container.

There are other patterns that are useful for everyday kubernetes workloads

Init Container Pattern
Adapter Container Pattern
Ambassador Container Pattern