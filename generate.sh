#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Define directories
PROTO_DIR="/app/mev-protos"
OUTPUT_DIR="/app/pb"
GO_PACKAGE_BASE="github.com/Prophet-Solutions/jito-sdk/pb"

# Add go_package to each .proto file
add_go_package() {
    local proto_file=$1
    local package_name=$(basename "$proto_file" .proto)
    local go_package_line="option go_package = \"$GO_PACKAGE_BASE/$package_name\";"

    if ! grep -q "^option go_package" "$proto_file"; then
        awk -v pkg_line="$go_package_line" '
        {
            print $0
            if ($1 == "package") {
                print pkg_line
            }
        }' "$proto_file" > "$proto_file.tmp" && mv "$proto_file.tmp" "$proto_file"
    fi
}

for proto_file in "$PROTO_DIR"/*.proto; 
do
    add_go_package "$proto_file"
done

# Generate Go code, organizing output by package name
for proto_file in "$PROTO_DIR"/*.proto; 
do
    # Extract the package name from the .proto filename
    package_name=$(basename "$proto_file" .proto)

    # Create an output directory for this package
    package_output_dir="$OUTPUT_DIR/$package_name"
    mkdir -p "$package_output_dir"

    # Generate Go code for this .proto file
    protoc -I "$PROTO_DIR" \
        --go_out="$package_output_dir" --go_opt=paths=source_relative \
        --go-grpc_out="$package_output_dir" --go-grpc_opt=paths=source_relative \
        "$proto_file"
done

echo "Protobuf generation completed successfully."
