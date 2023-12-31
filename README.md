# Entropy

### Overview
**Entropy** contains set of optimized functions for calculating file entropy using the Shannon algorithm. The implementation is designed for efficiency and accuracy, making it suitable for real-time entropy calculation, particularly for large files.

### Features
- **Shannon Entropy Calculation**: The repository utilizes the Shannon entropy formula to measure the randomness within a set of data. The formula is given by:  $$H = -1 \times \sum \frac{count_i}{N} \times \log_2\frac{count_i}{N}$$, where $H$ is the entropy, $count_i$ is the count of the $i$-th element, and $N$ is the total number of elements.
- **Optimized Implementation**: The functions are optimized to ensure efficient and accurate entropy calculation, even for large files.
- **File Compression and Encryption**: The functions take into account the effects of file compression and encryption on entropy, providing valuable insights for various applications such as malware detection and forensic analysis.
- **Cross-Platform Compatibility**: The code is designed to work with various file formats and operating systems, ensuring its versatility and applicability in different environments.

### Usage
To calculate the entropy of a file, simply use the provided function and pass the file as an input. The function will return the Shannon entropy value for the given file.

```go
package main

import (
	"fmt"
	"github.com/99pouria/entropy"
)

func main() {
	filePath := "path_to_your_file"
	entropy := entropy.CalculateFileEntropy(filePath)
	fmt.Println("File Entropy:", entropy)
}
```

Feel free to use this repository to incorporate efficient file entropy calculation into your Golang projects. If you have any questions or suggestions, please don't hesitate to reach out.
