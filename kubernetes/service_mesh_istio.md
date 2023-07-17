## Service Mesh

### Overview

A service mesh is a mechanism for managing communication between various invidividual services that make up modern applications in a microservice-based system.
- all inter-service communication is routed through proxies. Networking features such as encrption and load balancing can be applied. 
- decouples network logic from applicaiton or business logic
- like an API gateway, but it only handles communications between system services
- a dedicated infrastructure layer for handling service-to-service communication
- transparently add capabilities like observability, traffic management and security without adding them to your code
- functionalities: discovery, load balancing, failure recovery, metrics, monitoring, A/B testing, canary deployment, rate limiting, access control, encription and e2e authentication. 

### features
- Secure service-to-service communication in a cluster with TLS encryption, strong identity-based authentication and authorization
- Automatic load balancing for HTTP, gRPC, WebSocket, and TCP traffic
- Fine-grained control of traffic behavior with rich routing rules, retries, failovers, and fault injection
- A pluggable policy layer and configuration API supporting access controls, rate limits and quotas
- Automatic metrics, logs, and traces for all traffic within a cluster, including cluster ingress and egress

### How it works
**Control Plane**
- communication between services. Service mesh use a proxy to intercept all your network traffic, allowing a broad set of applicaiton-aware features based on configuration you set. 

**Data Plane**
- translating, forwarding, and monitoring every network packet flowing to and from an instance. 
- features: health checking, routing, service discovery, load balancing, security and telemetry. 
- implementation: prxy instances deployed in a side-car pattern