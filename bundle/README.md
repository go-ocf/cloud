# Scalable OCF Cloud Hosting / Testing
Being [plugged.in](https://pluggedin.cloud) provides you a complete set of tools and services to manage your devices at scale. Allowing for the processing of real-time device data and interconnection of your devices and applications based on an interoperable standard. Interconnect, monitor and manage your devices in a cloud native way.

# OCF Cloud Bundle
Provides a simple docker cloud image for **testing purpose**.

## Features
- [OCF Native Cloud](https://openconnectivity.org/specs/OCF_Device_To_Cloud_Services_Specification_v2.1.0.pdf)
- OAUTH Athorization code is not verified
- [GRPC](https://github.com/plgd-dev/cloud/blob/master/grpc-gateway/pb/service.proto)

## Supported clients
- [iotivity v2+](https://github.com/iotivity/iotivity)
- [iotivity-lite v2+](https://github.com/iotivity/iotivity-lite)


## Pull the image
```bash
docker pull plgd/bundle:vnext
```

## Configuration
Image can be configured via enviroment variables as argument `-e ENV=VALUE` of command `docker`:
| ENV variable | Type | Description | Default |
| --------- | ----------- | ------- | ------- |
| `FQDN` | string | public FQDN for bundle | `"localhost"` |
| `COAP_GATEWAY_UNSECURE_PORT` | uint16 | exposed public port for coap-tcp  | `"5683"` |
| `COAP_GATEWAY_UNSECURE_ADDRESS` | string | coap-tcp listen address | `"0.0.0.0:5683"` |
| `COAP_GATEWAY_PORT` | uint16 | exposed public port for coaps-tcp  | `"5684"` |
| `COAP_GATEWAY_ADDRESS` | string | coaps-tcp listen address | `"0.0.0.0:5684"` |
| `COAP_GATEWAY_CLOUD_ID` | string | cloud id | `"00000000-0000-0000-0000-000000000001"` |
| `COAP_GATEWAY_DISABLE_BLOCKWISE_TRANSFER`| bool | disable blockwise transfer | `false` |
| `COAP_GATEWAY_BLOCKWISE_TRANSFER_SZX` | string | blockwise transfer size | `"1024"` |
| `COAP_GATEWAY_DISABLE_PEER_TCP_SIGNAL_MESSAGE_CSMS` | bool | ignore tcp control signal message from peer | `"false"`|
| `COAP_GATEWAY_DISABLE_VERIFY_CLIENTS`| bool | disable verifying coap clients certificates | `true` |
| `GRPC_GATEWAY_ADDRESS`| string | secure grpc-tcp listen address | `"0.0.0.0:9084"` |
| `GRPC_GATEWAY_DISABLE_VERIFY_CLIENTS`| bool | disable verifying grpc clients certificates | `true` |
| `HTTP_GATEWAY_ADDRESS`| string | secure grpc-tcp listen address | `"0.0.0.0:9086"` |
| `HTTP_GATEWAY_DISABLE_VERIFY_CLIENTS`| bool | disable verifying http clients certificates | `true` |
| `INITIALIZE_CERITIFICATES` | bool | initialze certificates | `true` |
| `CERITIFICATES_PATH` | string | path to directory | `"/data/certs"` |
| `MONGO_PATH` | string | path to directory | `"/data/db"` |
| `MONGO_PORT` | uint16 | mongo listen port  | `"10000"` |
| `NATS_PORT` | uint16 | nats listen port  | `"10001"` |
| `LOGS_PATH` | string | path to directory | `"/data/log"` |

## Run
```bash
docker run -d --network=host --name=cloud -t plgd/bundle:vnext
```

## Device Onboarding
The onboarding values which should be set to the [coapcloudconf](https://github.com/openconnectivityfoundation/cloud-services/blob/c2c/swagger2.0/oic.r.coapcloudconf.swagger.json) device resource are:

### Unsecured device
| Attribute | Value |
| --------- | ------|
| `apn` | `test` |
| `cis` | `coap+tcp://127.0.0.1:5683` |
| `sid` | `same as is set in COAP_GATEWAY_CLOUD_ID` |
| `at` | `test` |

#### Unsecured iotivity-lite sample device example:

```bash
# Start the cloud container with "unsecured" parameters
docker run -d --network=host --name=cloud -t plgd/bundle:vnext \
-e COAP_GATEWAY_CLOUD_ID="00000000-0000-0000-0000-000000000001" \
-e COAP_GATEWAY_UNSECURE_PORT="5683" \
-e COAP_GATEWAY_UNSECURE_ADDRESS="0.0.0.0:5683" \
-e "coap+tcp://127.0.0.1:5683"

```

```bash
# Retrieve iotivity-lite project
git clone --recursive https://github.com/iotivity/iotivity-lite.git
cd ./iotivity-lite/port/linux

# Build and run unsecured applications
make CLOUD=1 SECURE=0 cloud_server cloud_client

# Start unsecured device sample
./cloud_server cloud_server test coap+tcp://127.0.0.1:5683 00000000-0000-0000-0000-000000000001 test

# Start unsecured client
./cloud_client cloud_client test coap+tcp://127.0.0.1:5683 00000000-0000-0000-0000-000000000001 test

```

### Secured device
| Attribute | Value |
| --------- | ------|
| `apn` | `test`|
| `cis` | `coaps+tcp://127.0.0.1:5684` |
| `sid` | `same as is set in COAP_GATEWAY_CLOUD_ID` |
| `at` | `test` |

#### Conditions:
- Device must be owned.
- Cloud CA must be set as TRUST CA with subject COAP_GATEWAY_CLOUD_ID in device.
- Cloud CA in PEM:
  ```bash
  docker exec -it cloud cat CERITIFICATES_PATH/root_ca.crt
  ```
- ACL for Cloud (Subject: COAP_GATEWAY_CLOUD_ID) must be set with full access to all published resources in device.

#### Secured iotivity-lite sample device example:

```bash
# Start the cloud container with "secured" parameters
docker run -d --network=host --name=cloud -t plgd/bundle:vnext \
-e COAP_GATEWAY_CLOUD_ID="00000000-0000-0000-0000-000000000001" \
-e COAP_GATEWAY_PORT="5684" \
-e COAP_GATEWAY_ADDRESS="0.0.0.0:5684" \
-e "coaps+tcp://127.0.0.1:5684"

```

```bash
# Retrieve iotivity-lite project
git clone --recursive https://github.com/iotivity/iotivity-lite.git
cd ./iotivity-lite/port/linux

# Then build secured applications and onboarding_tool
make CLOUD=1 SECURE=1 PKI=1 onboarding_tool cloud_server cloud_client

# Retrieve cloud CA and set it as "pki_certs/cloudca.pem", to make the sample import this as TRUST CA
docker exec -it cloud cat /data/certs/root_ca.crt > pki_certs/cloudca.pem

# Start secured device sample
./cloud_server cloud_server test coaps+tcp://127.0.0.1:5684 00000000-0000-0000-0000-000000000001 test
# Start secured client
./cloud_client cloud_client test coaps+tcp://127.0.0.1:5684 00000000-0000-0000-0000-000000000001 test

# Execute onboarding_tool and:
# - Onboard and set ACLs to the device sample as described in https://github.com/iotivity/iotivity-lite#step-1-onboard-and-provision-the-server
# - Onboard the client as described in https://github.com/iotivity/iotivity-lite#step-2-onboard-the-client
# - Pair client and server as described in https://github.com/iotivity/iotivity-lite#step-3-pair-server-and-client
./onboarding_tool

# Then restart device sample and client

```

## Build a COAP client application
To build the client you need to have **golang v1.13+**.
```bash
cd ./client/coap
go build
./coap --help
# gets a resource links of the registered devices from cloud
./coap --signUp test --href /oic/res
```

## Build a GRPC client application
To build the client you need to have **golang v1.13+**.
```bash
cd ./client/grpc
go build
./grpc --help
# gets all resources with contents from cloud
./grpc
# gets resources of device with contents from cloud
./grpc --deviceid {deviceID}
# gets devices from cloud
./grpc --getdevices
```

## HTTP access
[REST API](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/plgd-dev/cloud/master/http-gateway/swagger.yaml)
