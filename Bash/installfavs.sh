#!/bin/bash

# List of packages to install
packages=("vim" "btop" "julia")

# Output file
output_file="installeditems.txt"

# Write header with date
echo "Items installed on $(date)" > "$output_file"
echo "-------------------------" >> "$output_file"

# Function to install and report
install_package() {
  package=$1
  echo "Installing $package..."

  if sudo dnf install -y "$package" >/dev/null 2>&1; then
    echo "$package: SUCCESS"
    echo "$package" >> "$output_file"
  else
    echo "$package: FAILED"
  fi
}

# Loop through packages
for pkg in "${packages[@]}"; do
  install_package "$pkg"
done

echo "Installed items written to $output_file"
