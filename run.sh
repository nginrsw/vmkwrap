#!/bin/bash

# Create a bin folder if not there
mkdir -p bin

echo "Building VMKWrap..."
go build -o bin/vmkwrap main.go

# Check whether the build is successful
if [ $? -eq 0 ]; then
    echo "Build successful! 🎉"
    echo ""
    echo "You can now use VMKWrap by running:"
    echo "  ./bin/vmkwrap <file.vmk>"
    echo ""
    echo "For easier access, move the binary to a global location:"
    echo "  sudo ln -s ~/username/your-dir/bin/vmkwrap /usr/local/bin/"
    echo "Then you can run it from anywhere using:"
    echo "  vmkwrap <file.vmk>"
else
    echo "Build failed! ❌"
    exit 1
fi
