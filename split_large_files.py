import os

def split_file(filepath, chunk_size_mb=1):
    """Splits a file into smaller chunks of specified size in MB."""
    chunk_size = chunk_size_mb * 1024 * 1024
    base_name = os.path.basename(filepath)
    dir_name = os.path.dirname(filepath)
    name_without_ext, ext = os.path.splitext(base_name)
    
    new_dir = os.path.join(dir_name, name_without_ext)
    os.makedirs(new_dir, exist_ok=True)
    
    print(f"Splitting {base_name} into {new_dir}...")
    
    try:
        with open(filepath, 'r') as f:
            chunk_idx = 0
            while True:
                data = f.read(chunk_size)
                if not data:
                    break
                chunk_filename = f"part_{chunk_idx}{ext}"
                chunk_path = os.path.join(new_dir, chunk_filename)
                with open(chunk_path, 'w') as chunk_file:
                    chunk_file.write(data)
                chunk_idx += 1
        
        # Remove the original large file after successful split
        os.remove(filepath)
        return True
    except Exception as e:
        print(f"Error splitting {filepath}: {e}")
        return False

def main():
    target_dir = "extra-large-content"
    if not os.path.exists(target_dir):
        print(f"Directory {target_dir} not found.")
        return
    
    files_to_split = [f for f in os.listdir(target_dir) if f.endswith('.js') and os.path.isfile(os.path.join(target_dir, f))]
    
    for filename in files_to_split:
        filepath = os.path.join(target_dir, filename)
        split_file(filepath, chunk_size_mb=1)

    print("Splitting complete.")

if __name__ == "__main__":
    main()