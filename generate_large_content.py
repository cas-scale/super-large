import os
import random
import string

def generate_random_code(size_mb):
    """Generates a string of random 'code-like' text of a specific size in MB."""
    boilerplate = [
        "function processData(data) {\n",
        "  const result = data.map(item => {\n",
        "    return { ...item, processed: true, timestamp: Date.now() };\n",
        "  });\n",
        "  return result;\n",
        "}\n",
        "class DataManager {\n",
        "  constructor(name) { this.name = name; this.storage = []; }\n",
        "  addItem(item) { this.storage.push(item); }\n",
        "}\n"
    ]
    
    content = "".join(boilerplate)
    # Fill the rest with random comments to reach the desired size
    target_bytes = size_mb * 1024 * 1024
    current_bytes = len(content.encode('utf-8'))
    
    if current_bytes < target_bytes:
        padding_needed = target_bytes - current_bytes
        # Generate random string for padding
        padding = ''.join(random.choices(string.ascii_letters + string.digits + " ", k=padding_needed))
        content += f"\n/*\n{padding}\n*/\n"
        
    return content

def main():
    target_dir = "extra-large-content"
    os.makedirs(target_dir, exist_ok=True)
    
    print(f"Generating 1GB of additional code in {target_dir}...")
    
    # Generate 10 files of 100MB each
    for i in range(10):
        filename = f"large_module_{i}.js"
        filepath = os.path.join(target_dir, filename)
        content = generate_random_code(100)
        with open(filepath, 'w') as f:
            f.write(content)
        print(f"Created {filename} (100MB)")

    print("Generation complete.")

if __name__ == "__main__":
    main()