# 5G Core Network Simulator

A cloud-native 5G Core Network Simulator built with Go and Kubernetes, implementing the Service-Based Architecture (SBA) defined by 3GPP standards.

## Project Overview

This project simulates the key components and interactions of a 5G Core Network, focusing on:

- Service-Based Architecture implementation
- Network Function (NF) service discovery and registration
- UE registration and session management
- Network slicing capabilities
- Policy control and enforcement

The simulator is designed for educational purposes, proof of concept development, and as a platform for testing and experimentation with 5G core network concepts.

## Architecture

The simulator implements the following 5G Core Network Functions:

- **AMF** (Access and Mobility Management Function)
- **SMF** (Session Management Function)
- **UPF** (User Plane Function)
- **PCF** (Policy Control Function)
- **UDM** (Unified Data Management)
- **AUSF** (Authentication Server Function)
- **NRF** (Network Repository Function)
- **NSSF** (Network Slice Selection Function)

Each network function is implemented as a microservice, with RESTful APIs for service-based interfaces and gRPC for performance-critical interfaces.

## Project Structure

```
/
├── cmd/                 # Application entry points
│   ├── amf/             # AMF microservice
│   ├── smf/             # SMF microservice
│   └── ...              # Other network functions
├── pkg/
│   ├── common/          # Shared utilities
│   │   ├── logger/      # Logging framework
│   │   └── config/      # Configuration management
│   ├── models/          # Shared data models
│   └── sbi/             # Service-Based Interface library
├── api/
│   ├── proto/           # Protocol buffer definitions
│   └── openapi/         # OpenAPI specifications
├── deployments/         # Kubernetes manifests
│   ├── base/            # Base configurations
│   └── overlays/        # Environment-specific overlays
├── docs/                # Documentation
├── test/                # Test suite
│   ├── unit/            # Unit tests
│   └── integration/     # Integration tests
└── scripts/             # Helper scripts
```

## Implementation Checklist

### Phase 1: Infrastructure Setup

- [X] Set up Go modules and project structure
- [X] Create Kubernetes namespace and basic resources
- [X] Implement shared logging and metrics libraries
- [X] Set up CI/CD pipeline with GitHub Actions
- [X] Create Docker build pipeline

### Phase 2: Core Network Functions

#### NRF (Network Repository Function)
- [ ] Implement service registration endpoint
- [ ] Implement service discovery endpoint
- [ ] Create service profile storage
- [ ] Add heartbeat mechanism
- [ ] Implement proper error handling and logging
- [ ] Write unit tests
- [ ] Create Kubernetes deployment manifests
- [ ] Verify NRF functionality in isolation

#### AMF (Access and Mobility Management Function)
- [ ] Implement registration handling
- [ ] Create N1/N2 interface handlers
- [ ] Implement mobility management
- [ ] Set up connection with NRF
- [ ] Implement NAS security
- [ ] Add proper metrics and health checks
- [ ] Write unit and integration tests
- [ ] Create Kubernetes deployment manifests

#### SMF (Session Management Function)
- [ ] Implement session establishment
- [ ] Create IP address allocation mechanism
- [ ] Set up UPF selection logic
- [ ] Implement QoS handling
- [ ] Add proper metrics and health checks
- [ ] Write unit and integration tests
- [ ] Create Kubernetes deployment manifests

#### UPF (User Plane Function)
- [ ] Implement packet routing simulation
- [ ] Create QoS enforcement logic
- [ ] Implement packet inspection
- [ ] Add proper metrics and health checks
- [ ] Write unit and integration tests
- [ ] Create Kubernetes deployment manifests

#### UDM (Unified Data Management)
- [ ] Implement subscriber data storage
- [ ] Create profile management endpoints
- [ ] Set up authentication data handling
- [ ] Add proper metrics and health checks
- [ ] Write unit and integration tests
- [ ] Create Kubernetes deployment manifests

#### AUSF (Authentication Server Function)
- [ ] Implement authentication flows
- [ ] Create security context handling
- [ ] Set up integration with UDM
- [ ] Add proper metrics and health checks
- [ ] Write unit and integration tests
- [ ] Create Kubernetes deployment manifests

#### PCF (Policy Control Function)
- [ ] Implement policy rule engine
- [ ] Create policy enforcement points
- [ ] Set up QoS profile management
- [ ] Add proper metrics and health checks
- [ ] Write unit and integration tests
- [ ] Create Kubernetes deployment manifests

### Phase 3: Advanced Features

#### Network Slicing
- [ ] Define Custom Resource Definitions for Network Slices
- [ ] Implement NSSF functionality
- [ ] Create slice selection logic
- [ ] Implement slice differentiation in UPF
- [ ] Write integration tests for slice operations
- [ ] Create example slice templates (eMBB, URLLC, mMTC)

#### End-to-End Flows
- [ ] Implement UE registration flow across NFs
- [ ] Create PDU session establishment flow
- [ ] Implement mobility scenarios
- [ ] Add user plane data simulation
- [ ] Create comprehensive E2E tests

#### Observability
- [ ] Set up Prometheus metrics for all NFs
- [ ] Create Grafana dashboards
- [ ] Implement distributed tracing with Jaeger
- [ ] Set up structured logging with EFK stack
- [ ] Create alerting rules

### Phase 4: Simulation and Testing

- [ ] Build a basic RAN simulator for traffic generation
- [ ] Implement performance testing framework
- [ ] Create load testing scenarios
- [ ] Add chaos testing with Chaos Mesh
- [ ] Design demo scenarios showcasing key 5G features

## Development Requirements

- Go 1.21+
- Docker
- Kubernetes cluster (local: kind, minikube, or k3d)
- kubectl
- Helm
- Protocol Buffer compiler
- OpenAPI generator

## Documentation

Detailed documentation for each component can be found in the `/docs` directory.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
