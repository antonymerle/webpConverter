# webpConverter
A simple utility to batch-convert webp images to jpg.

## Usage
Invoke the program with the path of the directory where you have your webp pictures. The script will loop through every file, convert webp to jpg in place, then move the original webp in a subfolder.

```bash
webpConverter /home/user/Pictures
```

## Limitations
1. The script is not recursive.
2. You need ffmpeg installed on your system and included in your PATH.

## Building the program
```bash
go build
```

## License
GNU GPL