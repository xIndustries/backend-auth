### Create backendauth database:
CREATE DATABASE "backendauth";

### Create Users table:
CREATE TABLE users (
    id UUID PRIMARY KEY,             -- Unique user ID (UUID)
    auth0_id VARCHAR(255) UNIQUE NOT NULL, -- Auth0 unique identifier
    email VARCHAR(255) UNIQUE NOT NULL,    -- User's email address
    username VARCHAR(50),             -- Optional username
    created_at TIMESTAMP NOT NULL DEFAULT NOW() -- Timestamp of user creation
);

### Generate protobufs
protoc --proto_path=proto \
       --go_out=proto/Generated \
       --go-grpc_out=proto/Generated \
       proto/user.proto


### Generate the SWIFT protobuf on the root folder /Mahler
protoc --proto_path=Networking/Protobuf \
       --swift_out=Networking/Protobuf/Generated \
       --grpc-swift_out=Networking/Protobuf/Generated \
       Networking/Protobuf/user.proto
