<<<<<<< HEAD
# Goku - JSON ↔ YAML Converter CLI

**Goku** is a simple and fast command-line tool built in Go to convert between **JSON** and **YAML** formats.

---

## ✨ Features

- Convert JSON to YAML
- Convert YAML to JSON
- Supports `.json`, `.yaml`, and `.yml` files
- Automatic output filename generation
- Custom output filename support
- Clean and user-friendly CLI using Cobra
- Proper error handling and validation

---

## 📦 Installation

### From Source

```bash
git clone https://github.com/sharifulb07/goku.git
cd goku
go install .





## 🐳 Docker

You can run Goku without installing Go using Docker:

```bash
# Build
docker build -t goku .

# Usage
docker run --rm -v $(pwd):/data -w /data goku -i input.yaml -o json

or 
# Build the image
docker build -t goku:latest .

# Run the container
docker run --rm -v $(pwd):/data -w /data goku -i data.yaml -o json -f output.json
=======
# goku
a cli project using cobra cli 
>>>>>>> 14fb2773561b8cc0d6728bb57e613623d547bceb
