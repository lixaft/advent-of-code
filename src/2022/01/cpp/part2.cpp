#include <iostream>
#include <string>

int main() {
  int max[3] = {0, 0, 0};

  int current = 0;
  std::string line;
  while (std::getline(std::cin, line)) {
    if (line.empty()) {
      if (current > max[0]) {
        max[2] = max[1];
        max[1] = max[0];
        max[0] = current;
      } else if (current > max[1]) {
        max[2] = max[1];
        max[1] = current;
      } else if (current == max[1]) {
        max[2] += 1;
      }
      current = 0;
    } else {
      current += std::stoi(line);
    }
  }

  std::cout << max[0] + max[1] + max[2] << '\n';
}
