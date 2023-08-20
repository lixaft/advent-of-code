#include <iostream>
#include <string>

int main() {
  int max = 0;

  int current = 0;
  std::string line;
  while (std::getline(std::cin, line)) {
    if (line.empty()) {
      if (current > max) {
        max = current;
      }
      current = 0;
    } else {
      current += std::stoi(line);
    }
  }

  std::cout << max << '\n';
}
